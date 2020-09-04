package request

type IndexParams struct {
	MallId           int `validate:"required" label:"商城ID"`
	UserId           int
	CategoryId       int
	TopCategoryLimit int `validate:"required" label:"顶部分类分页条目"`
	CategoryLimit    int `validate:"required" label:"分类分页显示条目"`
	LabelGoodsPage   int //标签分类下商品列表的分页码
	LabelGoodsLimit  int //标签分类下商品列表的显示条目
}
