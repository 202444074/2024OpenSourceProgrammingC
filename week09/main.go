package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Printf("%f 나누기 %f은(는) %f입니다.\n", 1.0, 3.0, 1.0/3.0)    console 출력
	result := fmt.Sprintf("%-5.1f 나누기 %5.1f은(는) %0.1f입니다.\n", 1.0, 3.0, 1.0/3.0) // 서식을 문자열로 리턴
	fmt.Print(result)

	i := 1
	fmt.Printf("%T\n", i) // 다른 방법은 Println(reflect.Typeof(i))
	for i <= 10 {         // while과 똑같은 역할
		//fmt.Printf("%2d\n", i)
		fmt.Printf("%4d\n", rand.Intn(1000)+1) // %4d로 일의 자리수 맞추기
		i++
	}
}
