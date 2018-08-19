package main

import "fmt"

// Write a function with one variadic parameter that
// finds the greatest number in a list of numbers.

func main() {

	fmt.Println(largestNum(1, 2, 3, 4, 5))
	fmt.Println(largestNumNamedReturn(1, 2, 3, 4, 5))
	fmt.Println(largestNumMultipleNamedReturn(1, 2, 3, 4, 5))

}

func largestNum(nslice ...int) int {

	var largest int

	for _, n := range nslice {
		if n > largest {
			largest = n
		}
	}

	return largest
}

func largestNumNamedReturn(nslice ...int) (ans int) {
	for _, n := range nslice {
		if n > ans {
			ans = n
		}
	}

	return
}

func largestNumMultipleNamedReturn(nslice ...int) (loc int, num int) {
	for i, n := range nslice {
		if n > num {
			num = n
			loc = i
		}
	}
	fmt.Println("Largest num is ", num, "at index ", loc)

	return
}
