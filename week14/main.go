package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile" // go get 해야함
)

func main() {
	lines, err := datafile.GetStrings("votes.txt") // votes.txt 만들어야함, 여기서 lines는 GetString에 의해 slice 객체임
	if err != nil {
		log.Fatal(err)
	}
	/*
		var names []string
		var counts []int
		for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
		}
		for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
		}
		밑의 코드를 다른 방법으로 하는 방법, 책에 있음, map이 더 간편하지만 이 다른 방법을 아는 것도 중요하다고 강요하심
	*/
	counts := make(map[string]int) // Map 객체인 counts를 만듦(파이선의 딕셔너리라고 함) key와 value가 들어있음
	for _, line := range lines {
		counts[line]++ // 아무것도 없는 counts의 key에 line이 들어감, value에 zero value 0이 들어있었는데 ++로 1로 바뀜, 똑같은 key가 들어오면 딕셔너리처럼 똑같은 key에 ++됨
	} // 이 for문으로 txt가 모두 들어감
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
