package service

import (
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/pkg/util"
	"sort"
	"strconv"
)

type stores []*models.MallStore

func (s stores) Len() int      { return len(s) }
func (s stores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 嵌套结构体  将继承 MallStore 的所有属性和方法
// 所以相当于SortByName 也实现了 Len() 和 Swap() 方法
type sortByDistance struct{ stores }

// 根据元素的距离长度升序排序
func (p sortByDistance) Less(i, j int) bool {
	return p.stores[i].Distance < p.stores[j].Distance
}

// @title 计算店铺距离
// @params stores  店铺列表切片
// @params lon 经度
// @params lat 维度
func CalLonLatDistance(stores []*models.MallStore, lon float64, lat float64) []*models.MallStore {

	if len(stores) == 0 {
		return stores
	}
	for k, v := range stores {
		vLon, _ := strconv.ParseFloat(v.Lon, 64)
		vLat, _ := strconv.ParseFloat(v.Lat, 64)
		distance := util.Distance(vLon, vLat, lon, lat) //获取2个经纬度之间的距离
		stores[k].Distance = util.FloatRound(distance, 0)
	}
	sort.Sort(sortByDistance{stores})
	return stores
}

type Store struct {
	Id                    int     `json:"id"`
	UserId                int     `json:"user_id"`                // 用户id
	Name                  string  `json:"name"`                   // 店铺名称
	Contacts              string  `json:"contacts"`               // 联系人
	Phone                 string  `json:"phone"`                  // 联系电话
	Address               string  `json:"address"`                // 地址
	ShopStatus            int     `json:"shop_status"`            // 1 营业  -1 非营业
	BusinessStartHours    string  `json:"business_start_hours"`   // 营业开始时间
	BusinessEndHours      string  `json:"business_end_hours"`     // 营业结束时间
	DistributionFeeQuota  float32 `json:"distribution_fee_quota"` // 配送额度， 减免配送费
	DistributionFee       float32 `json:"distribution_fee"`       // 配送费
	DistributionRange     int     `json:"distribution_range"`     // 1 1公里  2 3公里 3 5公里  4 10公里
	Images                string  `json:"images"`
	Remark                string  `json:"remark"`                  // 店铺简介
	SendTime              int     `json:"send_time"`               // 配送时间（分钟）
	ShopLogo              string  `json:"shop_logo"`               // 店铺logo
	BasicDeliveryFee      int     `json:"basic_delivery_fee"`      // 起送费
	IsDistribution        int     `json:"is_distribution"`         // 1 支持配送 0 不支持配送
	IsDada                int     `json:"is_dada"`                 //支持达达配送 1支持 0不支持
	ShareRatio            float32 `json:"share_ratio"`             // 商品分享相关的分佣比例,0则不开启[收费店铺有效]
	MerchantDedProportion float32 `json:"merchant_ded_proportion"` // 商家支付配送费比例
	UserDedProportion     float32 `json:"user_ded_proportion"`     // 用户支付配送费比例
	LuckQuota             float32 `json:"luck_quota"`              // 红包补贴额度（红包抵扣时用）
	SmallPacket           int     `json:"small_packet"`            // 最小红包使用系数（%）
	MaxPacket             int     `json:"max_packet"`              // 最大红包使用系数（%）
	IsVip                 int     `json:"is_vip"`                  // 0 普通商户 1 收费商户
	IsOpenDadaShop        int     `json:"is_open_dada_shop"`       // 是否开启达达门店 0 : 未开启 1：开启
}
