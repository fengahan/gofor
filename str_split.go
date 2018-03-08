package main

import (
	"fmt"
	"strings"
)

/***
str_split — 将字符串转换为数组
http://php.net/manual/zh/function.str-split.php
   想你的时候就写一个实例，现在在上海，没有找到工作，如果我没猜错的话你应该也在这个城市，曾经好几次坐地铁，都会想象在地铁上遇到你
的情景，但是并没有。又或许你根本不在这里。
*/

func main() {
	str := "我该怎么办呢万能的表情,I miss you"
	str_arr := str_split(str, 4)
	fmt.Println(str_arr)

}
func str_split(str string, length int) []string {

	str = strings.Join(strings.Fields(str), "")
	strs := []rune(str)
	c := len(strs)
	arr := []string{}
	if length < 1 || length >= c {
		return []string{str}
	}
	for i := 0; i < c; i += length {
		arr = append(arr, string(strs[i:i+length]))
	}
	return arr
}
