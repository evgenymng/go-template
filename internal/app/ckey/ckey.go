// Context key. Basically constants for easier access to the
// [Context]-specific values.
package ckey

type ContextKey string

const (
	TraceId ContextKey = "trace_id"
)
