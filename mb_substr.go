/**
mb_substr — 获取部分字符串
http://php.net/manual/zh/function.mb-substr.php
在某公司过了试用期后，我加入了公司的群，按照男生的天性就是去群里找妹子，无意中 发现了你，但是我当时看中的
是你身旁的那个妹子，后来我看到了你在休息时间在看老友记。让我对你产生了兴趣
 */
package main

import "fmt"

func main() {
	str:="ohmygod 我的天啊"
	fmt.Println(mb_substr(str,1,11))
}

func mb_substr(str string ,s int,e int)string{
	str_rune:=[]rune(str)
	str_len:=len(str_rune)

	if s>=e || e>str_len {
		panic("长度有误")
	}

	return string(str_rune[s:e])

}




