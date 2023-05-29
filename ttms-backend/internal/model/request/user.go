/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:31
 */

package request

type RegisterParam struct {
	Email     string `json:"Email" binding:"required"`                     //邮箱
	Username  string `json:"Username" binding:"required,gte=3,lte=12"`     //用户名长度介于3和12位之间
	Password  string `json:"Password" binding:"required,gte=3,lte=12"`     //密码长度介于3和12之间
	EmailCode string `json:"EmailCode" binding:"required,gte=6,lte=6"`     //邮箱验证码，长度为6
	Signature string `json:"Signature" `                                   //个性签名
	Gender    string `json:"Gender" binding:"omitempty,oneof=male female"` //性别
	Avatar    string `json:"Avatar"`                                       //头像
}
