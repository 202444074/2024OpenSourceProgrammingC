package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("%f\n", math.Sqrt(19.0))   sqrt는 제곱근이라는 뜻
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
	if n <= 1 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { // 2를 제외한 짝수는 모두 소수가 아님
		isPrime = false
	} else {
		j := 2
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j)
			j++
		}
	}
	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
