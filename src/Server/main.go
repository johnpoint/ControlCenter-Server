package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Callback struct {
	Code int64
	Info string
}

func main() {
	if len(os.Args) <= 2 {
		if os.Args[1] == "init" {
			initServer()
		} else if os.Args[1] == "start" {
			start()
		}
	}
	fmt.Println("参数错误")
}

func start() {
	conf := loadConfig()
	e := echo.New()

	if conf.TLS {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(conf.HostAddress)
		e.AutoTLSManager.Cache = autocert.DirCache(conf.CachePath)
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
	}

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
	s.POST("/setup", setupServer)
	s.POST("/update/:token", serverUpdate)
	s.GET("/Certificate/:token/:id", serverGetCertificate)

	e.GET("/", accessible)

	w := e.Group("/web")
	w.Use(middleware.JWT([]byte("NFUCA")))
	w.POST("debug/check", checkPower)
	w.GET("/ServerInfo", getServerInfo)
	w.GET("/DomainInfo", getDomainInfo)
	w.PUT("/DomainInfo", updateDomainInfo)
	w.PUT("/ServerInfo", updateServerInfo)
	w.GET("/UserInfo", getUserInfo)
	w.PUT("/SiteInfo", addSiteInfo)
	w.GET("/SiteInfo", getSiteInfo)
	w.PUT("/Certificate", addCertificateInfo)
	w.GET("/Certificate", getCertificateInfo)
	w.POST("/Certificate", updateCertificateInfo)
	w.POST("/rmCertificate", deleteCertificateInfo)
	w.GET("/Service", getService)

	if conf.TLS {
		e.Logger.Fatal(e.StartAutoTLS(":" + conf.ListenPort))
	} else {
		e.Logger.Fatal(e.Start(":" + conf.ListenPort))
	}
}
