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
	fmt.Print("input number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	if n <= 1 { // A prime number has only one of the natural numbers greater than 1 and itself as a divisor.
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false
				break // break 하나로 큰 향상을 이룸
			}
			fmt.Printf("%d ", j) // check j loop, break 없을 때와 있을 때를 비교하기 위한 문장, 이 부분을 break위에 넣는게 좋아보임
			j++
		}
	}
	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
