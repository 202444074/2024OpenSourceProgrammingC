package main

import (
	"fmt"
	"reflect"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.2}
	gpa_slice := gpa[1:4]                       // 파이선과 같은 슬라이스 뜻임
	fmt.Println(gpa, reflect.TypeOf(gpa))       // 배열
	fmt.Println(gpa, reflect.TypeOf(gpa_slice)) // 슬라이스
}
