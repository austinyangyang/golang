package function

import (
	"fmt"
	"testing"
)

type calculation func(int, int) int

func add(x, y int) int {
	return x + y

}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)

}

func do(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		// err := errors.New("invalid")
		return sub

	}
}

func TestFunction(t *testing.T) {
	var c calculation

	c = add
	fmt.Printf("type of c:%T\n", c)
	fmt.Println(c(1, 2))

	f := add
	fmt.Printf("type of f:%T\n", f)
	fmt.Println(f(10, 20))

	ret2 := calc(10, 20, add)

	fmt.Println(ret2)
	// var ret3 calculation
	ret4 := do("+")

	fmt.Printf("type of ret3:%T\n", ret4)

}
