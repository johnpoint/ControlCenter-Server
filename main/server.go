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
	Hostname string `json:"hostname" xml:"hostname" form:"hostname" query:"hostname"`
	Ipv4     string `json:"ipv4" xml:"ipv4" form:"ipv4" query:"ipv4"`
	Ipv6     string `json:"ipv6" xml:"ipv6" form:"ipv6" query:"ipv6"`
	ID       int64  `gorm:"AUTO_INCREMENT"`
}

type Service struct {
	ID   int64 `gorm:"AUTO_INCREMENT"`
	Name string
}

type Domain struct {
	ID     int64 `gorm:"AUTO_INCREMENT"`
	Name   string
	Status string
	Cer    string
	Key    string
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
	e.POST("/auth/login", oaLogin)
	e.POST("/auth/register", oaRegister)
	e.POST("/server/setup", setupServer)
	e.GET("/", accessible)
	r := e.Group("/")
	r.Use(middleware.JWT([]byte("NFUCA")))
	r.POST("debug/check", checkPower)
	r.POST("web/getUpdate", getServerInfo)
	e.Logger.Fatal(e.Start(":1323"))
}
