package panicnil

import "fmt"

func Panic() {
	panic(nil)
}

func Catch() {
	defer func() {
		r := recover()
		fmt.Printf("type: %T\nvalue: %v\nis nil: %v\n", r, r, r == nil)
	}()

	panic(nil)
}
