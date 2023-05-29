/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:46
 */

package logic

import (
	"context"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app/errcode"
)

type user struct{}

func (u *user) Register(ctx context.Context, params request.RegisterParam) (err errcode.Err) {

}
