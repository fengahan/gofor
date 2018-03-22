package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
str_shuffle() 函数打乱一个字符串，使用任何一种可能的排序方案。
https://secure.php.net/manual/zh/function.str-shuffle.php
*/
func main() {
	str := []rune("不，这些都还不够，我必须是你近旁的一株木棉，作为树的形象和你站在一起。")
	str_len := len(str)
	var s []string
	iarr := generateRand(str_len)
	for i := 0; i < str_len; i++ {
		s = append(s, string(str[iarr[i]]))
	}
	fmt.Println(s)
}
/***
生成0到len的不重复随机数
可使用rand下的 perm函数
*/
func generateRand(l int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := make([]int, l)
	for i := 1; i < l; i++ {
		j := r.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}
