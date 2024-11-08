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

func getInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		//log.Fatal(err)  이 코드는 에러를 검출하면 바로 코드를 중단시키고 에러를 출력함
		return 0, err // 이 코드는 에러를 리턴해서 메인에서 에러를 처리함
	}
	return n, nil
}

func main() {
	fmt.Print("input start number : ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("input end number : ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
