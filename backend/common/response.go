package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ReturnError(ctx *gin.Context, err CustomError) {
	ctx.JSON(200, err)
	ctx.Abort()
}

type commonResponse struct {
	State   bool   `json:"state"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

func ReturnMessage(ctx *gin.Context, state bool, message string) {
	ctx.JSON(200, commonResponse{
		State:   state,
		Message: message,
		Time:    time.Now().Unix(),
	})
	ctx.Abort()
}

type dataResponse struct {
	State bool  `json:"state"`
	Data  any   `json:"data"`
	Time  int64 `json:"time"`
}

func ReturnData(ctx *gin.Context, state bool, data any) {
	ctx.JSON(200, dataResponse{
		State: state,
		Data:  data,
		Time:  time.Now().Unix(),
	})
	ctx.Abort()
}

func ReturnSuccess(ctx *gin.Context, obj any) {
	ctx.JSON(200, obj)
	ctx.Abort()
}
