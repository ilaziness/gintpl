package response

import (
	"gintpl/pkg/base/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

const successCode = 1

type Format struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	respData := Format{
		Code:    successCode,
		Message: "",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, respData)
}

func Error(ctx *gin.Context, err errcode.Code) {
	respData := Format{
		Code:    err.Code,
		Message: err.Message,
		Data:    err.Data,
	}
	ctx.JSON(http.StatusOK, respData)
}
