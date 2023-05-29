/**
 * @Author: lenovo
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:12
 */

package router

import (
	"mognolia/internal/routing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	root := r.Group("/api/v1")
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
		routing.Group.User.Init(root)
		routing.Group.Email.Init(root)
	}
	return r
}
