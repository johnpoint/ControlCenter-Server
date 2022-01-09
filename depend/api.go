package depend

import (
	"ControlCenter/app/controller"
	"ControlCenter/config"
	"ControlCenter/pkg/apiMiddleware"
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/bootstrap"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Api api服务
type Api struct{}

var _ bootstrap.Component = (*Api)(nil)

func (d *Api) Init(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	routerGin := gin.New()
	routerGin.Use(apiMiddleware.LogPlusMiddleware())
	routerGin.GET("/ping", controller.Pong)

	report := routerGin.Group("/report") // 数据上报接口
	{
		report.POST("/login", controller.Pong) // 用户登录上报
	}

	auth := routerGin.Group("/auth")
	{
		auth.POST("/login", controller.Login)   // 登录
		auth.POST("/register", controller.Pong) // 注册
	}

	api := routerGin.Group("/api", session.MiddlewareFunc())
	{
		api.GET("", controller.Pong) // 获取首页详情
	}

	user := api.Group("/user") // 用户模块
	{
		user.GET("", controller.Pong)        // 获取当前用户信息
		user.GET("/assets", controller.Pong) // 获取当前用户资产列表
	}

	assets := api.Group("/assets") // 资产相关(这些是资产的元数据信息，而不包括资产的内容)
	{
		assets.GET("/:uuid", controller.Pong)    // 获取资产信息
		assets.POST("/:uuid", controller.Pong)   // 修改资产相关信息
		assets.DELETE("/:uuid", controller.Pong) // 删除资产
	}

	server := api.Group("/server") // 服务器模块
	{
		server.GET("", controller.Pong)                         // 服务器列表
		server.GET("/:uuid", controller.Pong)                   // 服务器详细信息
		server.POST("/chart", controller.ServerChartController) // 服务器性能信息绘图
	}

	certificate := api.Group("/certificate") // 证书模块
	{
		certificate.GET("", controller.Pong)       // 证书列表
		certificate.GET("/:uuid", controller.Pong) // 证书详细信息
	}

	configuration := api.Group("/configuration") // 配置文件模块
	{
		configuration.GET("", controller.Pong)       // 配置文件列表
		configuration.GET("/:uuid", controller.Pong) // 配置文件详细信息
	}

	go func() {
		fmt.Println("[init] HTTP Listen at " + config.Config.HttpServerListen)
		err := routerGin.Run(config.Config.HttpServerListen)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
