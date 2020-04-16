package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "init" {
			initServer()
		} else if os.Args[1] == "start" {
			start()
		}
	}
	fmt.Println("参数错误")
}

func start() {
	go checkOnlineI()
	conf := loadConfig()
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowAddress,
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	u := e.Group("/user")
	u.POST("/auth/login", oaLogin)
	u.POST("/auth/register", oaRegister)

	s := e.Group("/server")
	s.POST("/setup/:token", setupServer)
	s.POST("/update/:token", serverUpdate)
	s.GET("/update/:token", getServerUpdate)
	s.GET("/Certificate/:token/:id", serverGetCertificate)

	e.GET("/", accessible)

	sys := e.Group("/system")
	sys.POST("/restart", sysRestart)
	jwtConfig := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(conf.Salt),
	}
	w := e.Group("/web")
	w.Use(middleware.JWTWithConfig(jwtConfig))
	w.POST("/debug/check", checkPower)
	w.GET("/ServerInfo", getServerInfo)
	w.PUT("/ServerInfo", updateServerInfo)
	w.GET("/ServerInfo/Certificate", getCertificateLinked)
	w.GET("/ServerInfo/Site", getSiteLinked)
	w.DELETE("/Server/:id", removeServer)
	w.GET("/DomainInfo", getDomainInfo)
	w.PUT("/DomainInfo", updateDomainInfo)
	w.PUT("/UserInfo/:mail/:key/:value", updateUserInfo)
	w.PUT("/SiteInfo", addSiteInfo)
	w.GET("/SiteInfo", getSiteInfo)
	w.DELETE("/SiteInfo", deleteSiteInfo)
	w.PUT("/Certificate", addCertificateInfo)
	w.GET("/Certificate", getCertificateInfo)
	w.POST("/Certificate", updateCertificateInfo)
	w.DELETE("/Certificate/:id", deleteCertificateInfo)
	w.PUT("/link/Certificate/:ServerID/:CerID", linkServerCer)
	w.DELETE("/link/Certificate/:ServerID/:CerID", unLinkServerCer)
	w.PUT("/link/Site/:ServerID/:SiteID", linkServerSite)
	w.DELETE("/link/Site/:ServerID/:SiteID", unLinkServerSite)
	w.POST("/backup", setBackupFile)
	w.GET("/SiteInfo/check", getCertificatesInfo)
	w.POST("/Setting/:name/:value", setSetting)
	w.GET("/Setting/:name", getSetting)
	w.GET("/DockerInfo/:id", getDockerInfo)
	w.PUT("/DockerInfo", addDockerInfo)
	w.PUT("/DockerInfo/:id", editDockerInfo)
	e.GET("/web/:token/backup", getBackupFile)
	user := w.Group("/UserInfo")
	user.PUT("/level/:uid/:level", changeLevel)
	user.GET("/Password/:oldpass/:newpass", reSetPassword)
	user.GET("/Token", getUserToken)
	user.PUT("/Token", getNewToken)
	user.GET("", getUserInfo)
	user.GET("/", getUserList)

	if conf.TLS {
		e.Logger.Fatal(e.StartTLS(":"+conf.ListenPort, conf.CERTPath, conf.KEYPath))
	} else {
		e.Logger.Fatal(e.Start(":" + conf.ListenPort))
	}
}

func accessible(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>ControlCenter</h1>(´・ω・`) 运行正常<br><hr>Ver: 1.7.4")
}

func checkOnlineI() {
	for true {
		checkOnline()
	}
}
