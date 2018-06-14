package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

/***
	http://www.php.net/mt_rand
	生成指定长度的纯数字随机数
 */

func main()  {
	for i:=0;i<1000 ;i++  {
		e:=randInt(9)
		fmt.Println(e)
	}
}

/**
	generate  m to n rand number
 */
func randInt(strlen int64)int64  {
	if strlen >19 || strlen<1{
		panic("error index of return value")
	}
	fl:=float64(strlen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s:=fl
	v:=9*math.Pow(float64(10),s)
	number:= r.Int63n(int64(v))
	var l float64=0
	n:=number
	for n>10{
		n=(n/10)+10/9
		l++
	}
	if l<fl {
		y:=fl-l
		number=int64(math.Pow(float64(10),y))*number
	}else if l>fl {
		number=number/10
	}
	return number
}
