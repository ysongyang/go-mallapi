package models

import (
	"gorm.io/gorm"
)

// 商城分类表
type LitestoreCategory struct {
	Model
	Id         int `gorm:"primary_key" json:"id"`
	Pid        int    `json:"-"`     // 父id
	Name       string `json:"name"`  // 分类名称
	Image      string `json:"image"` // 图片
	Weigh      int    `json:"-"`     // 权重
	Url        string `json:"url"`
	Status     string `json:"-"`
	Createtime int    `json:"-"`
	Updatetime int    `json:"-"`
	MallId     int    `json:"mall_id"`   // 商城id
	MallName   string `json:"mall_name"` // 商城名称
	IsTop      int    `json:"-"`         // 1 :是顶部菜单 2：不是
	Type       string `json:"type"`      // category 分类页 page 单页  mall 商城  qrcode 二维码 label 标签页
	Param      string `json:"param"`     // 参数
}

func (LitestoreCategory) TableName() string {
	return "litestore_category"
}

// @title 获取分类表
func (categoryModel *LitestoreCategory) GetCategory(MallId, IsPid, IsTop, Limit int) (LitestoreCategory []*LitestoreCategory) {
	db := db.Model(&categoryModel).Where("status = ? and mall_id = ? and pid = ? and is_top = ?", "normal", MallId, IsPid, IsTop).Order("weigh desc").Limit(Limit).Find(&LitestoreCategory).Error
	if db != nil && db != gorm.ErrRecordNotFound {
		return nil
	}
	return LitestoreCategory
}

// @title 获取分类信息
func (categoryModel *LitestoreCategory) GetCategoryById(CategoryId int) *LitestoreCategory {
	err := db.Model(&categoryModel).Where("id = ? ", CategoryId).First(&categoryModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categoryModel
}

// 返回分类id切片
func (categoryModel *LitestoreCategory) GetClassIds(MallId int, CategoryId int) ([]int, error) {
	var ids []int
	if CategoryId == 0 {
		db := db.Model(&categoryModel).Where("mall_id = ?", MallId).Pluck("id", &ids).Error
		if db != nil && db != gorm.ErrRecordNotFound {
			return nil, db
		}
	} else {
		db.Model(&categoryModel).Where("mall_id = ? and pid = ?", MallId, CategoryId).Pluck("id", &ids)
		if len(ids) > 0 {
			var pids []int
			db.Model(&categoryModel).Where("pid In ?", ids).Pluck("id", &pids)
			if len(pids) > 0 {
				ids = append(ids, pids...) //切片合并
			}
		} else {
			ids = []int{CategoryId}
		}
	}
	return ids, nil
}
