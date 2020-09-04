package response

import "golangPro/golang-mallapi/app/models"

// 首页homeIndex接口返回的结构体
type HomeIndexResult struct {
	TopBanner       []*models.LitestoreNews     `json:"topBanner"`    //首页顶部Banner
	AdBanner        []*models.LitestoreNews     `json:"adBanner"`     //首页顶部广告Banner
	AdCategory      []*models.LitestoreNews     `json:"adCategory"`   //首页广告分类
	Category        []*models.LitestoreCategory `json:"category"`     //首页分类
	HomePopUp       interface{}                 `json:"home_pop_up"`  //是否请求弹窗API
	IsNotice        int                         `json:"is_notice"`    //首页资讯未读数量
	IsShowMall      string                      `json:"is_show_mall"` //首页是否显示自营商城
	LabelList       []*models.LabelType         `json:"label_list"`
	Labels          []models.LabelType          `json:"labels"`
	MallName        string                      `json:"mall_name"`
	ShowStore       string                      `json:"show_store"`        //是否显示附近的店铺
	TplCategoryType int                         `json:"tpl_category_type"` //分类页模板
}
