package main

import (
	"fmt"
)

// function with one return value
func user_sum(n int) (int) {
	// print the sum from 1 to n
	var sum int = 0
	var i int 	= 1
	for i < (n+1) {		// range: 0,...,(n-1)
		sum += i
		i++
	}
	return sum
}

func user_avg(n int) (float32) {
	// computes the arithmetic mean

	s := user_sum(n)	// call user sum
	return float32(s)/float32(n)	// convert to float

}

// using variadic functions
func print_nums(nums ...int) {
	// print out a variable list of integers
	fmt.Print(nums, " ")

	// calculate the total sum
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Print("\n", total, "\n")
}

// function closures
func integer_sequence() func () int {
	i := 0
	fmt.Println("before closure: i := ", i)
	return func() int {
		i++
		fmt.Println("in closure: i := ", i)
		return i
	}
}


func main() {
	n := 10	/*set the stop value*/

	u_sum := user_sum(n)
	fmt.Println("Sum\n----")
	fmt.Println(u_sum)

	avg := user_avg(n)
	fmt.Println("Avg\n----")
	fmt.Println(avg)

	print_nums(1,2,3)
	print_nums(1,2,3,4,5,6)

	// create an array(arrays are immutable)
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	print_nums(nums...)

	// initialize function closure
	nextInt := integer_sequence()
	fmt.Println(nextInt())



}
