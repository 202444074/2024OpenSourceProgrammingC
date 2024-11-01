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
	// var isPrime bool = true / 리턴으로 true false 받을 거니까
	if n <= 1 {
		//isPrime = false
		return false
	} else if n == 2 {
		//isPrime = true
		return true
	} else if n%2 == 0 {
		//isPrime = false
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				// isPrime = false
				// break
				return false
			}
			fmt.Printf("%d ", j)
			j += 2
		}
	}
	return true // 모든 if문을 벗어난 소수임
}

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
	if isPrime(n) {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
