package log

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

type LogObject struct {
	tag     Tag
	error   error
	traceId string
	data    []any
}

func L() LogObject {
	return LogObject{}
}

func (l LogObject) Error(e error) LogObject {
	l.error = e
	return l
}

func (l LogObject) TraceId(id string) LogObject {
	l.traceId = id
	return l
}

func (l LogObject) Add(k string, v any) LogObject {
	// NOTE(evgenymng): no actual array copying is supposed to happen
	// when you call this method. Remember that slices are simply
	// views into the underlying data storage, so they may share it.
	l.data = append(l.data, k, v)
	return l
}

func (l LogObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	// tag should always be present
	enc.AddString("tag", l.tag.String())

	// in case there is an error
	if l.error != nil {
		enc.AddString("error", l.error.Error())
		enc.AddString("error_type", fmt.Sprintf("%T", l.error))
	}

	// optional stuff
	if l.traceId != "" {
		enc.AddString("trace_id", l.traceId)
	}

	if len(l.data)%2 == 0 {
		for i := 0; i < len(l.data)-1; i = i + 2 {
			var ok bool
			var strKey string
			if strKey, ok = l.data[i].(string); !ok {
				S.Error(
					"L's data array has a key that cannot be cast to string",
					L(),
				)
				continue
			}
			// NOTE(evgenymng): the docs say that this function might be slow
			// and allocation-heavy.
			err := enc.AddReflected(strKey, l.data[i+1])
			if err != nil {
				return err
			}
		}
	} else {
		S.Error(
			"L's data array isn't of even size, ignoring",
			L(),
		)
	}
	return nil
}
