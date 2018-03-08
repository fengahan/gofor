package main

import (
	"bytes"
	"fmt"
)

/**
https://secure.php.net/manual/zh/function.str-repeat.php
str_repeat — 重复一个字符串
 **/
func main() {
	str := "abc"
	new_str := str_repeat(str, 6) //	strings.Repeat()可用此函数代替
	fmt.Println(new_str)
}
func str_repeat(str string, mult int) string {
	new_str := ""
	if mult < 1 {
		return new_str
	}
	var buffer bytes.Buffer
	for i := 0; i < mult; i++ {
		buffer.WriteString(str)
	}
	new_str = buffer.String()
	return new_str
}
