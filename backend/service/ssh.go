package service

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	AUTH_TYPE_PASSWORD    = "pwd"
	AUTH_TYPE_CERTIFICATE = "cert"
)

type SSHService struct {
	Address  string
	Port     int
	AuthType string
	User     string
	Key      string
	Secret   string
}

// 创建客户端
func (s *SSHService) CreateClient() (*ssh.Client, error) {
	// 创建授权方式
	var auth []ssh.AuthMethod
	if s.AuthType == AUTH_TYPE_PASSWORD {
		// 使用密码授权登录
		auth = []ssh.AuthMethod{ssh.Password(s.Secret)}
	} else if s.AuthType == AUTH_TYPE_CERTIFICATE {
		var err error
		var signer ssh.Signer
		if len(s.Secret) == 0 {
			signer, err = ssh.ParsePrivateKey([]byte(s.Key))
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(s.Key), []byte(s.Secret))
		}
		// 使用证书授权登录
		if err != nil {
			log.Fatalln("证书无效")
			return nil, err
		}
		auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	// 创建连接配置
	config := ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            s.User,
		HostKeyCallback: SaveHostKey,
		Auth:            auth,
	}
	// 拨号连接
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Address, s.Port), &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 连接主机
func (s *SSHService) Connect(client *ssh.Client, height int, width int) (*ssh.Session, error) {
	var (
		session *ssh.Session
		err     error
	)
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		// 开启回显(不开自动补全功能就没了)
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", height, width, modes); err != nil {
		return nil, err
	}

	return session, nil
}

// 保存主机密钥
func SaveHostKey(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return nil
}
