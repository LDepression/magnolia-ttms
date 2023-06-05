/**
 * @Author: lenovo
 * @Description:
 * @File:  auto
 * @Version: 1.0.0
 * @Date: 2023/06/03 17:41
 */

package logic

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	t "mognolia/internal/pkg/task"
	"mognolia/internal/pkg/utils"
)

type auto struct{}

var tasks []t.Task

func (a *auto) Work() {
	ctx := context.Background()
	movieReadCountFlush2DBTask := t.Task{
		Name:            "movieReadCountFlush2DBTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.AutoFlushReadCount2DBTime,
		TimeoutDuration: global.Settings.Serve.DefaultContextTimeout,
		F:               movieVisitCountFlush(),
	}
	peopleFavorToCache := t.Task{
		Name:            "peopleFavorToCache",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.AutoFlushReadCount2DBTime,
		TimeoutDuration: global.Settings.Serve.DefaultContextTimeout,
		F:               peopleFavorToCache(),
	}

	tasks = append(tasks, movieReadCountFlush2DBTask, peopleFavorToCache)
	a.Init()
}

func (a *auto) Init() {
	for i := range tasks {
		t.NewTickerTask(tasks[i])
	}
}

func movieVisitCountFlush() t.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("start movieVisitCountFlush task....")
		ctx, cancle := context.WithTimeout(parentCtx, global.Settings.Serve.DefaultContextTimeout)
		defer cancle()
		readMap, err := dao.Group.Redis.GetAllDataAndFlushThem(ctx)
		if err != nil {
			global.Logger.Info("GetAllDataAndFlushThem failed...")
			return
		}
		tx := dao.Group.DB.Begin()
		for movieID, count := range readMap {
			i := utils.StringToIDMust(movieID)
			result := tx.Model(&automigrate.Movie{}).Where("id = ?", i).UpdateColumn("visit_count", gorm.Expr("visit_count + ?", count))
			if result.Error != nil {
				tx.Rollback()
				global.Logger.Info("GetMovie failed")
				return
			}
		}
		tx.Commit()
	}
}

func peopleFavorToCache() t.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("start peopleFavorToCache task ...")
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Serve.DefaultContextTimeout)
		defer cancel()
		page, size := global.Settings.Rule.DefaultUserFavorPage, global.Settings.Rule.DefaultUserFavorSize
		q := query.NewMovie()
		movieInfos, err := q.GetMovieOrderByFavorNum(page * size)
		if err != nil {
			zap.S().Info("get nothing ...")
			return
		}

		for i := 1; i <= page; i++ {
			start := (i - 1) * size
			end := start + size
			ok := false
			if end >= len(movieInfos) {
				end = len(movieInfos)
				ok = true
			}
			if err := dao.Group.Redis.Set(ctx, utils.IDToSting(uint(i)), movieInfos[start:end]); err != nil {
				zap.S().Infof("set to cache failed,err:%v", err)
				return
			}
			if ok {
				break
			}
		}
	}
}
