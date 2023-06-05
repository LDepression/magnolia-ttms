/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 0:40
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/global"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
)

type plan struct{}

func (plan) CreatePlan(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.CreatePlan
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		global.Logger.Info("Error creating plan")
		return
	}
	result, err := logic.Group.Plan.CreatePlan(param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

//func (plan) DelPlan(ctx *gin.Context) {
//	rly := app.NewResponse(ctx)
//	var req request.DelPlan
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		base.HandleValidatorError(ctx, err)
//		return
//	}
//	if err := logic.Group.Plan.DelPlan(req.PlanID); err != nil {
//		zap.S().Infof("logic.Group.Plan.DelPlan failed: %v", err)
//		return
//	}
//	rly.Reply(nil)
//}
