package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
func test(strs ...string) { // ...은 임의의 개수의 문자열 인자를 전달할 수 있다는 뜻
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	slices := os.Args[1:]
	fmt.Println(slices[1])
	for _, slice := range slices {
		fmt.Println(slice)
	}
	slices = append(slices, "forever", "!")
	fmt.Println(slices, len(slices))
	test("abc")
	test("abc", "123")
	test() // 모두 string인 모습
	test("abc", "123", "inha")
}
