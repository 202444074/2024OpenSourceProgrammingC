package main

import (
	"fmt"
)

type student struct { // 코드가 중복되는 것을 type로 바꿀 수 있음
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444074
	student1.name = "이영희"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)
	var student2 student
	student2.id = 20241234
	student2.name = "아이다"
	student2.gpa = 4.43
	fmt.Println(student2.id)

}
