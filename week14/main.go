package main

import (
	"fmt"
	// go get 해야함
)

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("what's ur name? (exit to 'q') ")
		fmt.Scanln(&name) // 이름이 같을 수 도 있어서 map은 못씀 그래서 전에 쓴 slice 코드를 이용해야 함
		if name == "q" || name == "Q" {
			break
		}
		fmt.Print("Ur age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}
	for name, age := range ages {
		fmt.Printf("%s is %d year old.\n", name, age)
	}
}
