package router

import (
	"github.com/johnpoint/ControlCenter-Server/src/apis"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	go checkOnlineI()
	conf := config.LoadConfig()
	e := echo.New()
	e.Debug = conf.Debug
	e.HideBanner = true

	//echo中间件配置
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowAddress,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	})) //CORS配置
	jwtConfig := middleware.JWTConfig{
		Claims:     &apis.JwtCustomClaims{},
		SigningKey: []byte(conf.Salt),
	} //JsonWebToken 配置

	//用户校验api
	u := e.Group("/user")
	u.POST("/auth/login", apis.OaLogin)       //登录
	u.POST("/auth/register", apis.OaRegister) //注册

	//客户端调用api
	s := e.Group("/server")
	s.POST("/setup/:token", apis.SetupServer)     //注册服务器至数据库
	s.GET("/now/:token", apis.GetNow)             //获取服务器当前任务
	s.POST("/update/:token", apis.ServerUpdate)   //更新服务器状态至中心控制服务器
	s.GET("/update/:token", apis.GetServerUpdate) //获取服务器配置

	e.GET("/", apis.Accessible) //服务端运行验证

	//服务端指令api
	sys := e.Group("/system")
	sys.POST("/restart", apis.SysRestart) //重启服务端 TODO
	sys.GET("/info", apis.GetSystemInfo)  //获取服务端服务器性能信息

	//前端调用部分
	w := e.Group("/web")
	w.Use(middleware.JWTWithConfig(jwtConfig))
	w.POST("/debug/check", apis.CheckPower)
	w.GET("/ServerInfo", apis.GetServerInfo)                                //获取服务器信息
	w.PUT("/ServerInfo", apis.UpdateServerInfo)                             //编辑服务器
	w.POST("/Server/:serverid/Server/:action", apis.AddClientEvent)         //服务器队列添加任务
	w.POST("/Server/:serverid/Docker/:action/:id", apis.ChangeDockerStatus) //服务器队列添加任务-docker管理
	w.GET("/ServerInfo/Certificate", apis.GetCertificateLinked)             //获取服务器证书列表
	w.GET("/ServerInfo/Task", apis.GetServerEvents)                         //获取服务器任务队列
	w.DELETE("/Server/:id", apis.RemoveServer)                              //删除服务器
	w.PATCH("/UserInfo/:id/:key/:value", apis.UpdateUserInfo)               //更新用户信息
	w.PUT("/Certificate", apis.AddCertificateInfo)                          //添加SSL证书
	w.GET("/Certificate", apis.GetCertificateInfo)                          //获取SSL证书信息
	w.POST("/Certificate", apis.UpdateCertificateInfo)                      //更新SSL证书信息
	w.DELETE("/Certificate/:id", apis.DeleteCertificateInfo)                //删除SSL证书
	w.GET("/Configuration", apis.GetConfigurationInfo)                      //获取配置文件信息
	w.PUT("/Configuration", apis.AddConfigurationInfo)                      //新增配置文件
	w.PUT("/link/Certificate/:ServerID/:CerID", apis.LinkServerCer)         //分配证书给服务器
	w.DELETE("/link/Certificate/:ServerID/:CerID", apis.UnLinkServerCer)    //将证书的分配记录删除
	w.POST("/backup", apis.SetBackupFile)                                   //获取系统数据库
	w.POST("/Setting/:name/:value", apis.SetSetting)                        //更改个人设置
	w.GET("/Setting/:name", apis.GetSetting)                                //获取个人设置
	e.GET("/web/:token/backup", apis.GetBackupFile)                         //上传数据库并覆盖

	//用户信息相关api
	user := w.Group("/UserInfo")
	user.PATCH("/level/:uid/:level", apis.ChangeLevel) //更新用户权限组
	user.POST("/Password", apis.ReSetPassword)         //更新用户密码
	user.GET("/Token", apis.GetUserToken)              //获取用户 token
	user.PUT("/Token", apis.GetNewToken)               //更新用户 token
	user.GET("", apis.GetUserInfo)                     //获取用户信息
	user.GET("/", apis.GetUserList)                    //获取用户列表

	if conf.TLS {
		e.Logger.Fatal(e.StartTLS(":"+conf.ListenPort, conf.CERTPath, conf.KEYPath))
	} else {
		e.Logger.Fatal(e.Start(":" + conf.ListenPort))
	}
}

func checkOnlineI() {
	for true {
		apis.CheckOnline()
	}
}
