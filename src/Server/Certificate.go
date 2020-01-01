package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/labstack/echo"
)

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
	user := checkAuth(c)
	cer := Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user.Level == 1 {
		if delCer(cer) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateCertificateInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
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
	user := checkAuth(c)
	cer := Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user.Level == 1 {
		return c.JSON(http.StatusOK, getCer(cer))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}
