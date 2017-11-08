package main

import (
	"fmt"
)

/**array_diff_key — 使用键名比较计算数组的差集
*  http://php.net/manual/zh/function.array-diff-key.php
*
*没时间写这些代码了，心情又特别糟糕，所以建了几个壳子放在这里 以后补上
*希望以后很多人可以看到这这些代码，也看到我对你的独白，想起来也是有趣
*你在我的幻想中出现，又在我的思念中离去
*
**/
func main() {
	var(
		m1 map[string]int
		m2 map[string]int
	)
	m1=make(map[string]int)
	m2=make(map[string]int)
	m1["a"]=1;
	m1["b"]=2;
	m1["c"]=3;
	m2["a"]=4;
	m2["d"]=5;
	fmt.Println(array_diff_key(m1,m2))
}
func array_diff_key(m1,m2 map[string]int) map[string]int {

	m3:=make(map[string]int)
	for key,val:=range m1{
			if _,ok:=m2[key];!ok{
				m3[key]=val
			}
	}
	return m3;
}
