package system

import (
	"github.com/gin-gonic/gin"
	"lf_web_gin/server/model/common/response"
)

type SystemUserApi struct{}

func (s *SystemUserApi) CreateUserApi(c *gin.Context) {
	response.OkWithMessage("请求成功", c)
}
