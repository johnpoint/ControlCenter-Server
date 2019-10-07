package main

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getServerInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getServer(server))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getSiteInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user != nil {
		site := Site{}
		if err := c.Bind(&site); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, getSite(site))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func setupServer(c echo.Context) error {
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	check := getServer(Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "Server already exists"})
	}
	time := time.Now().Unix()
	data := []byte(strconv.FormatInt(time, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !addServer(server) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: md5str1})
}

func addSiteInfo(c echo.Context) error {
	site := Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !addSite(site) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}

func addCertificateInfo(c echo.Context) error {
	certificate := Certificate{}
	if err := c.Bind(&certificate); err != nil {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
		panic(err)
	}
	var certPEMBlock []byte = []byte(certificate.Fullchain)
	var cert tls.Certificate
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain != nil {
		cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
	}
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
		panic(c)
	}
	certificate.DNSNames = x509Cert.DNSNames[0]
	certificate.Issuer = x509Cert.Issuer.String()
	certificate.IssuingCertificateURL = x509Cert.IssuingCertificateURL[0]
	certificate.NotAfter = x509Cert.NotAfter.Unix()
	certificate.NotBefore = x509Cert.NotBefore.Unix()
	certificate.OCSPServer = x509Cert.OCSPServer[0]
	certificate.Subject = x509Cert.Subject.String()
	if !addCer(certificate) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}

func deleteCertificateInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	cer := Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user["level"].(float64) == 1 {
		if delCer(cer) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateCertificateInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		certificate := Certificate{}
		if err := c.Bind(&certificate); err != nil {
			panic(err)
		}
		var certPEMBlock []byte = []byte(certificate.Fullchain)
		var cert tls.Certificate
		certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
		cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
		certDERBlockChain, _ := pem.Decode(restPEMBlock)
		if certDERBlockChain != nil {
			cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
		}
		x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
		if err != nil {
			panic(err)
		}
		certificate.DNSNames = x509Cert.DNSNames[0]
		certificate.Issuer = x509Cert.Issuer.String()
		certificate.IssuingCertificateURL = x509Cert.IssuingCertificateURL[0]
		certificate.NotAfter = x509Cert.NotAfter.Unix()
		certificate.NotBefore = x509Cert.NotBefore.Unix()
		certificate.OCSPServer = x509Cert.OCSPServer[0]
		certificate.Subject = x509Cert.Subject.String()
		if updateCer(Certificate{ID: certificate.ID}, certificate) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getCertificateInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	cer := Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getCer(cer))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getDomainInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getDomain(Domain{}))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateDomainInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		domain := Domain{}
		if err := c.Bind(&domain); err != nil {
			panic(err)
		}
		if updateDomain(Domain{Name: domain.Name}, domain) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateServerInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		server := Server{}
		if err := c.Bind(&server); err != nil {
			panic(err)
		}
		if updateServer(Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getUserInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}
