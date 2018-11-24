package main

import (
	L"./lib"
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

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}


// go pointers
/*
go pointers allow for you to pass
references to values and records within
your program
 */


func zeroval(ival int) {
	// takes the value
	ival = 0
}

func zeroptr(iptr *int) {
	// takes an int pointer then dereferences it
	*iptr = 0
}


// go structs
type person struct {
	name string
	age  int
	job  string
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

	// calling factorial
	fmt.Println(factorial(6))

	// using go pointers
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)		// takes in the memory address of i
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)	 // output memory location


	// setup for go structs
	fmt.Println(person{"William Murphy",
						31,
						"Machine Learning Engineer"})

	// initializing a struct like a python dictionary
	fmt.Println(person{name: "James", age: 26, job: "Artist"})

	// using pointers with structs
	s := person{name: "Melissa",
				age: 25,
				job: "Teacher"}

	// sp is a pointer to a struct
	sp := &s	// assign address of s to sp

	fmt.Println(sp.job)
	fmt.Println(sp.name)
	fmt.Println(sp.age)
	sp.name = "Melissa Brook Lin"
	fmt.Println(s.name)

	file_ := "/Users/williamrobertmurphy/Desktop/housing_data.rtf"




}
