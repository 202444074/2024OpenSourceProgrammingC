package main

import "fmt"

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4] // 4.1, 4.5, 3.9  1부터 4전까지(파이선과 똑같음)
	//pga_slice[1] = 2.71
	gpa[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpa)
}
