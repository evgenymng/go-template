package middleware

import (
	"net/http"

	"go-template/internal/errors"
	"go-template/pkg/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Middleware to perform custom response-handling logic,
// which includes error-handling and uniform responses, whenever
// error happens.
//
// Use [gin.Context.Error] to register a new error during
// the request processing.
func ResponseHandler(traceIdKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetString(traceIdKey)

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]

			if err.Type == gin.ErrorTypeBind {
				log.S.Error("Request parameter binding error", zap.Error(err))
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
				log.S.Warn("Service error", zap.Error(err))

				// NOTE(evgenymng): switch by code here, if needed
				switch serr.Code {
				default:
					c.JSON(http.StatusInternalServerError, serr)
				}
				return
			}

			log.S.Error("Unexpected error", zap.Error(err))
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
