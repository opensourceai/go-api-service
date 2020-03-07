package models

type Model struct {
	ID         int `gorm:"primary_key" json:"id" valid:"Required;Min(1)"` // ID
	CreatedOn  int `json:"created_on"`                                    // 新建时间
	ModifiedOn int `json:"modified_on"`                                   // 修改时间
	DeletedOn  int `json:"deleted_on"`                                    // 删除时间
}
