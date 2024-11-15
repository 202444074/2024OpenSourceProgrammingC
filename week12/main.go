package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpas [3]float64 // github.com/headfirstgo/keyboard 에서 가져올 함수에서 float를 리턴해서 float64로 만듬
	for i := 0; i < len(gpas); i++ {
		fmt.Print("Input float number : ")
		gpas[i], _ = keyboard.GetFloat() // float와 err가 들어가야하는데 에러가 안 난다고 가정하고 _를 넣음
	}
	for _, gpa := range gpas {
		fmt.Println(gpa)
	}
}
