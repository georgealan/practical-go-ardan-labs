package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                // s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"
	if s == nil {              // You can compare only a slice to nil, for compare a slice to another you have to write you own code.
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half-open range
	fmt.Printf("s3 = %#v\n", s3)

	//fmt.Println(s2[:100]) // panic

	s3 = append(s3, 100)
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2) // s2 is changed as well
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))
	fmt.Println()

	var s4 []int
	//s4 := make([]int, 0, 1_000) // Single allocation
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))
	fmt.Println()

	fmt.Println("Solving the program exercise")
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // Should return [A B C D E]
	fmt.Println()

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 3, 4}
	fmt.Printf("Before median: %v\n", vs)
	fmt.Println(median(vs))
	fmt.Printf("After median: %v", vs)

	fmt.Println("\nReturn an error:")
	fmt.Println(median(nil))
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Median of empty slice")
	}
	// Copy in order not to change values
	nums := make([]float64, len(values))
	copy(nums, values)

	/*
		Change the values variable to nums for avoid change values in pointers with this approach.
		In Go, you have to identify with types are pointers to avoid this type of behavior
		which cause changes that affect the same object, naturally slices have arrays inside
		and these are pointers.
	*/
	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}

func concat(s1, s2 []string) []string { // The type in params is the same for s1 and s2
	// Restriction: No for loops
	//var ar = make([]string, len(s1)+len(s2))
	ar := make([]string, len(s1)+len(s2))
	copy(ar, s1)
	copy(ar[len(s1):], s2) // Beginning in the end of ar slice to copy after it.
	return ar
}

func appendInt(s []int, v int) []int {
	// Here we make sure that we have an array with enough capacity
	i := len(s)
	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-allocate and copy
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	// Then we can do this
	s[i] = v
	return s
}
