package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args[1:], len(os.Args))
	slices := os.Args[1:] //는 첫번째 프로그램 이름을 제외한 나머지 인자를 가져온다는 뜻
	fmt.Println(slices[1])
	for _, slice := range slices {
		fmt.Println(slice)
	}
	slices = append(slices, "forever", "!")
	fmt.Println(slices, len(slices))
}
