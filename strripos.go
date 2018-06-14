package main

import (
	"fmt"
	"regexp"
)

/**
http://php.net/manual/zh/function.strripos.php
strripos — 计算指定字符串在目标字符串中最后一次出现的位置（不区分大小写）
 **/
func main() {
    var stringIng  ="计算指定hello你好字符串在目标字符串中最后一次出现hello你好的位置"
    fmt.Println(len(stringIng))
    var needle ="hello你好"
	ret:=strripos(stringIng,needle)
	fmt.Println(ret)
}
func strripos( s string,e string)int  {
	r, _ := regexp.Compile(e)
	se:=r.FindAllStringSubmatchIndex(s,-1)
	return se[len(se)-1][1]
}
