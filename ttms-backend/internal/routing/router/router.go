/**
 * @Author: lenovo
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:12
 */

package router

import (
	gs "github.com/swaggo/gin-swagger"
	"mognolia/internal/middleware"
	"mognolia/internal/routing"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	_ "mognolia/internal/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())
	root := r.Group("/api/v1")
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
		r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
		routing.Group.User.Init(root)
		routing.Group.Email.Init(root)
		routing.Group.Movie.Init(root)
		routing.Group.File.Init(root)
		routing.Group.Tag.Init(root)
		routing.Group.Pm.Init(root)
		routing.Group.Seat.Init(root)
		routing.Group.Cinema.Init(root)
	}
	return r
}
