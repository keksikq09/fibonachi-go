package main

import "fmt"

// fibonacci це функція, що повертає
// функцію, що повертає int.
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		temp := a
		a = b
		b = temp + b
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
