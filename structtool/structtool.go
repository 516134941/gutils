// Package structtool 结构体工具
package structtool

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// StructFloatFormat 格式化结构体内的float64,保留2位小数
// Format the float64 in the structure, retaining 2 decimal places
func StructFloatFormat(obj interface{}) interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Interface().(type) {
		case float64:
			formatFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.Field(i).Interface().(float64)), 64)
			v.Field(i).Set(reflect.ValueOf(formatFloat))
		default:
			continue
		}
	}
	return obj
}

func deepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}

// StructCopy 结构体相同字段拷贝
func StructCopy(dst interface{}, src interface{}) {
	srcv := reflect.ValueOf(src)
	dstv := reflect.ValueOf(dst)
	srct := reflect.TypeOf(src)
	dstt := reflect.TypeOf(dst)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := deepFields(reflect.ValueOf(src).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

// StructURLDecode 将接受的请求参数进行url-decode处理
// Will accept the request parameters for url-decode processing
func StructURLDecode(objReq interface{}) interface{} {
	t := reflect.TypeOf(objReq)
	v := reflect.ValueOf(objReq)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Interface().(type) {
		case string:
			decodeString, _ := url.QueryUnescape(v.Field(i).String())
			v.Field(i).Set(reflect.ValueOf(decodeString))
		default:
			continue
		}
	}
	return objReq
}
