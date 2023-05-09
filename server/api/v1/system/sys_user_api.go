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

// CreateUserApi 注册新用户
func (s *SystemUserApi) CreateUserApi(c *gin.Context) {

	var use system.SysUser
	err := c.ShouldBindJSON(&use)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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

// LoginUserApi 用户登陆
func (s *SystemUserApi) LoginUserApi(c *gin.Context) {
	var use system.SysUser
	err := c.ShouldBindJSON(&use)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	queryVO := &system.SysUser{
		UUID:     use.UUID,
		Password: use.Password,
	}
	//这里可以加验证码之类的

	userRep, err := userServer.Login(queryVO)

	if err != nil {
		global.PRO_LOG.Error("登陆失败！用户名不存在或密码错误", zap.Error(err))
		response.FailWithMessage("用户不存在或密码错误！", c)
		return
	}
	if userRep.Enable != 1 {
		global.PRO_LOG.Error("登陆失败！用户被冻结", zap.Error(err))
		response.FailWithMessage("用户禁止登陆！", c)
	}
	// TODO Token

	return

}

// TokenNext 登陆后添加jwt
func (s *SystemUserApi) TokenNext(c *gin.Context, user system.SysUser) {

}
