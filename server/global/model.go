package global

import (
	"gorm.io/gorm"
	"time"
)

type PRO_MODEL struct {
	Id           string         `gorm:"primarykey"` // 主键ID
	CreatedTime  time.Time      // 创建时间
	UpdatedTime  time.Time      // 更新时间
	DeletedTime  gorm.DeletedAt `json:"-"` // 删除时间
	ProvinceCode string         //省编码
	BureauCode   string         //局编码
}
