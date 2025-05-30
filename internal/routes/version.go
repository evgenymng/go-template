package routes

import (
	"net/http"

	"go-template/pkg/config"

	"github.com/gin-gonic/gin"
)

// Version controller.
//
//	@summary	Returns the version of the service.
//	@tags		misc
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{string}	string
//
//	@router		/version [get]
func GetVersion(c *gin.Context) {
	c.String(http.StatusOK, config.C.Version)
}
