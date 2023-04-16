package system

import (
	uuid "github.com/satori/go.uuid"
	"lf_web_gin/server/global"
)

type SysUser struct {
	global.PRO_MODEL           //通用属性
	UUID             uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username         string    `json:"userName" gorm:"comment:用户登陆名"`
	Password         string    `json:"password" gorm:"comment:用户密码"`
	NickName         string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	AuthorityId      uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Phone            string    `json:"phone" gorm:"comment:用户手机号"`
	Email            string    `josn:"email" gorm:"comment:用户邮箱"`
	Enable           int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

// 指定表名 gorm 有默认的表面规则
func (SysUser) TableName() string {
	return "sys_user"
}
