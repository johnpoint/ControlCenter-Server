package main

// Callback send some info to client
type Callback struct {
	Code int64
	Info string
}

// Server model of server
type Server struct {
	Status   string `json:"status" xml:"status" form:"status" query:"status"`
	Hostname string `json:"hostname" xml:"hostname" form:"hostname" query:"hostname"`
	Ipv4     string `json:"ipv4" xml:"ipv4" form:"ipv4" query:"ipv4"`
	Ipv6     string `json:"ipv6" xml:"ipv6" form:"ipv6" query:"ipv6"`
	ID       int64  `gorm:"AUTO_INCREMENT"`
	UID      int64  `json:"uid" xml:"uid" form:"uid" query:"uid"`
	Token    string
	Online   int64 `gorm:"default:true"`
}

// Service TODO
type Service struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	Name     string
	Enable   string
	Disable  string
	Status   int64
	Serverid int64
	UID      int64
}

// Site model of Site
type Site struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Config string `json:"config" xml:"config" form:"config" query:"config"`
	Cer    int64  `json:"cer" xml:"cer" form:"cer" query:"cer"`
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

// Config model of config
type Config struct {
	AllowAddress []string
	ListenPort   string
	TLS          bool
	CERTPath     string
	KEYPath      string
	Salt         string
	Database     string
}

// SysConfig model of config
type SysConfig struct {
	ID    int64 `gorm:"AUTO_INCREMENT"`
	UID   int64
	Name  string
	Value string
}

// Domain model of domain
type Domain struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Cer    string `json:"cer" xml:"cer" form:"cer" query:"cer"`
	Key    string `json:"key" xml:"key" form:"key" query:"key"`
}

// ServerCertificate model
type ServerLink struct {
	ID            int64 `gorm:"AUTO_INCREMENT"`
	CertificateID int64
	ServerID      int64
	SiteID        int64
}

// UpdateInfo model
type UpdateInfo struct {
	Code         int64
	Sites        []DataSite
	Certificates []DataCertificate
	Services     []DataService
}

// DataSite model
type DataSite struct {
	ID     int64
	Domain string
	CerID  int64
	Config string
}

// DataService model
type DataService struct {
	Name    string
	Enable  string
	Disable string
	Status  string
}

// DataCertificate model
type DataCertificate struct {
	ID        int64
	Domain    string
	FullChain string
	Key       string
}
