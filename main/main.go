package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64
}

type Server struct {
	Token    string
	Status   string `json:"status" xml:"status" form:"status" query:"status"`
	Hostname string `json:"hostname" xml:"hostname" form:"hostname" query:"hostname"`
	Ipv4     string `json:"ipv4" xml:"ipv4" form:"ipv4" query:"ipv4"`
	Ipv6     string `json:"ipv6" xml:"ipv6" form:"ipv6" query:"ipv6"`
	ID       int64  `gorm:"AUTO_INCREMENT"`
}

type ServerKey struct {
	ID      int64  `json:"id" xml:"id" form:"id" query:"id" gorm:"AUTO_INCREMENT"`
	Public  string `json:"public" xml:"public" form:"public" query:"public"`
	Private string `json:"private" xml:"private" form:"private" query:"private"`
}

type Site struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Server int64  `json:"server" xml:"server" form:"server" query:"server"`
	Status int64  `json:"status" xml:"status" form:"status" query:"status"`
	Config string `json:"config" xml:"config" form:"config" query:"config"`
	Cer    int64  `json:"cer" xml:"cer" form:"cer" query:"cer"`
}

type Certificate struct {
	ID                    int64  `json:"id" xml:"id" form:"id" query:"id" gorm:"AUTO_INCREMENT"`
	Name                  string `json:"name" xml:"name" form:"name" query:"name"`
	Fullchain             string `json:"fullchain" xml:"fullchain" form:"fullchain" query:"fullchain"`
	Key                   string `json:"key" xml:"key" form:"key" query:"key"`
	DNSNames              string `json:"DNSNames" xml:"DNSNames" form:"DNSNames" query:"DNSNames"`
	Issuer                string `json:"Issuer" xml:"Issuer" form:"Issuer" query:"Issuer"`
	IssuingCertificateURL string `json:"IssuingCertificateURL" xml:"IssuingCertificateURL" form:"IssuingCertificateURL" query:"IssuingCertificateURL"`
	NotAfter              int64  `json:"NotAfter" xml:"NotAfter" form:"NotAfter" query:"NotAfter"`
	NotBefore             int64  `json:"NotBefore" xml:"NotBefore" form:"NotBefore" query:"NotBefore"`
	OCSPServer            string `json:"OCSPServer" xml:"OCSPServer" form:"OCSPServer" query:"OCSPServer"`
	Subject               string `json:"Subject" xml:"Subject" form:"Subject" query:"Subject"`
}

type Service struct {
	ID     int64 `gorm:"AUTO_INCREMENT"`
	Server int64
	Site   int64
}

type Domain struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Cer    string `json:"cer" xml:"cer" form:"cer" query:"cer"`
	Key    string `json:"key" xml:"key" form:"key" query:"key"`
}

type Config struct {
	ID          int64 `gorm:"AUTO_INCREMENT"`
	configKey   string
	configValue string
}

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
