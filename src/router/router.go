package router

import (
	. "github.com/johnpoint/ControlCenter-Server/src/apis"
	. "github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	go checkOnlineI()
	conf := LoadConfig()
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowAddress,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	u := e.Group("/user")
	u.POST("/auth/login", OaLogin)
	u.POST("/auth/register", OaRegister)

	s := e.Group("/server")
	s.POST("/setup/:token", SetupServer)
	s.GET("/now/:token", GetNow)
	s.POST("/update/:token", ServerUpdate)
	s.GET("/update/:token", GetServerUpdate)

	e.GET("/", Accessible)

	sys := e.Group("/system")
	sys.POST("/restart", SysRestart)
	jwtConfig := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(conf.Salt),
	}
	w := e.Group("/web")
	w.Use(middleware.JWTWithConfig(jwtConfig))
	w.POST("/debug/check", CheckPower)
	w.GET("/ServerInfo", GetServerInfo)
	w.PUT("/ServerInfo", UpdateServerInfo)
	w.GET("/ServerInfo/Certificate", GetCertificateLinked)
	w.GET("/ServerInfo/Site", GetSiteLinked)
	w.DELETE("/Server/:id", RemoveServer)
	w.GET("/DomainInfo", GetDomainInfo)
	w.PUT("/DomainInfo", UpdateDomainInfo)
	w.PUT("/UserInfo/:mail/:key/:value", UpdateUserInfo)
	w.PUT("/SiteInfo", AddSiteInfo)
	w.GET("/SiteInfo", GetSiteInfo)
	w.DELETE("/SiteInfo", DeleteSiteInfo)
	w.PUT("/Certificate", AddCertificateInfo)
	w.GET("/Certificate", GetCertificateInfo)
	w.POST("/Certificate", UpdateCertificateInfo)
	w.DELETE("/Certificate/:id", DeleteCertificateInfo)
	w.PUT("/link/Certificate/:ServerID/:CerID", LinkServerCer)
	w.DELETE("/link/Certificate/:ServerID/:CerID", UnLinkServerCer)
	w.PUT("/link/Site/:ServerID/:SiteID", LinkServerSite)
	w.DELETE("/link/Site/:ServerID/:SiteID", UnLinkServerSite)
	w.POST("/backup", SetBackupFile)
	w.GET("/SiteInfo/check", GetCertificatesInfo)
	w.POST("/Setting/:name/:value", SetSetting)
	w.GET("/Setting/:name", GetSetting)
	w.GET("/DockerInfo/:id", GetDockerInfo)
	w.PUT("/DockerInfo", AddDockerInfo)
	w.PUT("/DockerInfo/:id", EditDockerInfo)
	e.GET("/web/:token/backup", GetBackupFile)
	user := w.Group("/UserInfo")
	user.PUT("/level/:uid/:level", ChangeLevel)
	user.GET("/Password/:oldpass/:newpass", ReSetPassword)
	user.GET("/Token", GetUserToken)
	user.PUT("/Token", GetNewToken)
	user.GET("", GetUserInfo)
	user.GET("/", GetUserList)

	if conf.TLS {
		e.Logger.Fatal(e.StartTLS(":"+conf.ListenPort, conf.CERTPath, conf.KEYPath))
	} else {
		e.Logger.Fatal(e.Start(":" + conf.ListenPort))
	}
}

func checkOnlineI() {
	for true {
		CheckOnline()
	}
}
