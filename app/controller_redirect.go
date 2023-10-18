package app

import (
	"hk4e-redirect/common/config"
	"hk4e-redirect/pkg/logger"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) redirectDispatch(context *gin.Context) {
	c.redirect(config.CONF.Redirect.DispatchUrl, context)
}

func (c *Controller) redirectSdk(context *gin.Context) {
	c.redirect(config.CONF.Redirect.SdkUrl, context)
}

func (c *Controller) redirect(url string, context *gin.Context) {
	client := &http.Client{}

	logger.Debug("url: %v", url+context.Request.RequestURI)
	req, err := http.NewRequest(context.Request.Method, url+context.Request.RequestURI, context.Request.Body)
	if err != nil {
		logger.Error("http new request error: %v", err)
		return
	}
	req.Header = context.Request.Header
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("body close error: %v", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("read all error: %v", err)
		return
	}
	_, err = context.Writer.Write(body)
	if err != nil {
		logger.Error("write body error: %v", err)
		return
	}
	for name, value := range resp.Header {
		for _, s := range value {
			context.Header(name, s)
		}
	}
	context.Status(resp.StatusCode)
}
