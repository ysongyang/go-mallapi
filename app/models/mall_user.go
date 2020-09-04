package models

import (
	"errors"
	"golangPro/golang-mallapi/pkg/util"
)

// 会员表
type MallUser struct {
	Id              int         `json:"id"`                // ID
	Source          string      `json:"source"`            // 来源
	Username        string      `json:"username"`          // 用户名
	Nickname        string      `json:"nickname"`          // 昵称
	Password        string      `json:"-"`                 // 密码
	OldPassword     string      `json:"-"`                 // 老库的密码
	Salt            string      `json:"-"`                 // 密码盐
	Email           string      `json:"email"`             // 电子邮箱
	MobileFix       int         `json:"mobile_fix"`        // 手机号国家代码
	Mobile          string      `json:"mobile"`            // 手机号
	Avatar          string      `json:"avatar"`            // 头像
	Gender          int         `json:"gender"`            // 性别
	Birthday        string      `json:"birthday"`          // 生日
	Bio             string      `json:"bio"`               // 格言
	Money           float32     `json:"money"`             // 余额
	FrozenMoney     float32     `json:"frozen_money"`      // 冻结的金额
	Score           int         `json:"score"`             // 积分
	FrozenScore     int         `json:"frozen_score"`      // 冻结的积分
	LuckMoney       float32     `json:"luck_money"`        // 红包余额
	FrozenLuckMoney float32     `json:"frozen_luck_money"` // 冻结的红包余额
	Createtime      int         `json:"createtime"`        // 创建时间
	AccessToken     interface{} `json:"access_token"`      //access_token
	DbModel         BaseModel   `json:"-"`
}

func (MallUser) TableName() string {
	return "user"
}

// @title 获取用户信息
// @description 根据用户ID返回用户信息
func (user *MallUser) GetUserById(userId int) (*MallUser, error) {
	db := user.DbModel.GetDbConnect("")
	db = db.Where("id = ?", userId)
	db.First(&user)
	if user.Id == 0 {
		return nil, errors.New("不存在的用户信息")
	}
	return user, nil
}

// @title    账号登录
// @description   用户名和密码登录
func (user *MallUser) LoginByAccount(uName string, pWord string) (*MallUser, error) {

	db := user.DbModel.GetDbConnect("")

	db = db.Where("username = ? or mobile = ? ", uName, uName)
	db.First(&user)

	if user.Id == 0 {
		return nil, errors.New("未注册的用户信息")
	}

	if user.Password == util.GetPassWord(pWord, user.Salt) {
		user, err := user.GetUserById(user.Id)
		if err != nil {
			return nil, errors.New("没有找到该用户")
		}
		return user, nil
	} else {
		return nil, errors.New("用户名或密码输入有误")
	}
	return nil, errors.New("未知的用户信息")
}
