package telemetry

import (
	"context"
	"errors"
	"fmt"
	"go-template/pkg/config"
	"go-template/pkg/log"

	"github.com/go-logr/zapr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	otlplog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func Init(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// configure intenral otel logger
	logger := zapr.NewLogger(log.S.Desugar())
	otel.SetLogger(logger)

	// set up tracer provider
	tp, err := newTracerProvider(ctx)
	if err != nil {
		handleErr(err)
		return
	}
	otel.SetTracerProvider(tp)
	log.S.Debug("TracerProvider is initialized")

	// set up meter provider
	mp, err := newMeterProvider(ctx)
	if err != nil {
		handleErr(err)
		return
	}
	otel.SetMeterProvider(mp)
	log.S.Debug("MeterProvider is initialized")

	// set up logger provider
	loggerProvider, err := newLoggerProvider(ctx)
	if err != nil {
		handleErr(err)
		return
	}
	global.SetLoggerProvider(loggerProvider)
	log.S.Debug("LoggerProvider is initialized")

	return
}

func newTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(getGrpcAddress()),
	}
	if !config.C.Otel.Secure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	exporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.C.Otel.ServiceName),
		)),
	)
	return tp, nil
}

func newMeterProvider(ctx context.Context) (*metric.MeterProvider, error) {
	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(getGrpcAddress()),
	}
	if !config.C.Otel.Secure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}

	metricExporter, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)
	return meterProvider, nil
}

func newLoggerProvider(ctx context.Context) (*otlplog.LoggerProvider, error) {
	opts := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(getGrpcAddress()),
	}
	if !config.C.Otel.Secure {
		opts = append(opts, otlploggrpc.WithInsecure())
	}

	logExporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	loggerProvider := otlplog.NewLoggerProvider(
		otlplog.WithProcessor(otlplog.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}

func getGrpcAddress() string {
	return fmt.Sprintf("%s:%d", config.C.Otel.Host, config.C.Otel.Port)
}
