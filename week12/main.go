package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	fmt.Println(scores[1], scores[0]) // , scores[3]) 인덱스 범위를 벗어나서 오류 발생
	fmt.Printf("%#v\n", scores)       // %#v는 배열 리터럴 형태로 보여줌
	fmt.Println(scores)               // 이거는 그냥 배열이 출력

	// var dates [3]time.Time
	// dates[1] = time.Unix(194720000, 0)
	// fmt.Println(dates[1])

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)} // := 사용 가능
	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Printf("%#v\n", dates)
	fmt.Println(dates)
}
