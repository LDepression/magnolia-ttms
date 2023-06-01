/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:16
 */

package request

import "mognolia/internal/pkg/app/errcode"

type CreateMovieReq struct {
	Name     string   `json:"name" binding:"required,gte=1,lte=20"`       //电影名称
	Duration int16    `json:"duration" binding:"required,gte=1"`          //时长(分钟)
	ShowTime int64    `json:"show_time" binding:"required"`               //上映时间戳
	Tags     []string `json:"tags" binding:"required"`                    //电影标签
	Actors   []string `json:"actors" binding:"required,gte=1,dive,max=5"` //演员
	Director string   `json:"director" binding:"required,get=1,lte=5"`    //导演
	Area     string   `json:"area" binding:"required,gte=1,lte=20"`       //地区
	Avatar   string   `json:"avatar"`
	Content  string   `json:"content" binding:"required,gte=1,maxLength=1000"`
}

type DeleteMovieReq struct {
	MovieID uint `json:"movie_id" binding:"required"`
}

type GetMovieByID struct {
	MovieID uint `json:"movieID" binding:"required"`
}

type GetMovieTagsAreaPeriod struct {
	Tag       string `json:"tag" `
	Area      string `json:"area" `
	Period    string `json:"period" `
	OrderType string `json:"order_type" binding:"required,oneof = period,score,readCount"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Page      int    `json:"page"`
}

func (g *GetMovieTagsAreaPeriod) Check() errcode.Err {
	var msg string
	switch {
	case g.StartTime > g.EndTime:
		msg = "起始时间不能大于结束时间"
	default:
		if g.Area == "" {
			g.Area = "%"
		}
		if g.Tag == "" {
			g.Tag = "%"
		}
		return nil
	}
	return errcode.ErrServer.WithDetails(msg)
}
