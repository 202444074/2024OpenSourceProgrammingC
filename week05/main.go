package main

import (
	"fmt" // c언어의 #include <stdio.h> 역할 정의를 해주지 않아도 밑에서 사용하고 저장하면 자동으로 생성됨
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int = 55
	//var f float32 = 3.99

	// var i int
	// i = 55

	i := 55
	//i := "55"
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("i는 %d\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)
}
