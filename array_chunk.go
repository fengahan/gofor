package main

import (
	"fmt"
	//"math"
	"math"
)

/**
	https://secure.php.net/manual/zh/function.array-chunk.php
    array_chunk — 将一个数组分割成多个 目前不支持是否 保留键值顺序 参数
*/

func main() {
	array := []interface{}{"a", "b", "c", "d", 10, 15, 16, "e", "f"}
	fmt.Println(array_chunk(array, 2))

}
func array_chunk(arr []interface{}, size int) interface{} {

	array_count := len(arr)
	if array_count-2 < size || array_count-1 == size {
		return arr
	}
	s_len := math.Ceil(float64(array_count / size))
	m := make(map[int]interface{}, int(s_len)) // 新的map长度
	e := 1
	l := size
	for i := 0; i < array_count; i += size {
		if i+size > array_count {
			l = array_count
		} else {
			l = i + size
		}
		m[e] = arr[i:l]
		e++
	}
	return m
}
