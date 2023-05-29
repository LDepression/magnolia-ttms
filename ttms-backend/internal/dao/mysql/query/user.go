/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 11:04
 */

package query

import (
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"

	"gorm.io/gorm"
)

type user struct{}

func NewUser() *user {
	return &user{}
}
func (u *user) FindUserByEmail(email string) (*automigrate.User, error) {
	var user automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("email =?", email).Find(&user); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (u *user) Register(us request.RegisterParam, hashPassword string, role automigrate.Roler) (*automigrate.User, error) {
	user := &automigrate.User{
		Email:     us.Email,
		Password:  hashPassword,
		Signature: us.Signature,
		Avatar:    us.Avatar,
		Gender:    automigrate.Gend(us.Gender),
		Role:      role,
	}
	if result := dao.Group.DB.Model(&automigrate.User{}).Create(user); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, nil

}
