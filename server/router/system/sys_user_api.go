package system

import (
	"github.com/gin-gonic/gin"
	v1 "lf_web_gin/server/api/v1"
)

type UserApiRouter struct{}

func (u *UserApiRouter) InitUserApiRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouerApi := v1.ApiGroupApp.System.SystemUserApi

	{
		userRouter.POST("add", userRouerApi.CreateUserApi) //创建用户
	}
}