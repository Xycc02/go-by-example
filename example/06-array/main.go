package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])   // 0
	fmt.Println("len:", len(a)) // 5

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) // 1 2 3 4 5

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/**
	0 1 2
	1 2 3
	*/
}
