package service

import (
	"github.com/gin-gonic/gin"
	"github.com/skye-z/colossus/backend/common"
)

type ConfigService struct {
}

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 获取配置信息
func (cs ConfigService) GetAll(ctx *gin.Context) {
	common.ReturnData(ctx, true, common.GetAll())
}

// 更新配置
func (cs ConfigService) Update(ctx *gin.Context) {
	var config Config
	if err := ctx.ShouldBindJSON(&config); err != nil {
		common.ReturnMessage(ctx, false, "非法参数")
		return
	}
	common.Set(config.Key, config.Value)
	common.ReturnMessage(ctx, true, "更新成功")
}
