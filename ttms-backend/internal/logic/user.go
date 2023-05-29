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
	"errors"
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
	"mognolia/internal/pkg/password"
	workemail "mognolia/internal/work/email"

	"gorm.io/gorm"
)

type user struct{}

func (u *user) Register(ctx context.Context, params request.RegisterParam) (*reply.RegisterRly, errcode.Err) {
	//先判断一下是否注册过了
	q := query.NewUser()
	_, err := q.FindUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//执行注册操作
			if ok := workemail.CheckEmailAndCodeValid(params.Email, params.EmailCode); !ok {
				return nil, myerr.CodeInvalid
			}
			hashPassword, err := password.HashPassword(params.Password)
			if err != nil {
				return nil, errcode.ErrServer
			}
			role := automigrate.Vistor
			if params.Avatar == "" {
				params.Avatar = global.Settings.Rule.DefaultAccountAvatar
			}
			if params.Gender == "" {
				params.Gender = string(automigrate.UnKnown)
			}
			userInfo, err := q.Register(params, hashPassword, role)
			if err != nil {
				return nil, errcode.ErrServer
			}
			return &reply.RegisterRly{UserID: userInfo.ID}, nil
		}
		return nil, errcode.ErrServer
	}
	return nil, myerr.UserAlreadyExists
}
