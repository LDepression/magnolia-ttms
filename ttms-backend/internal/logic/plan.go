/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 4:55
 */

package logic

import (
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app/errcode"
)

type plan struct{}

func (plan) CreatePlan(req request.CreatePlan) (*reply.CreatePlanRly, errcode.Err) {
	tx := tx2.NewPlanWithTX()
	rly, err := tx.CreatePlanWithTX(req)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	return rly, nil
}
