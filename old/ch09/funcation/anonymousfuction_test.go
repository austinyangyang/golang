package function

import (
	"fmt"
	"strings"
	"testing"
)

func addr2(x int) func(int) int {
	// var x int
	return func(y int) int {
		x += y
		return x

	}
}

// func make

func TestAnonymousFuncation(T *testing.T) {

	add := func(x, y int) {
		fmt.Println("x + y")

	}
	add(10, 20)
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	var f = addr2(10)
	fmt.Println("f value is: ", f(10))
	fmt.Println("f value is: ", f(20))

}

func makeSuffixFunc(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}

func TestMakeSuffix(t *testing.T) {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(txtFunc("test"))

	f1, f2 := calcs(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))

}

func calcs(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base

	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
