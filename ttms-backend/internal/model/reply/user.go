/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:54
 */

package reply

import "mognolia/internal/pkg/token"

type RegisterRly struct {
	UserID uint `json:"user_id"`
}

type UserInfoReply struct {
	AvatarURL string `json:"AvatarURL"`
	Role      string `json:"Role"`
	UserName  string `json:"user_name"`
}
type LoginRly struct {
	UserID       uint           `json:"user_id"`
	AccessToken  string         `json:"AccessToken"`
	RefreshToken string         `json:"RefreshToken"`
	PayLoad      *token.Payload `json:"Payload"`
	UserInfo     UserInfoReply  `json:"user_info"`
}

type RefreshRly struct {
	AccessToken string `json:"AccessToken"`
}

type UserInfo struct {
	Username  string `json:"username"`
	Signature string `json:"signature"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
}

type UserList struct {
	UserInfos []*UserInfo `json:"user_infos"`
	Total     int         `json:"total"`
}
