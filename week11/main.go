package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j)
			j += 2
		}
	}
	return true
}

func main() {
	fmt.Print("input start number : ")
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("input end number : ")
	b, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ { // 스타팅 n1 마지막 값 n2
		if isPrime(j) { // 소수만 출려하기 때문에 else 없애고 true일 때만 출력
			fmt.Printf("%d ", j)
		}
	}
}
