package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/common"
	"github.com/skye-z/colossus/backend/model"
)

type HostService struct {
	HostModel model.HostModel
}

// 添加主机
func (hs HostService) Add(ctx *gin.Context) {
	var form model.Host
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !check(ctx, form) {
		return
	}
	if form.AuthType == "cert" && len(form.Cert) == 0 {
		common.ReturnMessage(ctx, false, "证书私钥不能为空")
		return
	}
	if form.AuthType == "pwd" && len(form.Secret) == 0 {
		common.ReturnMessage(ctx, false, "登录密码不能为空")
		return
	}
	if len(form.AuthType) == 0 {
		common.ReturnMessage(ctx, false, "授权类型不能为空")
		return
	}

	state := hs.HostModel.Add(&form)
	common.ReturnMessage(ctx, state, fmt.Sprintf("%v", form.Id))
}

// 编辑主机
func (hs HostService) Edit(ctx *gin.Context) {
	// 获取主机编号
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "主机编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	var form model.Host
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !check(ctx, form) {
		return
	}

	form.Id = sid
	state := hs.HostModel.Edit(&form)
	common.ReturnMessage(ctx, state, id)
}

// 删除主机
func (hs HostService) Del(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "主机编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	state := hs.HostModel.Del(sid)
	common.ReturnMessage(ctx, state, id)
}

// 校验输入参数
func check(ctx *gin.Context, form model.Host) bool {
	if len(form.Name) == 0 {
		common.ReturnMessage(ctx, false, "主机名称不能为空")
		return false
	}
	if len(form.Address) == 0 {
		common.ReturnMessage(ctx, false, "访问地址不能为空")
		return false
	}
	if len(form.User) == 0 {
		common.ReturnMessage(ctx, false, "登录用户不能为空")
		return false
	}
	if len(form.AuthType) == 0 {
		common.ReturnMessage(ctx, false, "授权类型不能为空")
		return false
	}
	return true
}
