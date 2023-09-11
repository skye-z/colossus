package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/common"
	"github.com/skye-z/colossus/backend/model"
)

type CertService struct {
	CertModel model.CertModel
}

// 获取凭证列表
func (cs CertService) GetList(ctx *gin.Context) {
	page := ctx.Query("page")
	iPage, err1 := strconv.Atoi(page)
	num := ctx.Query("number")
	iNum, err2 := strconv.Atoi(num)
	if err1 != nil || err2 != nil {
		common.ReturnError(ctx, common.Errors.ParamIllegalError)
		return
	}
	if iNum == 0 {
		iNum = 20
	}

	keyword := ctx.Query("keyword")

	list, err := cs.CertModel.GetList(keyword, iPage, iNum)
	if err != nil {
		log.Println(err)
		common.ReturnError(ctx, common.Errors.UnexpectedError)
		return
	}
	common.ReturnData(ctx, true, list)
}

// 获取凭证列表
func (cs CertService) GetIdList(ctx *gin.Context) {
	page := ctx.Query("page")
	iPage, err1 := strconv.Atoi(page)
	num := ctx.Query("number")
	iNum, err2 := strconv.Atoi(num)
	if err1 != nil || err2 != nil {
		common.ReturnError(ctx, common.Errors.ParamIllegalError)
		return
	}
	if iNum == 0 {
		iNum = 20
	}

	keyword := ctx.Query("keyword")

	list, err := cs.CertModel.GetList(keyword, iPage, iNum)
	if err != nil {
		log.Println(err)
		common.ReturnError(ctx, common.Errors.UnexpectedError)
		return
	}
	// 清理数据
	for i := 0; i < len(list); i++ {
		list[i].PrivateKey = ""
		list[i].PublicKey = ""
		list[i].Passphrase = ""
	}
	common.ReturnData(ctx, true, list)
}

func (cs CertService) InitAdd(ctx *gin.Context) {
	level := ctx.Query("level")
	iLevel, err := strconv.Atoi(level)
	if err != nil {
		common.ReturnError(ctx, common.Errors.ParamIllegalError)
		return
	}
	hostCert := common.GenerateHostCert(iLevel)
	if hostCert == nil {
		common.ReturnMessage(ctx, false, "凭证生成失败")
	} else {
		common.ReturnData(ctx, true, hostCert)
	}
}

// 添加凭证
func (cs CertService) Add(ctx *gin.Context) {
	var form model.Cert
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(form.Name) == 0 {
		common.ReturnMessage(ctx, false, "凭证名称不能为空")
		return
	}
	if len(form.PrivateKey) == 0 {
		common.ReturnMessage(ctx, false, "私钥不能为空")
		return
	}

	state := cs.CertModel.Add(&form)
	common.ReturnMessage(ctx, state, fmt.Sprintf("%v", form.Id))
}

// 编辑凭证
func (cs CertService) Edit(ctx *gin.Context) {
	// 获取命令编号
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "凭证编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	var form model.Cert
	if err := ctx.ShouldBindJSON(&form); err != nil {
		common.ReturnMessage(ctx, false, "非法参数")
		return
	}
	form.Id = sid

	state := cs.CertModel.Edit(&form)
	common.ReturnMessage(ctx, state, id)
}

// 删除凭证
func (cs CertService) Del(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "凭证编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	state := cs.CertModel.Del(sid)
	common.ReturnMessage(ctx, state, id)
}
