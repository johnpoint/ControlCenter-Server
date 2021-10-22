package depend

import (
	"ControlCenter/app/controller"
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/apiMiddleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Enable bool
}

func (r *Api) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	gin.SetMode(gin.ReleaseMode)
	routerGin := gin.New()
	routerGin.GET("/ping", controller.Pong)

	report := routerGin.Group("/report") // 数据上报接口
	{
		report.POST("/login") // 用户登录上报
	}

	auth := routerGin.Group("/auth")
	{
		auth.POST("/login")    // 登录
		auth.POST("/register") // 注册
	}

	api := routerGin.Group("/api")
	{
		api.GET("")                                           // 获取首页详情
		api.POST("/token", apiMiddleware.JWTAuthMiddleware()) // 更新 jwt
	}

	user := api.Group("/user") // 用户模块
	{
		user.GET("")        // 获取当前用户信息
		user.GET("/assets") // 获取当前用户资产列表
	}

	server := api.Group("/server") // 服务器模块
	{
		server.POST("")       // 服务器列表
		server.POST("/:uuid") // 服务器详细信息
	}

	certificate := api.Group("/certificate") // 证书模块
	{
		certificate.POST("")       // 证书列表
		certificate.POST("/:uuid") // 证书详细信息
	}

	configuration := api.Group("/configuration") // 配置文件模块
	{
		configuration.POST("")       // 配置文件列表
		configuration.POST("/:uuid") // 配置文件详细信息
	}

	go func() {
		fmt.Println("[init] HTTP Listen at " + cfg.HttpServerListen)
		err := routerGin.Run(cfg.HttpServerListen)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
