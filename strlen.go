/**
写代码是无聊的 我决定每一个实例都会在头部注释发表一些有关于你的个人心情
希望有天你能看到
 */
package main
//http://php.net/manual/zh/function.strlen.php
//strlen统计字符串的长度
import "fmt"
func main() {
	str:=[]rune("我要成为程序员的精英 I want to do something");
	c:=strlen(str)
	fmt.Println(c)
}

func strlen(str []rune)int{
	e:=0
	for _,v:=range str{
		if string(v)!="32" {
		  e++
		}
	}
    return e
}