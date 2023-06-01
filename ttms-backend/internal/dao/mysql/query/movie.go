/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:55
 */

package query

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
	"time"
)

type movie struct{}

func NewMovie() *movie {
	return &movie{}
}
func (movie) DeleteMovieByID(movieID uint) error {
	result := dao.Group.DB.Model(&automigrate.Movie{}).Delete(&automigrate.Movie{}, movieID)
	return result.Error
}
func (movie) GetMovieByID(movieID uint) (*automigrate.Movie, error) {
	var u automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Where("id = ?", movieID).Find(&u); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &u, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByPeriod(t request.GetMovieTagsAreaPeriod) ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Scopes(Paginate(t.Page, 0)).Where("show_time between ? and ?", time.Unix(t.StartTime, 0), time.Unix(t.EndTime, 0)).Order("show_time desc").Find(&movies); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return movies, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByReadCount(t request.GetMovieTagsAreaPeriod) ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Scopes(Paginate(t.Page, 0)).Order("visit_count desc").Find(&movies); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movies, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByScore(t request.GetMovieTagsAreaPeriod) ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Scopes(Paginate(t.Page, 0)).Order("score desc").Find(&movies); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movies, nil
}
