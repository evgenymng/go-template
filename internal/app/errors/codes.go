package errors

// Lightweight error code identifier.
type Code uint

const (
	CodeUnset Code = iota
	CodeUnknown
	// client
	CodeBadInput
	// server
	CodeUnexpected
	CodeUnauthorized
)

var codeToString = map[Code]string{
	CodeBadInput:     "bad_input",
	CodeUnexpected:   "unexpected",
	CodeUnauthorized: "unauthorized",
	CodeUnknown:      "unknown",
}

// Implement [fmt.Stringer] interface.
func (e Code) String() string {
	if code, ok := codeToString[e]; ok {
		return code
	}
	return CodeUnknown.String()
}
