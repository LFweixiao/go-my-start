package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model/common/response"
	"lf_web_gin/server/model/system"

	systemRes "lf_web_gin/server/model/system/response"
)

type SystemUserApi struct{}

func (s *SystemUserApi) CreateUserApi(c *gin.Context) {

	var use system.SysUser
	err := c.ShouldBindJSON(&use)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	proModel := global.PRO_MODEL{
		ProvinceCode: use.ProvinceCode,
		BureauCode:   use.BureauCode,
	}

	user := &system.SysUser{
		PRO_MODEL: proModel,
		Username:  use.Username,
		Password:  use.Password,
		NickName:  use.NickName,
		Phone:     use.Phone,
		Email:     use.Email,
	}

	u, err := userServer.Register(*user)
	if err != nil {
		global.PRO_LOG.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: u}, "注册失败", c)
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: u}, "请求成功", c)
}
