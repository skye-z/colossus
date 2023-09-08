package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"github.com/skye-z/colossus/backend/common"
	"github.com/skye-z/colossus/backend/model"
)

type CodeService struct {
	CodeModel model.CodeModel
}

// 获取命令列表
func (cs CodeService) GetList(ctx *gin.Context) {
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

	list, err1 := cs.CodeModel.GetList(keyword, platform, system, iPage, iNum)
	if err1 != nil {
		common.ReturnError(ctx, common.Errors.UnexpectedError)
		return
	}
	common.ReturnData(ctx, true, list)
}

// 添加命令
func (cs CodeService) Add(ctx *gin.Context) {
	var form model.Code
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(form.Name) == 0 {
		common.ReturnMessage(ctx, false, "快捷命令名称不能为空")
		return
	}
	if len(form.Content) == 0 {
		common.ReturnMessage(ctx, false, "快捷命令内容不能为空")
		return
	}
	form.Keyword = cs.getPinyin(form.Name)

	state := cs.CodeModel.Add(&form)
	common.ReturnMessage(ctx, state, fmt.Sprintf("%v", form.Id))
}

// 编辑命令
func (cs CodeService) Edit(ctx *gin.Context) {
	// 获取命令编号
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "命令编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	var form model.Code
	if err := ctx.ShouldBindJSON(&form); err != nil {
		common.ReturnMessage(ctx, false, "非法参数")
		return
	}
	form.Keyword = cs.getPinyin(form.Name)
	form.Id = sid

	state := cs.CodeModel.Edit(&form)
	common.ReturnMessage(ctx, state, id)
}

// 删除命令
func (cs CodeService) Del(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		common.ReturnMessage(ctx, false, "命令编号不能为空")
		return
	}
	sid, _ := strconv.ParseInt(id, 10, 64)

	state := cs.CodeModel.Del(sid)
	common.ReturnMessage(ctx, state, id)
}

func (cs CodeService) getPinyin(txt string) string {
	py := pinyin.NewArgs()
	pyList := pinyin.Pinyin(txt, py)
	var pyText string
	for x := 0; x < len(pyList); x++ {
		for y := 0; y < len(pyList[x]); y++ {
			pyText += pyList[x][y]
		}
	}
	return pyText
}
