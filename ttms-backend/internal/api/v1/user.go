/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:42
 */

package v1

import (
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct{}

func (u *user) Register(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.RegisterParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Infof("should bind failed err:%v", err)
		return
	}
	rsp, err := logic.Group.User.Register(ctx, param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rsp)
}
