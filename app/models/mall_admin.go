package models

import "gorm.io/gorm"

// 管理员表
type Admin struct {
	Id       int        `json:"id"`       // ID
	Username string     `json:"username"` // 用户名
	Nickname string     `json:"nickname"` // 昵称
	Avatar   string     `json:"avatar"`   // 头像
	Email    string     `json:"email"`    // 电子邮箱
	AdminLog []AdminLog `json:"admin_log"`
}

func (Admin) TableName() string {
	return "mall_admin"
}

func GetAdmins() (admin []Admin) {
	db.Debug().Preload("AdminLog", func(query *gorm.DB) *gorm.DB {
		//Query the latest 2 operation logs of each user
		return query.Order("id desc").Limit(2)
	}).Order("id desc").Find(&admin)
	return admin
}
