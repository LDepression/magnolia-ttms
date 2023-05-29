/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:21
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
)

type user struct{}

func (u *user) Init(r *gin.RouterGroup) {

	g := r.Group("/user")
	{
		g.POST("/register", v1.Group.User.Register)
		g.POST("/login")
	}
}
