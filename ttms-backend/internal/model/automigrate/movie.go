/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/30 22:29
 */

package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Name       string   `gorm:"type:varchar(255);not null"`
	Area       string   `gorm:"type:varchar(255);not null"`
	Actors     []string `gorm:"type:varchar(255);not null"`
	Content    string   `gorm:"type:varchar(255);not null"`
	Duration   int64    `gorm:"type:int;not null"`
	ShowTime   time.Time
	Director   string  `gorm:"type:varchar(255);not null"`
	Score      float32 `gorm:"type:float;not null"`
	BoxOffice  float32 `gorm:"type:float;not null"`
	VisitCount int     `gorm:"type:int;not null"`
}
