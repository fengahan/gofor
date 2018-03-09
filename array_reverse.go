package main

import "fmt"

/**
array_reverse -以相反的顺序返回一个数组。
http://php.net/manual/en/function.array-reverse.php
*/
func main() {
	arr := []interface{}{"golang", 2017, []string{"red", "green"}, 15, "2018"}
	fmt.Println(array_reverse(arr))
}

func array_reverse(arr []interface{}) []interface{} {
	l := len(arr)
	var new_arr []interface{}
	for i := 0; i < l; i++ {
		new_arr = append(new_arr, arr[l-i-1])
	}
	return new_arr
}
