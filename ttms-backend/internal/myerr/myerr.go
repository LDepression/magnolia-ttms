/**
 * @Author: lenovo
 * @Description:
 * @File:  myerr
 * @Version: 1.0.0
 * @Date: 2023/05/29 11:02
 */

package myerr

import "mognolia/internal/pkg/app/errcode"

var (
	UserAlreadyExists = errcode.NewErr(20001, "用户已经存在了")
	CodeInvalid       = errcode.NewErr(30001, "验证码失效或者是不正确")
	SendTooMany       = errcode.NewErr(30002, "发送次数过多")
)
