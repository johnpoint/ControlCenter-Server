package main

import (
	// "crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"io/ioutil"
)

func parseCert(crt, privateKey string) *tls.Certificate {
	var cert tls.Certificate
	//加载PEM格式证书到字节数组
	certPEMBlock, err := ioutil.ReadFile(crt)
	if err != nil {
		return nil
	}
	//获取下一个pem格式证书数据 -----BEGIN CERTIFICATE----- -----END CERTIFICATE-----
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil
	}
	//附加数字证书到返回
	cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	//继续解析Certifacate Chan,这里要明白证书链的概念
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain != nil {
		//追加证书链证书到返回
		fmt.Println(certDERBlockChain.Headers)
		cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
		fmt.Println("存在证书链")
	}

	fmt.Println()
	//读取RSA私钥进文件到字节数组
	keyPEMBlock, err := ioutil.ReadFile(privateKey)
	if err != nil {
		return nil
	}

	//解码pem格式的私钥------BEGIN RSA PRIVATE KEY----- -----END RSA PRIVATE KEY-----
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		return nil
	}
	//打印出私钥类型
	fmt.Println(keyDERBlock.Type)
	fmt.Println(keyDERBlock.Headers)
	var key interface{}
	var errParsePK error
	if keyDERBlock.Type == "RSA PRIVATE KEY" {
		//RSA PKCS1
		key, errParsePK = x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	} else if keyDERBlock.Type == "PRIVATE KEY" {
		//pkcs8格式的私钥解析
		key, errParsePK = x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	}

	if errParsePK != nil {
		return nil
	} else {
		cert.PrivateKey = key
	}
	//第一个叶子证书就是我们https中使用的证书
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	fmt.Println(x509Cert.Subject)

	if err != nil {
		fmt.Println("x509证书解析失败")
		return nil
	} else {
		switch x509Cert.PublicKeyAlgorithm {
		case x509.RSA:
			{
				fmt.Println("Plublic Key Algorithm:RSA")
			}
		case x509.DSA:
			{
				fmt.Println("Plublic Key Algorithm:DSA")
			}
		case x509.ECDSA:
			{
				fmt.Println("Plublic Key Algorithm:ECDSA")
			}
		case x509.UnknownPublicKeyAlgorithm:
			{
				fmt.Println("Plublic Key Algorithm:Unknow")
			}
		}
	}
	return &cert
}

func main() {
	fmt.Println("---------pkcs8 private key ---------------")
	fmt.Println((parseCert("./server.crt", "pkcs8_server.key").Certificate)[])

	fmt.Println("---------pkcs1 private key ---------------")
	parseCert("./server.crt", "server.key")
}
