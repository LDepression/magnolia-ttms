/**
 * @Author: lenovo
 * @Description:
 * @File:  manager_movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 9:28
 */

package automigrate

import "gorm.io/gorm"

type ManagerMovie struct {
	gorm.Model
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID;references:ID"`

	UserID uint
	User   User `gorm:"foreignKey:UserID;references:ID"`

	Content string `gorm:"type:varchar(1000)"`
}
