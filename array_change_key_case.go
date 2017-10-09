package main

import (
	"strings"
	"fmt"
	"reflect"
)
/**
 *https://secure.php.net/manual/zh/function.array-change-key-case.php
 *array_change_key_case — 将数组中的所有键名修改为全大写或小写
 *go的数组没有键值 按照map来写 (后面所有情况都按照此情况处理)
 */
const CASE_UPPER =0 //大写
const CASE_LOWER =1 //小写 默认

func main() {
	m:=make(map[interface{}]interface{})
	m["Name"]="fengahan"
	m["job"]="coder"
	m[10]="age"

	new_map:=array_change_key_case(m,CASE_LOWER)
	fmt.Println(new_map)
}
func array_change_key_case(m map[interface{}]interface{},Do int)interface{}{
	//判断是不是字符串
	new_map :=make(map[interface{}]interface{})
	key:=""
	for k,v:=range m {
		ty:=reflect.TypeOf(k)
		k_temp:=reflect.ValueOf(k)
		if ty.String()=="string" { //如果是字符串则转换 否则为其他按照原来的返回
			if Do==CASE_LOWER {
				key=strings.ToLower(k_temp.String())
			}else if Do==CASE_UPPER {
				key=strings.ToUpper(k_temp.String())
			}else {
				panic("转换的格式错误")
			}
		}else{
			key=k_temp.String()
		}
		new_map[key]=v
	}
  	return new_map
}
