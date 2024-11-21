package errors

import "fmt"

// Any error this service might potentially produce.
type ServiceError struct {
	Code    Code    `json:"code"`
	Message string  `json:"message"`
	TraceId string  `json:"trace_id"`
	Inner   []error `json:"-"` // cause, etc.
}

// Any error this service might potentially produce.
type ServiceErrorBuilder struct {
	code    Code
	message string
	traceId string
	inner   []error
}

func E() ServiceErrorBuilder {
	return ServiceErrorBuilder{}
}

func (b ServiceErrorBuilder) Code(c Code) ServiceErrorBuilder {
	b.code = c
	return b
}

func (b ServiceErrorBuilder) Message(m string) ServiceErrorBuilder {
	b.message = m
	return b
}

func (b ServiceErrorBuilder) TraceId(t string) ServiceErrorBuilder {
	b.traceId = t
	return b
}

func (b ServiceErrorBuilder) Inner(i ...error) ServiceErrorBuilder {
	b.inner = i
	return b
}

func (b ServiceErrorBuilder) Build() ServiceError {
	return ServiceError{
		Code:    b.code,
		Message: b.message,
		TraceId: b.traceId,
		Inner:   b.inner,
	}
}

// Implement the [error] interface.
func (e ServiceError) Error() string {
	s := fmt.Sprintf("service error, code: %q", e.Code.String())
	if e.Message != "" {
		s += fmt.Sprintf(", message: %q", e.Message)
	}
	if e.TraceId != "" {
		s += fmt.Sprintf(", trace ID: %q", e.TraceId)
	}
	if e.Inner != nil {
		s += ", caused by"
		delim := ": "
		for _, ie := range e.Inner {
			s += delim + fmt.Sprintf("%q", ie.Error())
			delim = "; "
		}
	}
	return s
}

// Implement the [error] interface.
func (e ServiceError) Unwrap() []error {
	return e.Inner
}
