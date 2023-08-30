package service

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	// 获取文件列表(详细)
	// 文件索引号  权限  链接数  所有者  用户组 文件大小  修改日期  修改时间  时区  文件名
	// 103921544 -rw-------.  1 root root    59 2023-04-04 16:13:07.439361337 +0800 .Xauthority
	// 文件名中: ‘/’表示目录、‘@’表示链接、‘*’表示可执行
	CMD_GET_FILE_LIST = "ls -aliF --full-time %s"
)

type SFTPService struct {
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

type SFTPFile struct {
	// 索引
	Id int64 `json:"id"`
	// 文件名称
	Name string `json:"name"`
	// 文件类型(1文件、2目录、3可执行程序、4链接)
	Type int `json:"type"`
	// 权限
	Power string `json:"power"`
	// 链接数
	Link int `json:"link"`
	// 所有者
	User string `json:"user"`
	// 用户组
	Group string `json:"group"`
	// 文件大小
	Size int64 `json:"size"`
	// 修改时间
	Date int64 `json:"date"`
}

// 创建客户端
func (s *SFTPService) CreateClient() {
	client, err := sftp.NewClient(s.sshClient)
	if err != nil {
		log.Fatalln("SFTP创建失败")
	}
	s.sftpClient = client
}

// 关闭客户端
func (s *SFTPService) CloseClient() {
	s.RunShell("exit")
	s.sftpClient.Close()
	s.sshClient.Close()
}

// 执行命令
func (s *SFTPService) RunShell(shell string) string {
	log.Println("Run:", shell)
	var (
		session *ssh.Session
		err     error
	)
	if session, err = s.sshClient.NewSession(); err != nil {
		log.Println("Shell error:", err)
		return "ERROR"
	}
	if output, err := session.CombinedOutput(shell); err != nil {
		log.Println("Shell error:", err)
		return "ERROR"
	} else {
		return string(output)
	}
}

// 上传文件
func (s *SFTPService) Upload(localPath, cloudPath string) {
	localFile, _ := os.Open(localPath)
	cloudFile, _ := s.sftpClient.Create(cloudPath)
	defer func() {
		_ = localFile.Close()
		_ = cloudFile.Close()
	}()
	buf := make([]byte, 1024)
	for {
		n, err := localFile.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("error occurred:", err)
			} else {
				break
			}
		}
		_, _ = cloudFile.Write(buf[:n])
	}
}

// 下载文件
func (s *SFTPService) Download(localPath, cloudPath string) {
	localFile, _ := s.sftpClient.Open(localPath)
	cloudFile, _ := os.Create(cloudPath)
	defer func() {
		_ = localFile.Close()
		_ = cloudFile.Close()
	}()
	if _, err := localFile.WriteTo(cloudFile); err != nil {
		log.Fatalln("error occurred", err)
	}
	fmt.Println("文件下载完毕")
}
