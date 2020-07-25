package apis

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AddCertificateInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		certificate := model.Certificate{}
		if err := c.Bind(&certificate); err != nil {
			return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
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
			return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
		}
		certificate.DNSNames = x509Cert.DNSNames[0]
		certificate.Issuer = x509Cert.Issuer.String()
		certificate.IssuingCertificateURL = x509Cert.IssuingCertificateURL[0]
		certificate.NotAfter = x509Cert.NotAfter.Unix()
		certificate.NotBefore = x509Cert.NotBefore.Unix()
		certificate.OCSPServer = x509Cert.OCSPServer[0]
		certificate.Subject = x509Cert.Subject.String()
		certificate.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		if !database.AddCer(certificate) {
			return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
		}
		database.AddLog("Certificate", "addCertificateInfo:{user:{mail:"+user.Mail+"}}", 1)
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func DeleteCertificateInfo(c echo.Context) error {
	user := CheckAuth(c)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if user.Level <= 1 {
		uid := database.GetUser(model.User{Mail: user.Mail})[0].ID
		if database.DelCer(model.Certificate{ID: id, UID: uid}) {
			if database.UnLinkServer(model.ServerLink{CertificateID: id}) {
				database.AddLog("Certificate", "deleteCertificateInfo:{user:{mail:"+user.Mail+"},certificate:{id:"+strconv.FormatInt(id, 10)+"}}", 1)
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UpdateCertificateInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		certificate := model.Certificate{}
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
		if database.UpdateCer(model.Certificate{ID: certificate.ID, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID}, certificate) {
			database.AddLog("Certificate", "updateCertificateInfo:{user:{mail:"+user.Mail+"},certificate:{id:"+strconv.FormatInt(certificate.ID, 10)+"}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func GetCertificateInfo(c echo.Context) error {
	user := CheckAuth(c)
	cer := model.Certificate{}
	if err := c.Bind(&cer); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		cer.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, database.GetCer(cer))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func LinkServerCer(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		sid := c.Param("ServerID")
		cid := c.Param("CerID")
		Isid, _ := strconv.ParseInt(sid, 10, 64)
		Icid, _ := strconv.ParseInt(cid, 10, 64)
		payload := model.ServerLink{ServerID: Isid, CertificateID: Icid}
		data := database.GetLinkCer(payload)
		if len(data) == 0 {
			if (database.LinkServer(model.ServerLink{ServerID: Isid, CertificateID: Icid})) {
				database.AddLog("Certificate", "linkServerCer:{user:{mail:"+user.Mail+"},link:{serverID:"+strconv.FormatInt(Isid, 10)+",certificateID:"+strconv.FormatInt(Icid, 10)+"}}", 1)
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
		} else {
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "already linked"})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UnLinkServerCer(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		sid := c.Param("ServerID")
		cid := c.Param("CerID")
		Isid, _ := strconv.ParseInt(sid, 10, 64)
		Icid, _ := strconv.ParseInt(cid, 10, 64)
		payload := model.ServerLink{ServerID: Isid, CertificateID: Icid}
		data := database.UnLinkServer(payload)
		if data {
			if len(database.GetLinkCer(payload)) == 0 {
				database.AddLog("Certificate", "unLinkServerCer:{user:{mail:"+user.Mail+"},link:{serverID:"+strconv.FormatInt(Isid, 10)+",certificateID:"+strconv.FormatInt(Icid, 10)+"}}", 1)
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
