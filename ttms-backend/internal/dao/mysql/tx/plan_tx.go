/**
 * @Author: lenovo
 * @Description:
 * @File:  plan_tx
 * @Version: 1.0.0
 * @Date: 2023/06/04 5:09
 */

package tx

import (
	"errors"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"time"
)

var ErrPlanConflict = errors.New("演出计划出现冲突")

type planTX struct{}

func NewPlanWithTX() *planTX {
	return &planTX{}
}

func (planTX) CreatePlanWithTX(req request.CreatePlan) (*reply.CreatePlanRly, error) {

	res := &reply.CreatePlanRly{}
	tx := dao.Group.DB.Begin()
	//先判断该电影是否存在
	var movieInfo automigrate.Movie
	if result := tx.Model(&automigrate.Movie{}).Where("id =?", req.MovieID).Find(&movieInfo); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.Cinema{}).Where("id=?", req.CinemaID).Find(&automigrate.Cinema{}); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	endAt := time.Unix(req.StartTime, 0).Add(time.Duration(movieInfo.Duration) * time.Minute)
	var ids []uint
	if result := tx.Raw(`
	SELECT
	p.id 
FROM
	plans p 
WHERE
	p.cinema_id =? 
	AND (
		p.start_at <= ? AND P.end_at >=? 
		OR p.start_at <=? 
		AND p.end_at <=? 
		OR p.start_at BETWEEN ? 
		AND ? 
		OR P.end_at BETWEEN ? 
	AND ? 
	)
	for update
`, req.MovieID, req.StartTime, endAt, req.StartTime, endAt, req.StartTime, endAt).Scan(&ids); result.RowsAffected != 0 {
		tx.Rollback()
		return nil, ErrPlanConflict
	}

	//创建plan表
	plan := &automigrate.Plan{
		MovieID:  req.MovieID,
		CinemaID: req.CinemaID,
		StartAt:  time.Unix(req.StartTime, 0),
		EndAt:    endAt,
		Price:    req.Price,
	}
	if result := tx.Model(&automigrate.Plan{}).Create(&plan); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}

	//TODO:获取seat表,以及ticket
	var seats []automigrate.Seat
	if result := tx.Model(&automigrate.Seat{}).Where("cinema_id = ?", req.CinemaID).Find(&seats); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	//创建票

	type CreateTicketParam struct {
		Price  float64 `json:"price"`
		PlanID uint    `json:"plan_id"`
		SeatID uint    `json:"seat_id"`
	}
	CreateTicketArg := make([]CreateTicketParam, 0, len(seats))
	for i := range seats {
		CreateTicketArg = append(CreateTicketArg, CreateTicketParam{
			Price:  req.Price,
			PlanID: plan.ID,
			SeatID: seats[i].ID,
		})
	}
	if result := tx.Model(&automigrate.Ticket{}).Create(&CreateTicketArg); result.RowsAffected == 0 {
		return nil, result.Error
	}
	res.EndAt = endAt
	res.StartAt = plan.StartAt

	tx.Commit()
	return res, nil
}
