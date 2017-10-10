package main

import "fmt"

/***
array_column — 返回数组中指定的一列
http://php.net/manual/zh/function.array-column.php
 */
func main() {

	m:=[]map[string]interface{}{
		{"id":10,"first_name":"feng","last_name":1},
		{"id":11,"first_name":"aaa","last_name":2},
		{"id":12,"first_name":"han","last_name":2},
		}

  first_arr:=array_column(m,"last_name")
  fmt.Println(first_arr)
}

func array_column(m []map[string]interface{},needle_key string)[]interface{}{
	var arr []interface{}
	for _,v:=range m  {
		for k,mv:= range v{
			if k ==needle_key {
				arr=append(arr,mv)
			}
		}
	}
   return  arr
}