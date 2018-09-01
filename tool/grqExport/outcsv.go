package grqExport

import (
	"reflect"
	"errors"
	"encoding/csv"
)

/**
 csv导出操作对象
 */
type OutCsv struct {
	Head interface{}
	Data interface{}
}

/*
   返回 标题行对象的名称和字段名
 */
func (o *OutCsv) getHead() ([]string, []string, error) {
	var headTitle, headField []string
	t := reflect.TypeOf(o.Head)
	v := reflect.ValueOf(o.Head)
	if v.Kind() == reflect.Struct {
		for i, num := 0, v.NumField(); i < num; i++ {
			headField = append(headField, t.Field(i).Name)
			headTitle = append(headTitle, t.Field(i).Tag.Get("title"))
		}
	}
	return headTitle, headField, nil
}
/*
   返回 标题行对象的名称和字段名
 */
func (o *OutCsv) getBody() ([]map[string]string, error) {
	v := reflect.ValueOf(o.Data)
	vHead := reflect.ValueOf(o.Head)
	eles := make([]map[string]string, 0)
	if v.Kind() == reflect.Slice {
		for ii, dataNum := 0, v.Len(); ii < dataNum; ii++ {
			if !reflect.DeepEqual(v.Index(ii).Type(), vHead.Type()) {
				return nil, errors.New("数据格式错误")
			}
			dataEle:=reflect.TypeOf(v.Index(ii).Interface())
			ele := make(map[string]string)
			for i, num := 0, dataEle.NumField (); i < num; i++ {
				field:=dataEle.Field(i).Name
				ele[field] = v.Index(ii).Field(i).String()
			}
			eles = append(eles, ele)
		}
	}
	return eles, nil

}
/**
  写入标题行
 */
func (o *OutCsv) WriteHead(writer *csv.Writer) {
	headTitle, _, _ := o.getHead()
	writer.Write(headTitle)
}
/**
  内容写入csv文件中
 */
func (o *OutCsv) WriteBody(writer *csv.Writer) {
	_, headField, _ := o.getHead()
	csvBody, _ := o.getBody()
	for _, vv := range csvBody {
		ele := make([]string, 0)
			for _, v := range headField {
			ele = append(ele, vv[v])
		}
		writer.Write(ele)
	}
}
