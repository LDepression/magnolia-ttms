/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:43
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type movie struct{}

func (m *movie) Init(r *gin.RouterGroup) {
	g := r.Group("/movie", middleware.Auth(), middleware.AuthManager())
	{
		g.POST("/create", v1.Group.Movie.CreateMovie)
		g.DELETE("/:id", v1.Group.Movie.DeleteMovie)
		g.GET("/:id", v1.Group.Movie.GetMovieDetails)
	}
}
