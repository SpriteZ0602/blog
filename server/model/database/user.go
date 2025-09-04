package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/gofrs/uuid"
)

type User struct {
	global.MODEL                   // gorm.Model包含ID、CreatedAt、UpdatedAt、DeletedAt字段
	UUID         uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`   // 用户UUID
	Username     string            `json:"username"`                           // 用户名
	Password     string            `json:"-"`                                  // 密码
	Email        string            `json:"email"`                              // 邮箱
	OpenId       string            `json:"openid"`                             // QQ登录的OpenID
	Avatar       string            `json:"avatar" gorm:"size:255"`             // 头像
	Address      string            `json:"address"`                            // 地址
	Signature    string            `json:"signature" gorm:"default:'签名是空白的。'"` // 个性签名
	RoleID       appTypes.RoleID   `json:"role_id"`                            // 角色ID
	Register     appTypes.Register `json:"register"`                           // 注册来源
	Freeze       bool              `json:"freeze"`                             // 是否冻结
}
