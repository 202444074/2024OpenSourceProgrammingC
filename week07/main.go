package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input your score : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                // 줄바꿈등을 제거(줄바꿈 \n이 에러를 발생), 파이썬의 strip과 비슷함
	score, _ := strconv.ParseInt(i, 16, 32) // 정수형으로 변환 16진수 32bit
	if score >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", score)
	}
}
