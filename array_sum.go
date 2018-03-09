package main

import "fmt"

/**
	array_sum — array_sum—计算数组中值的总和。
	http://php.net/manual/en/function.array-sum.php
 */
func main() {
	var arr []float64
	arr=[]float64{100,12.58,18,16,14.3,17.6}
	fmt.Println(array_sum(arr))
}
func array_sum(arr []float64) float64{
	l:=len(arr)
	var s float64
	for i:=0;i<l;i++  {
		s+=arr[i]
	}
	return s
}
