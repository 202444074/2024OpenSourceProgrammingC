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
	if n <= 1 { // A prime number has only one of the natural numbers greater than 1 and itself as a divisor.
		isPrime = false
	} else {
		j := 2
		for j <= int(math.Sqrt(float64(n))) { // n이 일시적으로 float로 바뀌고 이를 다시 int로 바꿈, int로 바꾸면서 나머지를 버리게 되니까 <=으로 바꿈
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
