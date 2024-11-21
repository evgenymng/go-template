package log

// Logger event tag. Guaranteed to be present in every log message.
//
// It defaults to the [TagUnset].
//
// [Tag] is a lightweight identifier. It is simply a [uint] under the hood,
// which transforms into [string] when needed.
type Tag uint

const (
	TagUnset Tag = iota
	TagUnknown
	// Lifecycle
	TagStartup
	TagRunning
	TagShutdown
	// Log
	TagLogParsing
	// Request/response
	TagRequest
	TagResponse
	TagMiddleware
	// Integrations
	TagSqlConnect
	TagSqlQuery
	// Custom
	// ...
)

var tagToString = map[Tag]string{
	TagUnknown:    "unknown",
	TagStartup:    "startup",
	TagShutdown:   "shutdown",
	TagLogParsing: "log_parsing",
	TagRequest:    "request",
	TagResponse:   "response",
}

// Implement [fmt.Stringer] interface.
func (e Tag) String() string {
	if tag, ok := tagToString[e]; ok {
		return tag
	}
	return TagUnknown.String()
}
