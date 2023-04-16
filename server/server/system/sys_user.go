package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model/system"
	"time"
)

type UserServer struct{}

// Register 注册新用户
func (userServer *UserServer) Register(u system.SysUser) (userVO system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.PRO_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userVO, errors.New("用户名以注册")
	}
	//添加uuid  密码加密
	u.CreatedTime = time.Now()
	u.UpdatedTime = time.Now()
	err = global.PRO_DB.Create(&u).Error
	return u, err
}
