package models

type Auth struct {
	Model
	// 用户名
	Username string `json:"username" grom:"column:username;not null"`
	// 昵称
	Name     string `json:"name"  grom:"column:name;not null"`
	Password string `json:"password" grom:"column:password;not null"`
}
