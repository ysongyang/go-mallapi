package request

// 订单列表的参数定义
type GoodsListParams struct {
	MallId     int    `validate:"required"` //商城ID
	Limit      int    `validate:"required"` //分页条目
	Page       int    `validate:"required"` //页码
	Rand       int    //是否随机获取
	LabelId    int    //标签分类Id
	UserId     int    //用户uid
	GoodsName  string //商品名
	GoodsIds   string //商品id
	CategoryId int    //分类id
	Sort       int    //点击排序ID
}
