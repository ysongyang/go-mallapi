package models

import "gorm.io/gorm"

// 管理员表
type Admin struct {
	Id           int        `json:"id"`       // ID
	Username     string     `json:"username"` // 用户名
	Nickname     string     `json:"nickname"` // 昵称
	Password     string     `json:"-"`        // 密码
	Salt         string     `json:"-"`        // 密码盐
	Avatar       string     `json:"avatar"`   // 头像
	Email        string     `json:"email"`    // 电子邮箱
	Loginfailure int        `json:"-"`        // 失败次数
	Logintime    int        `json:"-"`        // 登录时间
	Loginip      string     `json:"-"`        // 登录IP
	Createtime   int        `json:"-"`        // 创建时间
	Updatetime   int        `json:"-"`        // 更新时间
	Token        string     `json:"-"`        // Session标识
	Status       string     `json:"-"`        // 状态
	UserId       int        `json:"user_id"`  // 用户id
	CompanyId    int        `json:"-"`
	MallId       int        `json:"mall_id"`
	ShopId       int        `json:"shop_id"`
	Phone        string     `json:"phone"`
	AdminLog     []AdminLog `json:"admin_log"`
}

func GetAdmin(adminId int) Admin {
	var admin Admin
	db.Debug().Model(&Admin{}).
		Preload("AdminLog", func(query *gorm.DB) *gorm.DB {
			return query.Order("id desc")
		}).
		Where("id = ?", adminId).
		First(&admin)
	return admin
}
