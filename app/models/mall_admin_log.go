package models

// 管理员日志表
type AdminLog struct {
	Id         int    `json:"id"`         // ID
	AdminId    int    `json:"admin_id"`   // 管理员ID
	Username   string `json:"username"`   // 管理员名字
	Url        string `json:"url"`        // 操作页面
	Title      string `json:"title"`      // 日志标题
	Content    string `json:"content"`    // 内容
	Ip         string `json:"ip"`         // IP
	Useragent  string `json:"useragent"`  // User-Agent
	Createtime int    `json:"createtime"` // 操作时间
}
