package main

import "fmt"

func main() {
	var i any // rule of thumb: don't use any
	// If use Go < 1.18
	// var i interface{}

	i = 7
	fmt.Println(i)

	i = "George"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println("s:", s)

	/*
		n := i.(int) // will panic
		fmt.Println("n:", n)
	*/

	// comma, ok
	n, ok := i.(int) // won't panic
	if ok {
		fmt.Println("n:", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}

	fmt.Println(max([]float64{3, 1, 2}))
	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]int{3, 9, 7}))
}

// Go generics, with use of interface
func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]

	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

type Number interface {
	int | float64
}
