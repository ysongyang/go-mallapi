package service

import (
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/app/models/request"
	"golangPro/golang-mallapi/app/models/response"
)

// 获取首页参数
func GetIndex(params *request.IndexParams) *response.HomeIndexResult {
	var newsModel = &models.LitestoreNews{}

	var categoryModel = &models.LitestoreCategory{}

	var labelTypeModel = &models.LabelType{}

	var shopModel = models.MallShop{}
	topBanner := newsModel.GetBanners(params.MallId, 1, 0)
	adBanner := newsModel.GetBanners(params.MallId, 2, 0)
	adCategory := newsModel.GetBanners(params.MallId, 10, params.TopCategoryLimit)
	category := categoryModel.GetCategory(params.MallId, 0, 0, params.CategoryLimit)
	showStore := "off"
	isShowMall := "off"

	if params.MallId == 8 {
		showStore = "off"
		isShowMall = "on"
	}
	if params.MallId == 1 || params.MallId == 18 {
		showStore = "on"
	}

	shopInfo := shopModel.GetShopInfoById(params.MallId)

	labelList := labelTypeModel.GetAll(params)

	labels := models.GetLabels()

	result := &response.HomeIndexResult{
		TopBanner:       topBanner,
		AdBanner:        adBanner,
		AdCategory:      adCategory,
		Category:        category,
		IsNotice:        0,
		IsShowMall:      isShowMall,
		LabelList:       labelList,
		Labels:          labels,
		MallName:        shopInfo.Name,
		ShowStore:       showStore,
		TplCategoryType: shopInfo.CategoryType,
	}
	if params.UserId > 0 {
		popUp := newsModel.HomePopUp(params.MallId, params.UserId, 5)
		result.HomePopUp = popUp
	}

	return result
}
