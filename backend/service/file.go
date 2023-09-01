package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/common"
	"github.com/skye-z/colossus/backend/model"
)

const layout = "2023-01-01 12:00:00.000000000 +0800"

type FileService struct {
	HostModel model.HostModel
}

type FileParam struct {
	Id   int64  `json:"id"`
	Path string `json:"path"`
}

// 获取文件列表
func (fs FileService) GetFileList(ctx *gin.Context) {
	var param FileParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}
	// 执行查询
	result := sftp.RunShell(fmt.Sprintf(CMD_GET_FILE_LIST, param.Path))
	if result == "" || result == "ERROR" {
		common.ReturnMessage(ctx, false, "目录地址不可用")
		return
	}
	results := strings.Split(result, "\n")

	var files []SFTPFile
	for i := 0; i < len(results); i++ {
		if i == 0 {

		} else {
			// 去除前后空格
			cache := strings.TrimSpace(results[i])
			cacheList := strings.Split(cache, " ")

			var metas []string
			for x := 0; x < len(cacheList); x++ {
				value := strings.TrimSpace(cacheList[x])
				if len(value) != 0 {
					metas = append(metas, value)
				}
			}
			if len(metas) == 0 {
				continue
			}

			fileId, _ := strconv.ParseInt(metas[0], 10, 64)
			fileLink, _ := strconv.Atoi(metas[2])
			fileSize, _ := strconv.ParseInt(metas[5], 10, 64)
			fileDate, _ := time.Parse(layout, fmt.Sprintf("%s %s %s", metas[6], metas[7], metas[8]))

			var fileName string
			for x := 9; x < len(metas); x++ {
				fileName += metas[x]
			}

			suffix := fileName[len(fileName)-1]
			var fileType int
			switch suffix {
			case '/':
				fileType = 2
				fileName = fileName[0 : len(fileName)-1]
				break
			case '*':
				fileType = 3
				fileName = fileName[0 : len(fileName)-1]
				break
			case '@':
				fileType = 4
				fileName = fileName[0 : len(fileName)-1]
				break
			default:
				fileType = 1
				break
			}

			file := &SFTPFile{
				Id:    fileId,
				Name:  fileName,
				Type:  fileType,
				Power: metas[1][0 : len(metas[1])-1],
				Link:  fileLink,
				User:  metas[3],
				Group: metas[4],
				Size:  fileSize,
				Date:  fileDate.Unix(),
			}
			files = append(files, *file)
		}
	}

	common.ReturnData(ctx, true, files)
}

type UpDownParam struct {
	Id         int64  `json:"id"`
	LocalPath  string `json:"localPath"`
	ServerPath string `json:"serverPath"`
}

func (fs FileService) DownloadFile(ctx *gin.Context) {
	var param UpDownParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}

	// 下载文件
	sftp.Download(param.LocalPath, param.ServerPath)

	common.ReturnMessage(ctx, true, "下载开始")
}

func (fs FileService) UploadFile(ctx *gin.Context) {
	var param UpDownParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}

	// 上传文件
	sftp.Upload(param.LocalPath, param.ServerPath)

	common.ReturnMessage(ctx, true, "上传开始")
}

// 获取SFTP
func (fs FileService) getSFTP(hostId int64, ctx *gin.Context) *SFTPService {
	// 获取主机信息
	host := &model.Host{Id: hostId}
	fs.HostModel.GetItem(host)
	if len(host.Address) == 0 {
		common.ReturnMessage(ctx, false, "主机地址为空")
		return nil
	}
	// 组装SSH连接配置
	sshConfig := SSHService{
		Address:  host.Address,
		Port:     host.Port,
		AuthType: host.AuthType,
		User:     host.User,
		Secret:   host.Secret,
	}
	// 创建SSH客户端
	sshClient, err := sshConfig.CreateClient()
	if err != nil {
		common.ReturnMessage(ctx, false, "无法连接主机")
		return nil
	}
	// 组装SFTP连接配置
	sftpConfig := SFTPService{
		sshClient: sshClient,
	}
	// 创建SFTP客户端
	sftpConfig.CreateClient()
	return &sftpConfig
}
