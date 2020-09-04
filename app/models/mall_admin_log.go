package models

// 管理员日志表
type AdminLog struct {
	Id       int    `json:"id"`       // ID
	AdminId  int    `json:"admin_id"` // 管理员ID
	Username string `json:"username"` // 管理员名字
	Url      string `json:"url"`      // 操作页面
	Title    string `json:"title"`    // 日志标题
	Ip       string `json:"ip"`       // IP
}

func (AdminLog) TableName() string {
	return "mall_admin_log"
}
