//go routines

package main

import "fmt"

// Variadic functions_ f() with N number of arguments (print)
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closures --
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	//sequences ...
	nextInt := intSeq() //1
	fmt.Println(nextInt())
	fmt.Println(nextInt()) //2
	newInts := intSeq()    // Start over 1
	fmt.Println(newInts())
	// test

}
