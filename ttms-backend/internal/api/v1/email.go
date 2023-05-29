/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/05/29 12:11
 */

package v1

import (
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type email struct{}

func (e *email) SendEmail(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.SendEmailParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Info("ctx.ShouldBindJSON failed", zap.Any("err", err))
		return
	}
	if err := logic.Group.Email.SendEmail(param.Email); err != nil {
		zap.S().Info("logic.Group.Email.SendEmail ")
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}
