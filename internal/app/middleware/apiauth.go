package middleware

import (
	"regexp"

	"app/internal/app/ckey"
	"app/internal/app/errors"
	"app/internal/util"

	"github.com/gin-gonic/gin"
)

// Middleware implementing the API key authentication mechanism.
func ApiAuthMiddleware(
	apiKeys []string,
	header string,
	allowUri []string,
) gin.HandlerFunc {
	patterns := util.Map(
		allowUri,
		func(s string) *regexp.Regexp { return regexp.MustCompile(s) },
	)

	return func(c *gin.Context) {
		traceId := c.GetString(string(ckey.TraceId))
		reqPath := []byte(c.Request.URL.Path)
		allow := util.Any(util.Map(patterns, func(re *regexp.Regexp) bool {
			return re.Match(reqPath)
		}))
		if len(apiKeys) == 0 || allow {
			c.Next()
			return
		}

		key := c.GetHeader("Authorization")
		if !util.Contains(apiKeys, key) {
			_ = c.Error(
				errors.E().
					Code(errors.CodeUnauthorized).
					Message("API key is invalid or not provided. Check the Authorization header.").
					TraceId(traceId).
					Build(),
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
