package models

// 第三方登录表
type MallThird struct {
	Id           int    `json:"id"`      // ID
	UserId       int    `json:"user_id"` // 会员ID
	UnionId      string `json:"union_id"`
	ProjectType  string `json:"project_type"` // 不同项目
	Platform     string `json:"platform"`     // 第三方应用
	Openid       string `json:"openid"`       // 第三方唯一ID
	MiniOpenid   string `json:"mini_openid"`
	Openname     string `json:"openname"` // 第三方会员昵称
	AccessToken  string `json:"-"`        // AccessToken
	RefreshToken string `json:"-"`
	ExpiresIn    int    `json:"-"` // 有效期
	Avatar       string `json:"avatar"`
	Gender       int    `json:"gender"` // 性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language     string `json:"language"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	City         string `json:"city"`
	IsSubscribe  int    `json:"is_subscribe"` // 是否关注 0否 1关注
	Createtime   int    `json:"createtime"`   // 创建时间
	Updatetime   int    `json:"-"`            // 更新时间
	Logintime    int    `json:"-"`            // 登录时间
	Expiretime   int    `json:"-"`            // 过期时间
}

func (MallThird) TableName() string {
	return "third"
}
