package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13              //var i int = 13
	var f float64 = 12.9 //f := 12.9
	c1 := 'Z'            // 99
	c2 := '김'            // 44608, unicode

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)   c언어와 다르게 go는 묵시적 형변환 안함
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) //밑의 type를 보면 i가 일시적으로만 변한것
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) //i와 f가 int지만 %f에서 오류가 나서 %d로 바꿈
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

}
