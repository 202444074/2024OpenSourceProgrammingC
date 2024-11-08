package greeting

import "fmt"

func hi(name string) {
	fmt.Printf("Hi %s~\n", name)
}

func hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func EnglishGreetings(name string) { // 함수를 소문자로 쓸 때 사용하는 경우
	hello(name)
	hi(name)
}
