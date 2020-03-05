package page

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
)

type Page struct {
	PageNum int    `json:"page" valid:"Min(0)"`
	Size    int    `json:"size" valid:"Min(1)"`
	OrderBy string `json:"order_by"` // 排序字段
	Sorter  string `json:"sorter"`   // 升序, 降序

}
type Result struct {
	Page
	Data      interface{} `json:"data"`       // 数据
	Total     int         `json:"total"`      // 总条数
	TotalPage int         `json:"total_page"` // 总页数
	NextPage  int         `json:"next_page"`  // 上一页
	PrevPage  int         `json:"prev_page"`  // 下一页
}

func NewPage(page, size int, orderBy, sorter string) *Page {

	return &Page{
		PageNum: page,
		Size:    size,
		OrderBy: orderBy,
		Sorter:  sorter,
	}
}
func NewDefaultPage(page, size int) *Page {
	return &Page{
		PageNum: page,
		Size:    size,
		OrderBy: "id",
		Sorter:  "asc",
	}
}
func ExeMysqlPage(db *gorm.DB, model interface{}, page *Page, query interface{}, args ...string) (result *Result, err error) {
	var total int
	if err = db.Model(model).Where(query, args).Count(&total).Error; err != nil {
		return
	}

	totalPageNum := int(math.Floor(float64(total / page.Size)))
	offset := totalPageNum * (page.PageNum)
	order := page.OrderBy + " " + page.Sorter
	if err = db.Offset(offset).Limit(page.Size).Order(order).Where(query, args).Find(model).Error; err != nil {
		fmt.Println("err", err)

		panic(err)
		return
	}
	fmt.Println(model)
	var pervPage = page.PageNum + 1
	if pervPage > totalPageNum {
		pervPage = -1
	}
	result = &Result{
		Page:      *page,
		Data:      model,
		Total:     total,
		TotalPage: totalPageNum,
		NextPage:  page.PageNum - 1,
		PrevPage:  pervPage,
	}

	return
}
func PageHelper(condition *gorm.DB, model interface{}, page *Page) (result *Result, err error) {

	var total int
	if err = condition.Model(model).Count(&total).Error; err != nil {
		return
	}

	var nextPage int                                                       // 下一页
	var pervPage int                                                       // 上一页
	var totalPageNum = int(math.Ceil(float64(total) / float64(page.Size))) // 总页数

	if page.PageNum > totalPageNum {
		pervPage = -1
		nextPage = -1
	} else {
		offset := page.Size * (page.PageNum)
		order := page.OrderBy + " " + page.Sorter
		if err = condition.Offset(offset).Limit(page.Size).Order(order).Find(model).Error; err != nil {
			return
		}
		nextPage = page.PageNum - 1
		pervPage = page.PageNum + 1
		if pervPage+1 > totalPageNum {
			pervPage = -1
		}
	}
	// 结果
	result = &Result{
		Page:      *page,
		Data:      model,
		Total:     total,
		TotalPage: totalPageNum,
		NextPage:  nextPage,
		PrevPage:  pervPage,
	}
	return
}
