package model

import (
	"osstp-go-hive/system/global"
)

type SysUser struct {
	global.BaseModel
	UserId      int    `gorm:"comment:编码"  json:"user_id"`
	Username    string `json:"user_name" gorm:"size:64;comment:用户名"`
	Password    string `json:"-" gorm:"size:128;comment:密码"`
	NickName    string `json:"nick_name" gorm:"size:128;comment:昵称"`
	Phone       string `json:"phone" gorm:"size:11;comment:手机号"`
	Avatar      string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex         string `json:"sex" gorm:"size:255;comment:性别"`
	Email       string `json:"email" gorm:"size:128;comment:邮箱"`
	SideMode    string `json:"side_mode" gorm:"default:dark;comment:用户侧边主题" example:"dark"`        // 用户侧边主题
	BaseColor   string `json:"base_color" gorm:"default:#fff;comment:基础颜色" example:"#fff"`         // 基础颜色
	ActiveColor string `json:"active_color" gorm:"default:#1890ff;comment:活跃颜色" example:"#1890ff"` // 活跃颜色
}

func (s SysUser) TableName() string {
	return "sys_user"
}
