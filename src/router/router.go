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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowAddress,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	u := e.Group("/user")
	u.POST("/auth/login", apis.OaLogin)
	u.POST("/auth/register", apis.OaRegister)

	s := e.Group("/server")
	s.POST("/setup/:token", apis.SetupServer)
	s.GET("/now/:token", apis.GetNow)
	s.POST("/update/:token", apis.ServerUpdate)
	s.GET("/update/:token", apis.GetServerUpdate)

	e.GET("/", apis.Accessible)

	sys := e.Group("/system")
	sys.POST("/restart", apis.SysRestart)
	jwtConfig := middleware.JWTConfig{
		Claims:     &apis.JwtCustomClaims{},
		SigningKey: []byte(conf.Salt),
	}
	w := e.Group("/web")
	w.Use(middleware.JWTWithConfig(jwtConfig))
	w.POST("/debug/check", apis.CheckPower)
	w.GET("/ServerInfo", apis.GetServerInfo)
	w.PUT("/ServerInfo", apis.UpdateServerInfo) // TODO: remove this
	w.POST("/Server/:serverid/Server/:action", apis.AddClientEvent)
	w.POST("/Server/:serverid/Docker/:action/:id", apis.ChangeDockerStatus)
	w.GET("/ServerInfo/Certificate", apis.GetCertificateLinked)
	w.GET("/ServerInfo/Task", apis.GetServerEvents)
	w.GET("/ServerInfo/Site", apis.GetSiteLinked)
	w.DELETE("/Server/:id", apis.RemoveServer)
	w.GET("/DomainInfo", apis.GetDomainInfo)
	w.PUT("/DomainInfo", apis.UpdateDomainInfo)
	w.PUT("/UserInfo/:mail/:key/:value", apis.UpdateUserInfo)
	w.PUT("/SiteInfo", apis.AddSiteInfo)
	w.GET("/SiteInfo", apis.GetSiteInfo)
	w.DELETE("/SiteInfo", apis.DeleteSiteInfo)
	w.PUT("/Certificate", apis.AddCertificateInfo)
	w.GET("/Certificate", apis.GetCertificateInfo)
	w.POST("/Certificate", apis.UpdateCertificateInfo)
	w.DELETE("/Certificate/:id", apis.DeleteCertificateInfo)
	w.PUT("/link/Certificate/:ServerID/:CerID", apis.LinkServerCer)
	w.DELETE("/link/Certificate/:ServerID/:CerID", apis.UnLinkServerCer)
	w.PUT("/link/Site/:ServerID/:SiteID", apis.LinkServerSite)
	w.DELETE("/link/Site/:ServerID/:SiteID", apis.UnLinkServerSite)
	w.POST("/backup", apis.SetBackupFile)
	w.GET("/SiteInfo/check", apis.GetCertificatesInfo)
	w.POST("/Setting/:name/:value", apis.SetSetting)
	w.GET("/Setting/:name", apis.GetSetting)
	w.GET("/DockerInfo/:id", apis.GetDockerInfo)
	w.PUT("/DockerInfo", apis.AddDockerInfo)
	w.PUT("/DockerInfo/:id", apis.EditDockerInfo)
	e.GET("/web/:token/backup", apis.GetBackupFile)
	user := w.Group("/UserInfo")
	user.PUT("/level/:uid/:level", apis.ChangeLevel)
	user.GET("/Password/:oldpass/:newpass", apis.ReSetPassword)
	user.GET("/Token", apis.GetUserToken)
	user.PUT("/Token", apis.GetNewToken)
	user.GET("", apis.GetUserInfo)
	user.GET("/", apis.GetUserList)

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
