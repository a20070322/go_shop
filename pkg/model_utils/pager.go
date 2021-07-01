package model_utils

import (
	"github.com/a20070322/shop-go/ent"
	"github.com/cjrd/allocate"
	"math"
	"reflect"
)

type PageOptions struct {
	Page int `json:"page" form:"page"; query:"page"`
	Size int `json:"size" form:"size"; query:"size"`
}

type PageSort struct {
	SortColumn string `json:"sort_column" form:"sort_column"`
	SortOrder  string `json:"sort_order" form:"sort_order"`
}

func PageSortDB(client interface{}, pageSort PageSort) {
	t := reflect.ValueOf(client)
	if pageSort.SortOrder == "ascend" && pageSort.SortColumn != "" {
		params := []reflect.Value{reflect.ValueOf(ent.Asc(pageSort.SortColumn))}
		t.MethodByName("Order").Call(params)
	}
	if pageSort.SortOrder == "descend" && pageSort.SortColumn != "" {
		params := []reflect.Value{reflect.ValueOf(ent.Desc(pageSort.SortColumn))}
		t.MethodByName("Order").Call(params)
	}
}

func GetDefaultPager(num int, defaultNum int) int {
	if num <= 0 {
		return defaultNum
	}
	return num
}


func PipePagerFn(dataSource interface{}) {
	allocate.Zero(dataSource)
	t := reflect.ValueOf(dataSource).Elem()
	page := t.FieldByName("Page")
	size := t.FieldByName("Size")
	if page.Int() < 1 {
		page.SetInt(1)
	}
	if size.Int() < 1 {
		size.SetInt(10)
	}
}

func PipeLimitFn(db interface{}, dataSource interface{}) {
	data := reflect.ValueOf(dataSource).Elem()
	page := data.FieldByName("Page").Int()
	size := data.FieldByName("Size").Int()
	offset := (page - 1) * size
	refDb := reflect.ValueOf(db)
	refDb.MethodByName("Limit").Call([]reflect.Value{reflect.ValueOf(int(size))})
	refDb.MethodByName("Offset").Call([]reflect.Value{reflect.ValueOf(int(offset))})
}


func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}