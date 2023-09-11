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
	var cache model.AddHost
	if err := ctx.ShouldBindJSON(&cache); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	form := addHost2Host(cache)

	if !check(ctx, form) {
		return
	}
	if form.AuthType == "cert" && form.Cert == 0 {
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

	var cache model.AddHost
	if err := ctx.ShouldBindJSON(&cache); err != nil {
		common.ReturnMessage(ctx, false, "非法参数")
		return
	}

	// 查询已有数据
	form := &model.Host{
		Id: sid,
	}
	hs.HostModel.GetItem(form)
	// 更新已有数据
	form.Name = cache.Name
	form.Platform = cache.Platform
	form.System = cache.System
	form.Region = cache.Region
	form.Usage = cache.Usage
	form.Period = cache.Period
	form.Address = cache.Address
	form.Port = cache.Port
	form.AuthType = cache.AuthType
	form.User = cache.User
	// 判断证书是否需要更新
	if cache.Cert == 0 && cache.Cert != form.Cert {
		form.Cert = cache.Cert
	}
	// 判断密码是否需要更新
	if len(cache.Secret) > 0 && cache.Secret != form.Secret {
		form.Secret = cache.Secret
	}
	if !check(ctx, *form) {
		return
	}

	state := hs.HostModel.Edit(form)
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

// 获取主机列表
func (hs HostService) GetList(ctx *gin.Context) {
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
	platform := ctx.Query("platform")
	system := ctx.Query("system")
	region := ctx.Query("region")
	usage := ctx.Query("usage")
	period := ctx.Query("period")

	list, err1 := hs.HostModel.GetList(keyword, platform, system, region, usage, period, iPage, iNum)
	if err1 != nil {
		common.ReturnError(ctx, common.Errors.UnexpectedError)
		return
	}
	common.ReturnData(ctx, true, list)
}

// 获取主机
func (hs HostService) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "主机编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	form := &model.Host{
		Id: sid,
	}

	hs.HostModel.GetItem(form)
	if len(form.Name) == 0 {
		common.ReturnMessage(ctx, false, "主机不存在")
	}
	form.Secret = ""
	common.ReturnData(ctx, true, form)
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

func addHost2Host(cache model.AddHost) model.Host {
	form := model.Host{
		Id:       cache.Id,
		Name:     cache.Name,
		Platform: cache.Platform,
		System:   cache.System,
		Region:   cache.Region,
		Usage:    cache.Usage,
		Period:   cache.Period,
		Address:  cache.Address,
		Port:     cache.Port,
		AuthType: cache.AuthType,
		User:     cache.User,
		Cert:     cache.Cert,
		Secret:   cache.Secret,
	}
	return form
}
