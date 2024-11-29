package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	// visitors : 구조체 슬라이스
	totalCost := 0
	for _, value := range visitors {
		totalCost = totalCost + value.cost
	}
	return totalCost
}

func main() {
	var numVisitors int
	fmt.Println("How many visitores? ")
	fmt.Scanln(&numVisitors)

	vs := make([]visitor, numVisitors)

	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Print("Input age: ")
		fmt.Scan(&age)

		switch { // switch로 바꾸는 방법
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 7000}
		}
		/*
			if age < 12 {
				vs[i] = visitor{age: age, cost: 5000}
			} else if age >= 12 && age < 65 {
				vs[i] = visitor{age: age, cost: 10000}
			} else {
				vs[i] = visitor{age: age, cost: 7000}
			}
		*/
	}
	fmt.Printf("Total price is %d won", calculateCost(vs))
}
