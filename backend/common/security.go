package common

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"

	"crypto/aes"
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

// AES加密
func EncryptAES(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainText)

	return ciphertext, nil
}

// AES解密
func DecryptAES(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, ciphertext)

	return plainText, nil
}

type HostCert struct {
	Type       string `json:"type"`
	Level      int16  `json:"level"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

// 生成密钥
func GenerateHostCert(level int) *HostCert {
	// 生成RSA私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, level)
	if err != nil {
		log.Println("[Cert]", err)
		return nil
	}
	// 获取RSA公钥
	publicKey := privateKey.PublicKey

	sshPub, err := ssh.NewPublicKey(&publicKey)
	if err != nil {
		log.Println("[Cert]", err)
		return nil
	}
	sshRsaBytes := ssh.MarshalAuthorizedKey(sshPub)

	// 将私钥和公钥转为字节切片
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	// 将私钥和公钥编码为PEM格式
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	hostCert := &HostCert{
		Type:       "RSA",
		Level:      2048,
		PrivateKey: string(privateKeyPEM),
		PublicKey:  string(sshRsaBytes),
	}
	return hostCert
}
