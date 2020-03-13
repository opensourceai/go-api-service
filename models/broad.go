package models

// 版块
type Board struct {
	Model
	Name string `json:"name" gorm:"column:name" valid:"Required"` // 名称
}
