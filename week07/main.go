package main

import (
	"fmt"
	"time"
)

func main() {
	// 오늘은 2024년 10월 11일 입니다.
	var now time.Time = time.Now()
	// var month time.Month = now.Month() October로 나옴
	// var month int = int(now.Month()) // 위 코드의 이유로 int로 묶어줌
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())
}
