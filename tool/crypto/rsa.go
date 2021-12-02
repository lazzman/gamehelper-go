package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/wenzhenxi/gorsa"
)

//GenerateRsaKeyPair 生成pem格式的rsa公私密钥
func GenerateRsaKeyPair() (privateKeyPem string, publicKeyPem string) {
	// 生成 RSA 密钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	// 通过x509标准将得到的ras私钥序列化为PKCS #1、ASN.1 DER 形式
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	// 通过x509标准将得到的ras公钥序列化为PKIX、ASN.1 DER 形式
	x509PublicKey, _ := x509.MarshalPKIXPublicKey(publicKey)

	priBlock := &pem.Block{
		Type:  "RSA Private Key",
		Bytes: x509PrivateKey,
	}
	pubBlock := &pem.Block{
		Type:  "RSA Public Key",
		Bytes: x509PublicKey,
	}
	priKey := pem.EncodeToMemory(priBlock)
	pubKey := pem.EncodeToMemory(pubBlock)

	return string(priKey), string(pubKey)
}

//EncryptRsaPriKey 使用rsa私钥加密
func EncryptRsaPriKey(content string, privateKeyPem string) (string, error) {
	return gorsa.PriKeyEncrypt(content, privateKeyPem)
}

//DecryptRsaPubKey 使用rsa公钥解密
func DecryptRsaPubKey(encryptContent string, publicKeyPem string) (string, error) {
	return gorsa.PublicDecrypt(encryptContent, publicKeyPem)
}
