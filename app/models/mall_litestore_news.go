package models

import (
	"gorm.io/gorm"
	"time"
)

// 首页banner
type LitestoreNews struct {
	Model
	Id         int `gorm:"primary_key" json:"id"`
	MallId     int    `json:"mall_id"` // 商城id
	Title      string `json:"title"`   // 标题
	Desc       string `json:"desc"`    // 描述
	Image      string `json:"image"`   // 图片
	Bgcolor    string `json:"-"`       // 幻灯图背景颜色
	Content    string `json:"content"` // 内容
	Type       string `json:"type"`
	Url        string `json:"url"`
	Level      int    `json:"-"`
	Starttime  int    `json:"-"` // 活动开始时间
	Endtime    int    `json:"-"` // 活动结束时间
	Status     string `json:"-"` // 状态
	Weigh      int    `json:"-"` // 权重
	Createtime int    `json:"-"` // 添加时间
	Updatetime int    `json:"-"` // 更新时间
}

func (LitestoreNews) TableName() string {
	return "litestore_news"
}

// @title 获取新闻表数据
func (newsModel *LitestoreNews) GetBanners(MallId int, Level int, limit int) (litestoreNews []*LitestoreNews) {

	db := db.Model(&litestoreNews).Where("mall_id = ? and level = ? and status = ?", MallId, Level, "normal").Order("weigh desc")
	if limit > 0 {
		db = db.Limit(limit)
	}
	err := db.Find(&litestoreNews).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return litestoreNews
}

// @title 获取新闻表单行数据
func (newsModel *LitestoreNews) GetBanner(MallId int, Level int) (litestoreNews *LitestoreNews) {
	db := db.Model(&litestoreNews).Where("mall_id = ? and level = ? and status = ?", MallId, Level, "normal").Order("weigh desc")
	if Level == 5 {
		dtime := time.Now().Unix()
		db = db.Where("type = ? and endtime > ?", "popup", dtime)
	}
	err := db.First(&litestoreNews).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return litestoreNews
}

// @title 查询用户首页是否有弹窗广告
func (newsModel *LitestoreNews) HomePopUp(MallId int, UserId int, Level int) interface{} {
	popRes := newsModel.GetBanner(MallId, Level)
	if popRes == nil {
		return nil
	}
	newsUserLog := &MallNewsUserLog{
		UserId: UserId,
		MallId: MallId,
		NewsId: popRes.Id,
	}
	rest := newsUserLog.Get(newsUserLog)
	if rest.Id == 0 {
		newsUserLog.CreatedAt = time.Now().Unix()
		newsUserLog.Add()
		return popRes
	}
	return nil
}
