package ginc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

func ResFailWithData(ctx context.Context, c *gin.Context, err error, data any) {
	var resErr *BizError
	switch e := err.(type) {
	case *BizError:
		resErr = e
	default:
		resErr = NewCommBizErr("服务不见了~")
	}

	c.JSON(http.StatusOK, Response{
		Code:    resErr.Code,
		Message: resErr.Message,
		Error:   resErr.Error(),
		Data:    data,
	})
}

func ResFail(ctx context.Context, c *gin.Context, err error) {
	ResFailWithData(ctx, c, err, nil)
}

func ResSuccess(ctx context.Context, c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    "0",
		Message: "成功",
		Data:    data,
	})
}
