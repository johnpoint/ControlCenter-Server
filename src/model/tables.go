package model

import "time"

// Server model of server
type Server struct {
	Status   string `json:"status" xml:"status" form:"status" query:"status"`
	Hostname string `json:"hostname" xml:"hostname" form:"hostname" query:"hostname"`
	Ipv4     string `json:"ipv4" xml:"ipv4" form:"ipv4" query:"ipv4"`
	Ipv6     string `json:"ipv6" xml:"ipv6" form:"ipv6" query:"ipv6"`
	ID       int64  `gorm:"AUTO_INCREMENT"`
	UID      int64  `json:"uid" xml:"uid" form:"uid" query:"uid"`
	Token    string
	Online   int64 `gorm:"default:1"`
}

// User model of user
type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64  //用户等级 0 = 特权用户 1 = 普通用户 2 = 游客
	Token    string
}

// Certificate model of Certificate
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
	UID                   int64  `json:"uid" xml:"uid" form:"uid" query:"uid"`
}

type SysConfig struct {
	ID    int64 `gorm:"AUTO_INCREMENT"`
	UID   int64
	Name  string
	Value string
}

// ServerCertificate model
type ServerLink struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	Type     string
	ServerID int64
	ItemID   int64
}

type LogInfo struct {
	ID        int64 `gorm:"AUTO_INCREMENT"`
	Service   string
	Info      string
	Level     int64 // 1 Info | 2 Warn | 3 Error
	CreatedAt time.Time
}

type Event struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	Type     int64 // 1 client | 2 docker
	TargetID int64
	Code     int64
	Info     string
	Active   int64
}

type Configuration struct {
	ID    int64  `gorm:"AUTO_INCREMENT"`
	Type  string `json:"type" xml:"type" form:"type" query:"type"`
	Value string `json:"value" xml:"value" form:"value" query:"value"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Path  string `json:"path" xml:"path" form:"path" query:"path"`
	UID   int64
}
