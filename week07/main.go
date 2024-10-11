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
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                // 줄바꿈등을 제거(줄바꿈 \n이 에러를 발생), 파이썬의 strip과 비슷함
	score, _ := strconv.ParseInt(i, 10, 32) // 정수형으로 변환 10진수 32bit
	var grade string
	if score >= 90 {
		grade = "A" // 위에서 grade 전역 변수를 정의 해줘서 :=가 아닌 =로 바로 넣으면 됨
	} else {
		grade = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다.\n", score, grade)
}
