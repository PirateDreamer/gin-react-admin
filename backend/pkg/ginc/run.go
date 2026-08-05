package ginc

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Run[T, R any](fn func(context.Context, *gin.Context, T) (*R, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var req T
		if err := c.Bind(&req); err != nil {
			ResFail(ctx, c, err)
			return
		}

		res, err := fn(ctx, c, req)
		if err != nil {
			ResFail(ctx, c, err)
			return
		}
		ResSuccess(ctx, c, res)
		c.Next()
	}
}
