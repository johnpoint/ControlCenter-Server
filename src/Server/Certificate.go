package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func addCertificateInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
		certificate := Certificate{}
		if err := c.Bind(&certificate); err != nil {
			return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
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
		}
		certificate.DNSNames = x509Cert.DNSNames[0]
		certificate.Issuer = x509Cert.Issuer.String()
		certificate.IssuingCertificateURL = x509Cert.IssuingCertificateURL[0]
		certificate.NotAfter = x509Cert.NotAfter.Unix()
		certificate.NotBefore = x509Cert.NotBefore.Unix()
		certificate.OCSPServer = x509Cert.OCSPServer[0]
		certificate.Subject = x509Cert.Subject.String()
		certificate.UID = getUser(User{Mail: user.Mail})[0].ID
		if !addCer(certificate) {
			return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func deleteCertificateInfo(c echo.Context) error {
	user := checkAuth(c)
	cer := Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user.Level == 1 {
		cer.UID = getUser(User{Mail: user.Mail})[0].ID
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
		if updateCer(Certificate{ID: certificate.ID, UID: getUser(User{Mail: user.Mail})[0].ID}, certificate) {
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
		cer.UID = getUser(User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, getCer(cer))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func linkServerCer(c echo.Context) error {
	//user := checkAuth(c)
	sid := c.Param("ServerID")
	cid := c.Param("CerID")
	Isid, _ := strconv.ParseInt(sid, 10, 64)
	Icid, _ := strconv.ParseInt(cid, 10, 64)
	payload := ServerCertificate{ServerID: Isid, CertificateID: Icid}
	data := getLinkCer(payload)
	if len(data) == 0 {
		if (LinkCer(ServerCertificate{ServerID: Isid, CertificateID: Icid})) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
}
