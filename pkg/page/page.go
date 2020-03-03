package page

type Page struct {
	Page    uint   `json:"page"`
	Size    uint   `json:"size"`
	OrderBy string `json:"order_by"` // 排序字段
	Sorter  string `json:"sorter"`   // 升序, 降序

}
