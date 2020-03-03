package models

type Model struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedOn  int  `json:"created_on"`
	ModifiedOn int  `json:"modified_on"`
	DeletedOn  int  `json:"deleted_on"`
}
