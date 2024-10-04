package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 13
	f := 12.9
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)   c언어와 다르게 go는 묵시적 형변환 안함
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i))

}
