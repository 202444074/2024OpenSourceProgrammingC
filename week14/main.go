package main

import (
	"fmt"
)

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}
	student.id = 202444074
	student.name = "이영희"
	student.gpa = 4.5
	fmt.Println(student.gpa)
	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 20241234
	student2.name = "아이다"
	student2.gpa = 4.43
	fmt.Println(student2.id)

}
