package models

type User struct {
	Model
	// 用户名
	Username string `json:"username" grom:"column:username;not null" valid:"Required"`
	// 昵称
	Name string `json:"name"  grom:"column:name;not null" valid:"Required;MaxSize(50)"`
	// 密码
	Password string `json:"password" grom:"column:password;not null" valid:"Required"`
	// 描述
	Description string `json:"description" grom:"column:description" valid:"MaxSize(200)"`
	// 性别
	Sex int `json:"sex" grom:"column:sex;not null" valid:"Min(1)"`
	// 头像地址
	AvatarSrc string `json:"avatar_src" grom:"column:avatar_src;not null"`
	// 电子邮件
	Email string `json:"email" grom:"column:email" valid:"Required;Email;MaxSize(100)"`
	// 网站
	WebSite string `json:"web_site" grom:"column:web_site" valid:"MaxSize(50)"`
	// 公司
	Company string `json:"company" grom:"column:company" valid:"MaxSize(50)"`
	// 职位
	Position string `json:"position" grom:"column:position" valid:"MaxSize(50)"`
}
