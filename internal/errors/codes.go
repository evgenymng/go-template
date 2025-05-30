package errors

// Error code identifier.
type Code string

const (
	CodeUnset      Code = ""
	CodeUnknown    Code = "unknown"
	CodeBadInput   Code = "bad_input"
	CodeUnexpected Code = "unexpected"
)
