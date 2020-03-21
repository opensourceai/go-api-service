/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package page

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
)

type Page struct {
	PageNum int    `json:"page" valid:"Min(0)"` // 页数
	Size    int    `json:"size" valid:"Min(1)"` // 条数
	OrderBy string `json:"order_by"`            // 排序字段
	Sorter  string `json:"sorter"`              // 升序, 降序

}
type Option struct {
	//Data      interface{} `json:"data"`       // 数据
	Total     int    `json:"total"`      // 总条数
	TotalPage int    `json:"total_page"` // 总页数
	NextPage  int    `json:"next_page"`  // 上一页
	PrevPage  int    `json:"prev_page"`  // 下一页
	Offset    int    `json:"-"`          // 偏移
	Limit     int    `json:"-"`          // 条数
	Order     string `json:"-"`          // 排序规则
}

type Result struct {
	*Page
	Data interface{} `json:"data"` // 数据
	*Option
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

// 单表查询分页工具函数
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
		Page: page,
		Data: model,
		Option: &Option{
			Total:     total,
			TotalPage: totalPageNum,
			NextPage:  nextPage,
			PrevPage:  pervPage,
		},
	}
	return
}

// 从gin.Context绑定Page
func BindPage(context *gin.Context) (p *Page) {
	var err error
	// 第几页
	pageNum := context.Query("page")
	var pageNumInt int

	if pageNum == "" {
		pageNumInt = 0
	} else {
		if pageNumInt, err = strconv.Atoi(pageNum); err != nil {
			panic("pageNum 无法转换.")
		}
	}
	// 每页数据个数
	pageSize := context.Query("size")
	var pageSizeInt int

	if pageSize == "" {
		pageSizeInt = 20
	} else {
		if pageSizeInt, err = strconv.Atoi(pageSize); err != nil {
			panic("pageSize 无法转换.")
		}
	}

	// 排序字段
	sorter := context.Query("sorter")
	if sorter == "" {
		sorter = "asc"
	}

	// 排序方式
	orderBy := context.Query("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}
	p = &Page{
		PageNum: pageNumInt,
		Size:    pageSizeInt,
		OrderBy: orderBy,
		Sorter:  sorter,
	}
	return
}

// 返回分页option属性,
func PageHelperOption(condition *gorm.DB, model interface{}, page *Page) (result *Option, err error) {

	var total int
	if err = condition.Model(model).Count(&total).Error; err != nil {
		return
	}

	var nextPage int                                                       // 下一页
	var pervPage int                                                       // 上一页
	var totalPageNum = int(math.Ceil(float64(total) / float64(page.Size))) // 总页数
	var offset int
	var order string
	if page.PageNum > totalPageNum {
		pervPage = -1
		nextPage = -1
	} else {
		offset = page.Size * (page.PageNum)
		order = page.OrderBy + " " + page.Sorter

		nextPage = page.PageNum - 1
		pervPage = page.PageNum + 1
		if pervPage+1 > totalPageNum {
			pervPage = -1
		}
	}
	// 结果
	result = &Option{
		Total:     total,
		TotalPage: totalPageNum,
		NextPage:  nextPage,
		PrevPage:  pervPage,
		Offset:    offset,
		Limit:     page.Size,
		Order:     order,
	}
	return
}
