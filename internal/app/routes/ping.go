package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping controller.
//
//	@summary	Ping the service
//	@tags		misc
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{string}	string
//
//	@router		/ping [get]
func GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
