package main

import (
	"github.com/labstack/echo"
)

type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64
}

type Server struct {
	Hostname string
	Ipv4     string
	Ipv6     string
	ID       int64 `gorm:"AUTO_INCREMENT"`
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
	e := echo.New() /*
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, get")
		})

		e.POST("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, post")
		})
		e.GET("/show", show)*/
	e.POST("/auth/login", oaLogin)
	e.POST("/auth/register", oaRegister)
	/*
		type User struct {
			Name  string `json:"name" xml:"name" form:"name" query:"name"`
			Email string `json:"email" xml:"email" form:"email" query:"email"`
		}

		e.POST("/users", func(c echo.Context) error {
			u := new(User)
			if err := c.Bind(u); err != nil {
				return err
			}
			return c.JSON(http.StatusCreated, u)
		})

		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == "johnpoint" && password == "200632482" {
				return true, nil
			}
			return false, nil
		}))
	*/
	e.Logger.Fatal(e.Start(":1323"))
}
