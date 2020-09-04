package service

import "golangPro/golang-mallapi/app/models"

// @title 优惠付的订单列表
func MakeOrderList(orders []*models.LitestoreOrder) []*models.LitestoreOrder {
	if len(orders) == 0 {
		return orders
	}
	var storeModel models.MallStore
	for k, v := range orders {
		orders[k].ConsumeDed = v.CouponFee + v.PlatformDed + v.MoneyDed
		orders[k].ResultStoreInfo = storeModel.GetShopInfo(v.ShopId)
	}
	return orders
}


