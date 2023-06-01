/**
 * @Author: lenovo
 * @Description:
 * @File:  movie_tx
 * @Version: 1.0.0
 * @Date: 2023/05/31 9:26
 */

package tx

import (
	"fmt"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"time"
)

type MovieTX struct{}

func NewMovie() *MovieTX {
	return &MovieTX{}
}
func (MovieTX) CreateMovieWithTX(managerID uint, param request.CreateMovieReq) (*reply.CreateMovieRly, error) {
	tx := dao.Group.DB.Begin()
	m := &automigrate.Movie{
		Name:     param.Name,
		Area:     param.Area,
		Actors:   param.Actors,
		Content:  param.Content,
		Duration: int64(param.Duration),
		ShowTime: time.Unix(int64(param.ShowTime), 0),
		Director: param.Director,
	}
	if result := tx.Model(&automigrate.Movie{}).Create(m); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	if result := tx.Model(&automigrate.Tag{}).Create(&automigrate.Tag{
		MovieID: m.ID,
	}); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	var u automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", m.ID).Find(&u); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	tx.Model(&automigrate.ManagerMovie{}).Create(&automigrate.ManagerMovie{
		MovieID: m.ID,
		UserID:  managerID,
		Content: fmt.Sprintf("用户:%s创建了电影%s", u.UserName, param.Name),
	})
	tx.Commit()
	return &reply.CreateMovieRly{MovieID: m.ID}, nil
}

func (m *MovieTX) DeleteMovieByID(mID, movieID uint) error {
	var movie automigrate.Movie
	tx := dao.Group.DB.Begin()
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Where("id =?", movieID).Find(&movie); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.Movie{}).Delete(&automigrate.Movie{}, movieID); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	var u automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", mID).Find(&u); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	tx.Model(&automigrate.ManagerMovie{}).Create(&automigrate.ManagerMovie{
		MovieID: movieID,
		UserID:  mID,
		Content: fmt.Sprintf("用户:%s创建了电影%s", u.UserName, movie.Name),
	})
	return nil
}

func (m *MovieTX) GetMovieDetails(movieID uint) (*reply.GetMovieDetails, error) {
	tx := dao.Group.DB.Begin()
	var movie automigrate.Movie
	if result := tx.Model(&automigrate.Movie{}).Where("id =?", movieID).Find(&movie); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	//从tag中获取tag字段
	t := automigrate.Tag{}
	if result := dao.Group.DB.Model(&automigrate.Tag{}).Where("movie_id =?", movieID).Find(&t); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	//todo: 查询是否是该用户关注过或者评论过
	return &reply.GetMovieDetails{
		MovieInfoRow: reply.MovieInfo{
			MovieID:    movieID,
			Name:       movie.Name,
			Area:       movie.Area,
			Actors:     movie.Actors,
			Content:    movie.Content,
			Duration:   movie.Duration,
			ShowTime:   movie.ShowTime,
			Director:   movie.Director,
			Score:      movie.Score,
			BoxOffice:  movie.BoxOffice,
			VisitCount: movie.VisitCount,
		},
		Tag: t.Tags,
	}, nil
}
