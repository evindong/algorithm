package main

import "fmt"

func main() {
	n := 5
	fmt.Print(fib1(n))
	fmt.Println()
	f := fib2()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}

}

//method1
func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

//method2
func fib2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//1 1 2 3 5 8 13
