package main

import (
	//"strings"
	"fmt"
	"strings"
	"bytes"
)

/**
 str_replace — 子字符串替换
 https://secure.php.net/manual/zh/function.str-replace.php

 **/
func main() {

	s:="fengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahanfengahan"
	r:="a"//把a替换成阿
	nw:="阿"
	new_str:=str_replace(r,nw,s)
	fmt.Println(new_str)
}
/**
	r 要替换的字符
	nw 被替换为nw
	s 被替换的字符串
 */
func str_replace(r string,nw string,old_str string)string  {

	//string.Counts 某字符串在目标字符串出现的次数
	//这里使用string包的函数 后续会单独封装这个函数
	rc:=strings.Count(r,old_str)
	sc:=len(old_str)
	s_len:=len(nw)+len(old_str)-rc
	s:=make([]string,s_len)
	for i:=0;i<sc;i++{
		s[i]=string(old_str[i])
		if string(old_str[i])==r {
			s[i]=nw
		}
	}
	var b bytes.Buffer

	for i:=0;i<s_len;i++ {
		b.WriteString(string(s[i]))
	}
	return b.String()
}
