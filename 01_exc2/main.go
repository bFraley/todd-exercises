package main

import "fmt"

func main() {
	ans1, ans2 := twoReturns(10) // 5, true
	fmt.Println(ans1, ans2)

	ans3, ans4 := twoNamedReturns(20) // 10, true
	fmt.Println(ans3, ans4)

	ans5, ans6 := twoReturnsWithFloat(11) // 5.5 false
	fmt.Println(ans5, ans6)

	// With Func Expression 02_exc2
	func_exp := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(func_exp(10))
	fmt.Println(func_exp(20))

}

func twoReturns(n int) (int, bool) {
	half := n / 2
	even := n%2 == 0
	return half, even
}

// With named returns
func twoNamedReturns(n int) (half int, even bool) {
	half = n / 2
	even = n%2 == 0
	return
}

// With Float Division
func twoReturnsWithFloat(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}
