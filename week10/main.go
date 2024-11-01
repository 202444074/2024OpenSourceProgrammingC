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
			}
			j++
		}
	}
	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
