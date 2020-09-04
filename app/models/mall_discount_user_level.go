package models

// 推广等级表
type MallDiscountUserLevel struct {
	Id             int       `json:"id" gorm:"primary_key;auto_increment;`
	Name           string    `json:"name"`            // 级别名称
	CommissionRate float32   `json:"commission_rate"` // 佣金比例（利润百分比）
	Amount         float32   `json:"amount"`          // 费用
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	CreateTime     int       `json:"create_time"`
	BaseModel      BaseModel `json:"-"`
	IsChoose       int       `json:"is_choose"`
}

func (MallDiscountUserLevel) TableName() string {
	return "discount_user_level"
}

// @title 获取等级列表
func (model *MallDiscountUserLevel) GetLevels(UserId int) []*MallDiscountUserLevel {
	db := model.BaseModel.GetDbConnect("")

	db = db.Where("type = ?", "default")
	db = db.Where("status = ?", "normal")

	var models []*MallDiscountUserLevel

	db.Find(&models)
	var mallUpgradeOrder MallUpgradeOrder
	if len(models) > 0 && UserId > 0 {
		levelId := 0
		res := mallUpgradeOrder.GetUpgradeOrder(UserId)
		if res.Id > 0 {
			levelId = res.LevelId
			for k, v := range models {
				models[k].IsChoose = 0
				if levelId > v.Id {
					models[k].IsChoose = 1
				}
			}
		}

	}

	return models
}
