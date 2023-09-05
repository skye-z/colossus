package service

import (
	"fmt"
	"log"
	"os"
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
	Hide bool   `json:"hide"`
}

func (fs FileService) GetHomePath(ctx *gin.Context) {
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
	result := sftp.RunShell(CMD_GET_HOME)
	if result == "" || result == "ERROR" {
		common.ReturnMessage(ctx, false, "寻址不可用")
		return
	}
	common.ReturnMessage(ctx, true, strings.TrimSpace(result))
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
	// 是否显示全部
	showAll := "a"
	if param.Hide {
		showAll = ""
	}
	// 执行查询
	result := sftp.RunShell(fmt.Sprintf(CMD_GET_FILE_LIST, showAll, fs.cleanPath(param.Path)))
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
			if len(metas) == 10 {
				fileName = metas[9]
			} else {
				for x := 9; x < len(metas); x++ {
					fileName += metas[x] + " "
				}
				fileName = fileName[0 : len(fileName)-1]
			}
			suffix := fileName[len(fileName)-1]
			var fileType int
			switch suffix {
			case '/':
				fileType = 2
				fileName = fileName[0 : len(fileName)-1]
			case '*':
				fileType = 3
				fileName = fileName[0 : len(fileName)-1]
			case '@':
				fileType = 4
				fileName = fileName[0 : len(fileName)-1]
			default:
				fileType = 1
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

// 获取文件信息
func (fs FileService) GetFileInfo(ctx *gin.Context) {
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
	result := sftp.RunShell(fmt.Sprintf(CMD_GET_FILE_INFO, fs.cleanPath(param.Path)))
	if result == "" || result == "ERROR" {
		common.ReturnMessage(ctx, false, "目录地址不可用")
		return
	}
	results := strings.Split(result, " ")

	common.ReturnData(ctx, true, results)
}

type EditParam struct {
	Id         int64  `json:"id"`
	Model      string `json:"model"`
	FileName   string `json:"fileName"`
	LocalPath  string `json:"localPath"`
	ServerPath string `json:"serverPath"`
}

// 下载文件
func (fs FileService) DownloadFile(ctx *gin.Context) {
	var param EditParam
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
	switch param.Model {
	case "zip":
		now := time.Now()
		result := sftp.RunShell(fmt.Sprintf(CMD_ZIP_FILE, param.ServerPath, now.Unix(), param.FileName))
		if result == "ERROR" {
			common.ReturnMessage(ctx, false, "压缩出错")
			return
		}
		zipName := fmt.Sprintf("%v.tar.gz", now.Unix())
		sftp.Download(param.LocalPath, param.ServerPath, zipName)
		result = sftp.RunShell(fmt.Sprintf(CMD_RM_FILE, fs.cleanPath(param.ServerPath+"/"+zipName)))
		if result == "ERROR" {
			common.ReturnMessage(ctx, false, "残留文件删除出错")
			return
		}
		autoUnzip := common.GetBool("download.auto_unzip")
		if autoUnzip {
			err := common.UnzipTarGz(param.LocalPath, zipName)
			if err != nil {
				log.Println(err)
				common.ReturnMessage(ctx, false, "自动解压出错")
				return
			}
			os.Remove(param.LocalPath + "/" + zipName)
		}
	default:
		log.Println(param.LocalPath, param.ServerPath, param.FileName)
		sftp.Download(param.LocalPath, param.ServerPath, param.FileName)
	}

	common.ReturnMessage(ctx, true, "下载完成")
}

// 上传文件
func (fs FileService) UploadFile(ctx *gin.Context) {
	var param EditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}

	switch param.Model {
	case "directory":
		// 压缩目录
		now := time.Now()
		zipName := fmt.Sprintf("%v.tar.gz", now.Unix())
		err := common.ZipTarGz(param.LocalPath, param.FileName, zipName)
		if err != nil {
			log.Println(err)
			common.ReturnMessage(ctx, false, "目录压缩出错")
			return
		}
		// 上传压缩包
		sftp.Upload(param.LocalPath, param.ServerPath, zipName)
		// 解压目录
		result := sftp.RunShell(fmt.Sprintf(CMD_UNZIP_FILE, zipName, fs.cleanPath(param.ServerPath)))
		if result == "ERROR" {
			common.ReturnMessage(ctx, false, "解压文件出错")
			return
		}
		// 清理垃圾
		result = sftp.RunShell(fmt.Sprintf(CMD_RM_FILE, fs.cleanPath(param.ServerPath+"/"+zipName)))
		if result == "ERROR" {
			common.ReturnMessage(ctx, false, "残留文件删除出错")
			return
		}
		os.Remove(param.LocalPath + "/" + zipName)
	default:
		sftp.Upload(param.LocalPath, param.ServerPath, param.FileName)
	}

	common.ReturnMessage(ctx, true, "上传完成")
}

// 移动文件
func (fs FileService) MoveFile(ctx *gin.Context) {
	var param EditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}

	result := sftp.RunShell(fmt.Sprintf(CMD_MV_FILE, fs.cleanPath(param.LocalPath), fs.cleanPath(param.ServerPath)))
	if result == "ERROR" {
		common.ReturnMessage(ctx, false, "重命名出错")
		return
	}
	common.ReturnMessage(ctx, true, "操作成功")
}

// 删除文件
func (fs FileService) RemoveFile(ctx *gin.Context) {
	var param EditParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		common.ReturnMessage(ctx, false, "传入参数非法")
		return
	}

	// 获取SFTP操作对象
	sftp := fs.getSFTP(param.Id, ctx)
	if sftp == nil {
		return
	}

	result := sftp.RunShell(fmt.Sprintf(CMD_RM_FILE, fs.cleanPath(param.ServerPath+"/"+param.FileName)))
	if result == "ERROR" {
		common.ReturnMessage(ctx, false, "删除出错")
		return
	}
	common.ReturnMessage(ctx, true, "操作成功")
}

// 获取SFTP
func (fs FileService) getSFTP(hostId int64, ctx *gin.Context) *SFTPService {
	// 获取主机信息
	host := &model.Host{Id: hostId}
	fs.HostModel.GetItem(host)
	if len(host.Address) == 0 {
		common.ReturnMessage(ctx, false, "主机不存在")
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

func (fs FileService) cleanPath(path string) string {
	cache := strings.Replace(path, "//", "/", -1)
	cache = strings.Replace(path, " ", "\\ ", -1)
	return cache
}
