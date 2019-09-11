package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	//"log"
)

func main() {
	var cert tls.Certificate
	certPEMBlock, err := ioutil.ReadFile("server.crt")
	//获取下一个pem格式证书数据 -----BEGIN CERTIFICATE----- -----END CERTIFICATE-----
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	//附加数字证书到返回
	cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	//继续解析Certifacate Chan,这里要明白证书链的概念
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain != nil {
		//追加证书链证书到返回
		cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
		fmt.Println("存在证书链")
	}
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(x509Cert.DNSNames)
	fmt.Println(x509Cert.Issuer.String())
	fmt.Println(x509Cert.IssuingCertificateURL)
	fmt.Println(x509Cert.NotAfter.UnixNano)
	fmt.Println(x509Cert.NotBefore.UnixNano)
	fmt.Println(x509Cert.OCSPServer[0])
	fmt.Println(x509Cert.Subject.CommonName)
}
