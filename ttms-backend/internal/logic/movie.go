/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:51
 */

package logic

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/middleware"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
)

type movie struct{}

const (
	OrderByPeriod    = "period"
	OrderByReadCount = "read_count"
	OrderByScore     = "score"
)

func (m *movie) CreateMovie(ctx *gin.Context, param request.CreateMovieReq) (*reply.CreateMovieRly, errcode.Err) {
	tx := tx2.NewMovie()
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	rly, err1 := tx.CreateMovieWithTX(content.ID, param)
	if err1 != nil {
		return nil, errcode.ErrServer.WithDetails(err1.Error())
	}
	return &reply.CreateMovieRly{MovieID: rly.MovieID}, nil
}

func (m *movie) DeleteMovie(ctx *gin.Context, movieID uint) errcode.Err {
	//TODO:先去判断一下该电影时候在演出计划中
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	tx := tx2.NewMovie()
	if err := tx.DeleteMovieByID(uid, movieID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (m *movie) GetMovieByID(movieID uint) (*reply.GetMovieDetails, errcode.Err) {
	tx := tx2.NewMovie()
	movieDetail, err := tx.GetMovieDetails(movieID)
	if err != nil {
		return nil, myerr.NoRecords
	}
	if err := dao.Group.Redis.AddReadCountToMovie(context.Background(), movieID); err != nil {
		return nil, errcode.ErrServer
	}
	return movieDetail, nil
}

// 我对不起这个接口 呜呜呜~~~~~~~~~~
func (m *movie) GetMovieByTagAreaPeriod(param request.GetMovieTagsAreaPeriod) (*reply.GetMovieRly, errcode.Err) {
	var result reply.GetMovieRly
	switch param.OrderType {
	case OrderByPeriod:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByPeriod(param)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:    movie.ID,
				Name:       movie.Name,
				Area:       movie.Area,
				Actors:     movie.Actors,
				Content:    movie.Content,
				ShowTime:   movie.ShowTime,
				Director:   movie.Director,
				Score:      movie.Score,
				BoxOffice:  movie.BoxOffice,
				VisitCount: movie.VisitCount,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	case OrderByReadCount:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByReadCount(param)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:    movie.ID,
				Name:       movie.Name,
				Area:       movie.Area,
				Actors:     movie.Actors,
				Content:    movie.Content,
				ShowTime:   movie.ShowTime,
				Director:   movie.Director,
				Score:      movie.Score,
				BoxOffice:  movie.BoxOffice,
				VisitCount: movie.VisitCount,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	case OrderByScore:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByScore(param)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:    movie.ID,
				Name:       movie.Name,
				Area:       movie.Area,
				Actors:     movie.Actors,
				Content:    movie.Content,
				ShowTime:   movie.ShowTime,
				Director:   movie.Director,
				Score:      movie.Score,
				BoxOffice:  movie.BoxOffice,
				VisitCount: movie.VisitCount,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	}
	return &result, nil
}
