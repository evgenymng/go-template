package middleware

import (
	"net/http"

	"app/internal/app/ckey"
	"app/internal/app/errors"
	"app/internal/log"

	"github.com/gin-gonic/gin"
)

// Middleware to perform custom response-handling logic,
// which includes error-handling and uniform responses, whenever
// error happens.
//
// Use [gin.Context.Error] to register a new error during
// the request processing.
func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetString(string(ckey.TraceId))

		c.Next()

		l := log.L().
			TraceId(traceId)

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			l = l.Error(err)

			if err.Type == gin.ErrorTypeBind {
				log.S.Error("Request parameter binding error", l)
				c.JSON(http.StatusUnprocessableEntity, errors.E().
					Code(errors.CodeBadInput).
					Message("Request parameter binding error").
					TraceId(traceId).
					Inner(err).
					Build(),
				)
				return
			}

			if serr, ok := err.Err.(*errors.ServiceError); ok {
				log.S.Warn("Service error", l)

				// NOTE(evgenymng): switch by code here, if needed
				switch serr.Code {
				case errors.CodeUnauthorized:
					c.JSON(http.StatusUnauthorized, serr)
				default:
					c.JSON(http.StatusInternalServerError, serr)
				}
				return
			}

			log.S.Error("Unexpected error", l)
			c.JSON(http.StatusInternalServerError, errors.E().
				Code(errors.CodeUnexpected).
				Message("Unexpected error").
				TraceId(traceId).
				Inner(err).
				Build(),
			)
			// NOTE(evgenymng): we do not expect more than one error
		}
	}
}
