package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your name : ")
	name, err := in.ReadString('\n')
	// fmt.Println(i, err)  _ 언더스코어로 err를 안 보이게 처리해서 출력할 때도 안 보이게 할 수 있다.

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}
}
