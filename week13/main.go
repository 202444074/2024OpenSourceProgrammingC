package main

import "fmt"

func main() {
	var emptySlice []bool
	//emptySlice = make([]bool, 5)
	fmt.Printf("%#v\n", emptySlice) // slice의 기본값은 nil임

	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	//gpa_slice[1] = 2.71  여기 주석이 풀리면 원본 배열도 바뀜, gpa[2]가 바뀜(이 상태에서는 slice1이 gpa2를 가져온거니까)
	gpa[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpa)
	fmt.Printf("%#v\n", gpa_slice) // s%#v 가 go코드에서 어떻게 작성될지를 보여준다는데 잘 모르겠음
}
