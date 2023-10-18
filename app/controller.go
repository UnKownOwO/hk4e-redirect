package app

import (
	"hk4e-redirect/common/config"
	"hk4e-redirect/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	c := new(Controller)

	go c.initRoute()

	return c
}

// initRoute 初始化路由
func (c *Controller) initRoute() {
	if config.GetConfig().Logger.Level == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	{
		// dispatch
		engine.GET("/query_security_file", c.redirectDispatch)
		engine.GET("/query_region_list", c.redirectDispatch)
		engine.GET("/query_cur_region", c.redirectDispatch)
	}
	{
		// 登录
		// 转发至sdk
		// 账号登录
		engine.POST("/hk4e_:name/mdk/shield/api/login", c.redirectSdk)
		// token登录
		engine.POST("/hk4e_:name/mdk/shield/api/verify", c.redirectSdk)
		// 获取combo token
		engine.POST("/hk4e_:name/combo/granter/login/v2/login", c.redirectSdk)
	}
	{
		// 日志
		engine.POST("/sdk/dataUpload", c.ignoreLog)
		engine.GET("/perf/config/verify", c.ignoreLog)
		engine.POST("/perf/dataUpload", c.ignoreLog)
		engine.POST("/log", c.ignoreLog)
		engine.POST("/crash/dataUpload", c.ignoreLog)
	}
	{
		// 收集数据
		engine.GET("/device-fp/api/getExtList", c.redirectSdk)
		engine.POST("/device-fp/api/getFp", c.redirectSdk)
	}
	{
		// 返回固定数据
		// Windows
		engine.GET("/hk4e_:name/mdk/agreement/api/getAgreementInfos", c.redirectSdk)
		engine.POST("/hk4e_:name/combo/granter/api/compareProtocolVersion", c.redirectSdk)
		engine.POST("/account/risky/api/check", c.redirectSdk)
		engine.GET("/combo/box/api/config/sdk/combo", c.redirectSdk)
		engine.GET("/hk4e_:name/combo/granter/api/getConfig", c.redirectSdk)
		engine.GET("/hk4e_:name/mdk/shield/api/loadConfig", c.redirectSdk)
		engine.POST("/data_abtest_api/config/experiment/list", c.redirectSdk)
		// Android
		engine.POST("/common/h5log/log/batch", c.redirectSdk)
		engine.GET("/hk4e_:name/combo/granter/api/getFont", c.redirectSdk)
	}
	{
		// 静态资源
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_oversea/m2020030410/m2020030410-version.json HTTP/1.1
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_oversea/m2020030410/m2020030410-zh-cn.json HTTP/1.1
		engine.StaticFS("/admin/mi18n/plat_oversea/m2020030410", http.Dir("./static/m2020030410"))
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_os/m09291531181441/m09291531181441-version.json HTTP/1.1
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_os/m09291531181441/m09291531181441-zh-cn.json HTTP/1.1
		engine.StaticFS("/admin/mi18n/plat_os/m09291531181441", http.Dir("./static/m09291531181441"))
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_oversea/m202003049/m202003049-version.json HTTP/1.1
		// GET https://webstatic-sea.hoyoverse.com/admin/mi18n/plat_oversea/m202003049/m202003049-zh-cn.json HTTP/1.1
		engine.StaticFS("/admin/mi18n/plat_oversea/m202003049", http.Dir("./static/m202003049"))
	}
	port := config.GetConfig().HttpPort
	addr := ":" + strconv.Itoa(int(port))
	if port == 443 {
		err := engine.RunTLS(addr, config.CONF.CertPath, config.CONF.KeyPath)
		if err != nil {
			logger.Error("gin run tls error: %v", err)
			return
		}
	} else {
		err := engine.Run(addr)
		if err != nil {
			logger.Error("gin run error: %v", err)
		}
	}
}
