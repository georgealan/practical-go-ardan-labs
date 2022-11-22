package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(safeDiv(1, 0))
}

// You can name return values
func safeDiv(a, b int) (q int, err error) {
	// q & r are local variables in safeDiv, just like (a & b)
	defer func() {
		if e := recover(); e != nil { // e's type is any or interface{} not error
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}
