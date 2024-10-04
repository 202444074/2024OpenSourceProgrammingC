package main

import (
	"fmt"
)

func main() { //v0.6 unicode 확인할 것
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i) // zero values
	f = 2.7
	i = 3
	fmt.Print("\n\n", f <= float64(i), "\n") // true false로 나오는 것이 bool
	// go build main.go -> main.exe 생성 -> ./main.exe -> week06 main.go 실행
}
