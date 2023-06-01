/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:12
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
)

type movie struct{}

// CreateMovie
// @Tags      movie
// @Summary   创建电影
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query     request.CreateMovieReq  true  "创建电影相关参数"
// @Success   200            {object}  common.State{reply.CreateMovieRly}  "1001:参数有误 1003:系统错误 2001:鉴权失败 20001:用户已存在 30002:发送次数过多"
// @Router    /api/v1/email/send [post]
func (e *movie) CreateMovie(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	param := request.CreateMovieReq{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rep, err := logic.Group.Movie.CreateMovie(ctx, param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

// DeleteMovie
// @Tags      movie
// @Summary   删除电影(不能删除在演出计划里面的)
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query     request.DeleteMovieReq  true  "根据movieID删除电影"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误 2001:鉴权失败"
// @Router    /api/v1/movie/delete [delete]
func (e *movie) DeleteMovie(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	param := request.DeleteMovieReq{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Movie.DeleteMovie(ctx, param.MovieID); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// GetMovieDetails
// @Tags      movie
// @Summary   获取电影详细信息
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query    request.GetMovieByID  true  "创建电影相关参数"
// @Success   200            {object}  common.State{reply.GetMovieDetails}  "1001:参数有误 1003:系统错误 2001:鉴权失败 "
// @Router    /api/v1/movie/detail [post]
func (e *movie) GetMovieDetails(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.GetMovieByID
	if err := ctx.ShouldBindQuery(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	movieDetail, err1 := logic.Group.Movie.GetMovieByID(param.MovieID)
	if err1 != nil {
		rly.Reply(err1)
		return
	}
	rly.Reply(nil, movieDetail)
}

func (e *movie) GetMovieByTagAreaPeriod(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.GetMovieTagsAreaPeriod
	if err := ctx.ShouldBindQuery(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := param.Check(); err != nil {
		rly.Reply(err)
		return
	}
	result, err := logic.Group.Movie.GetMovieByTagAreaPeriod(param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}
