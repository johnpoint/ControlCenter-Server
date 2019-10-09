package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Callback struct {
	Code int64
	Info string
}

func main() {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://lvcshu.test.io"},
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
	e.Logger.Fatal(e.Start(":1323"))
}