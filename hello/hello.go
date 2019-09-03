// hello world in Go
package main

import (
	"fmt"
	"math"
)

func plus(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, 世界")
	//operations -----
	var a = "lang"
	fmt.Println("go" + a)
	var b, c int = 1, 1
	//c declared and not used ... all variables must be used.
	fmt.Println("1+1 =", b+1)
	fmt.Println("1+1 =", c+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	var d = true
	fmt.Println(!d)
	var e int
	fmt.Println(e) // initial value = 0
	f := "apple"
	fmt.Println(f)

	// maths
	const n = 500000000
	fmt.Println(math.Sin(n))

	//loops
	for n := 0; n <= 5; n++ { // prime numbers
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	//arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Ranges ...
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//functions -- I cannot have functions inside other functions
	res := plus(2, 3)
	fmt.Println("plus(2,3) = ", res)

}
