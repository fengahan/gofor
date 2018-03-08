package main

import "fmt"

/**
	array_keys — 返回map中的key，不返回子集mao的key
 */
func main() {

	var m  map[string]interface{}
	m=map[string]interface{}{"name":"feng","age":24,"sex":"男"}
	fmt.Println(array_keys(m))
}
/**
	返回一个数组
 */
func array_keys(m map[string]interface{}) []string {
	var arr []string
	for k:=range m{
		arr=append(arr,k)
	}
	return arr
}
