package main

import "fmt"

func main() {
	var emptySlice []bool                               // bool로 만든 배열에는 true, false 만 들어갈 수 있는 배열임
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice)) // slice zero value (nil), 0
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, true)
	}
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice)) // []bool{true}, 1

	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	gpa[2] = 2.71
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpa)
}
