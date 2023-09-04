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
	CMD_GET_HOME = "cd ~ && pwd"
	// 获取文件列表
	// 文件索引号  权限  链接数  所有者  用户组 文件大小  修改日期  修改时间  时区  文件名
	// 103921544 -rw-------.  1 root root    59 2023-04-04 16:13:07.439361337 +0800 .Xauthority
	// 文件名中: ‘/’表示目录、‘@’表示链接、‘*’表示可执行
	CMD_GET_FILE_LIST = "ls -%sliF --group-directories-first --full-time %s"
	// 获取文件详情
	// 文件名 索引节号 文件大小 I/O块大小 文件占用的块数 块大小 硬链接数量 文件类型 所有者id 所有者 用户组id 用户组 权限位 权限位和文件类型 主要设备类型 次要设备类型 最后访问时间 最后修改时间 最后更改时间
	CMD_GET_FILE_INFO = "stat --format=\"%%n %%i %%s %%o %%b %%B %%H %%f %%u %%U %%g %%G %%a %%A %%t %%T %%X %%Y %%Z\" %s"
	// 移动文件 重命名文件
	CMD_MV_FILE = "mv %s %s"
	// 压缩文件
	CMD_ZIP_FILE = "cd %s && tar -zcvf %v.tar.gz ./%s"
	// 删除文件
	CMD_RM_FILE = "rm -rf --preserve-root %s"
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
func (s *SFTPService) Upload(localPath, cloudPath, fileName string) {
	log.Println(localPath, "->", cloudPath)
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
func (s *SFTPService) Download(localPath, cloudPath, fileName string) {
	// 判断应用目录是否存在
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		// 目录不存在,创建目录
		os.Mkdir(localPath, os.ModePerm)
	}
	log.Println(cloudPath + "/" + fileName)
	cloudFile, err := s.sftpClient.Open(cloudPath + "/" + fileName)
	if err != nil {
		log.Println("cloudFile error", err)
		return
	}
	defer cloudFile.Close()
	localFile, err := os.Create(localPath + "/" + fileName)
	if err != nil {
		log.Println("localFile error", err)
		return
	}
	defer localFile.Close()
	number, err := io.Copy(localFile, cloudFile)
	if err != nil {
		log.Println("Copy occurred", err)
		return
	}
	fmt.Printf("Downloaded %d bytes\n", number)
}
