package main

import "fmt"

/*
array_combine — 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
http://php.net/manual/zh/function.array-combine.php
 */
func main() {
	keys:=[]string{"name","age","job"}
	values:=[]interface{}{"fengahan",21,"coder"}
	
	m:=array_combine(keys,values)
	fmt.Println(m)
}
func array_combine(keys []string,values []interface{})map[string]interface{}{

	keys_count :=len(keys)
	if keys_count !=len(values) {
		panic("错误键和值数量不同")
	}
	m :=make(map[string]interface{},keys_count)
	for i:=0;i<keys_count;i++ {
		m[keys[i]]=values[i]
	}
	return m
}
