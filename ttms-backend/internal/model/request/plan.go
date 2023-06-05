/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 4:37
 */

package request

type CreatePlan struct {
	MovieID   uint    `json:"movieID" binding:"required"`
	CinemaID  uint    `json:"cinemaID" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	StartTime int64   `json:"startTime" binding:"required"`
}

type DelPlan struct {
	PlanID uint `json:"planID" binding:"required"`
}
