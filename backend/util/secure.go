package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/bcrypt"
)

func RandomString(length int) string {
	s := make([]byte, length)
	_, _ = rand.Read(s)
	return string(s)
}

func GenerateRSAKeys() (privateKeyBytes, publicKeyBytes []byte) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	publicKey := &privateKey.PublicKey
	privateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	privateKeyBytes = pem.EncodeToMemory(privateBlock)
	publicKeyBytes = pem.EncodeToMemory(publicBlock)

	return privateKeyBytes, publicKeyBytes
}

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func JWTSign(sigKey []byte) {

}
