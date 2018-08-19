package main

import "fmt"

/**
  Write a function, foo, which can be called in all of these ways:
  foo(1, 2)
  foo(1, 2, 3)
  aSlice := []int{1, 2, 3, 4}
  foo(aSlice...)
  foo()
*/
func main() {

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

// Just needs to accept variadic args of int
func foo(nargs ...int) {

	if len(nargs) < 1 {
		fmt.Println("0 args passed to func foo")
	} else {

		for _, n := range nargs {

			if len(nargs) == 1 {
				fmt.Println("The only arg passed to foo is ", n)
			} else {
				fmt.Println(n, " was passed to foo")
			}
		}
	}
}
