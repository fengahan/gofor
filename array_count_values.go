package main

import (
	"fmt"
)

/**

array_count_values — 统计数组中所有的值
array_count_values() 返回一个数组： 数组的键是 array 里单元的值； 数组的值是 array 单元的值出现的次数。
https://secure.php.net/manual/zh/function.array-count-values.php
*/
func main() {

	arr := []interface{}{1, "hello", 1, 1, "world", "hello"}
	m := array_count_values(arr)
	fmt.Println(m)
}

func array_count_values(arr []interface{}) map[interface{}]int {
	l := len(arr)
	m := make(map[interface{}]int, l)
	for i := 0; i < l; i++ {
		m[arr[i]] = 1
		for e := 0; e < i; e++ {
			if arr[i] == arr[e] {
				m[arr[i]]++
			}
		}
	}
	return m
}
