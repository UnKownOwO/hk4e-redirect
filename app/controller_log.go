package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ignoreLog(context *gin.Context) {
	context.JSON(http.StatusOK, "ok")
}
