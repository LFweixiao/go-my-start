package system

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model/system"
	"lf_web_gin/server/utils"
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
	u.UUID = uuid.NewV4()
	u.Password = utils.BcryptHash(u.Password)
	u.CreatedTime = time.Now()
	u.UpdatedTime = time.Now()
	err = global.PRO_DB.Create(&u).Error
	return u, err
}

// Login 用户登陆
func (UserServer *UserServer) Login(vo *system.SysUser) (userRep *system.SysUser, err error) {

	if nil == global.PRO_DB {
		return nil, fmt.Errorf("do not init")
	}

	var user system.SysUser
	err = global.PRO_DB.Where("UUID = ?", vo.UUID).First(&user).Error
	if nil == err {
		if ok := utils.BcryptCheck(vo.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		//其他操作
	}
	return &user, err
}
