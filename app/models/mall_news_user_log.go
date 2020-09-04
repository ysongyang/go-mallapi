package models

import "gorm.io/gorm"

// 用户弹窗记录表
type MallNewsUserLog struct {
	Model
	Id        int   `json:"id"`
	UserId    int   `json:"user_id"`
	MallId    int   `json:"mall_id"`
	NewsId    int   `json:"news_id"` // 活动id
	CreatedAt int64 `json:"created_at"`
}

func (MallNewsUserLog) TableName() string {
	return "news_user_log"
}

func (log *MallNewsUserLog) Add() int {
	var UserLog = MallNewsUserLog{
		UserId:    log.UserId,
		MallId:    log.MallId,
		NewsId:    log.NewsId,
		CreatedAt: log.CreatedAt,
	}
	result := db.Create(&UserLog) // 通过数据的指针来创建
	if result.RowsAffected > 0 {
		return UserLog.Id
	}
	return 0
}

func (log *MallNewsUserLog) Get(userLog *MallNewsUserLog) *MallNewsUserLog {
	var mallNewsUserLog *MallNewsUserLog
	err := db.Where("user_id = ? and mall_id = ? and news_id = ?", userLog.UserId, userLog.MallId, userLog.NewsId).First(&mallNewsUserLog).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return mallNewsUserLog
}
