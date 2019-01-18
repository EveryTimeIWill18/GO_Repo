package DataStructs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Open_file(file_path string, delim string)([]string) {
	// Open a file for data extraction

		file, err := os.Open(file_path)

		if err != nil { // an error occurred
			panic(err)
		}

		// read the entire contents of the file
		raw_data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		_ = file.Close() // close the file

		parsed_data := strings.SplitAfter(string(raw_data), delim)

		return parsed_data
}

func data_pointer(dp *string) {
	// pointer to the current data source
	*dp = " "
}


func Open_web_file(url string) {
	// Make a GET request to specified url


	r, err := http.Get(url)	// make get request
	if err != nil {
		panic(err)	// an error was raised
	}
	defer r.Body.Close()

	if r.Status == "200 OK" {
		fmt.Println(r.Status)
		fmt.Println("Successfully connected to url")
		goto End

	} else if r.Status != "200 ok" {
		panic(err)
	}
	//fmt.Println(reflect.TypeOf(r.Status))
	//fmt.Println("%s\n", html)


	End:
		// read in the entire body of the url
		html, err := io.Copy(os.Stdout, r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(html)
}

package main

import (
	"DataScience/DataStructs"
)

func main() {

	//file_ := "/Users/williamrobertmurphy/Desktop/housing_data.rtf"

	//data_array := DataStructs.Open_file(file_, ",")

	//fmt.Println(data_array)

	url_ := "https://assets.datacamp.com/production/course_3423/datasets/crime_sampler.csv"
	//url_ := "http://tour.golang.org/welcome/1"

	//data_ := DataStructs.Open_web_file(url_)
	//fmt.Println(data_)

	//DataStructs.Strings()
	//fmt.Println(reflect.TypeOf(DataStructs.Strings()))
	DataStructs.Open_web_file(url_)
}


package DataStructs

import (
	"fmt"
	"strings"

)

var p = fmt.Println


func Strings() ([]string) {

	p("Contains", strings.Contains("test", "es"))
	p(strings.Split("Name,Data,Age,Job", ","))

	str_arr := strings.Split("Name,Data,Age,Job", ",")
	return str_arr
}



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
package lib

import (
"fmt"
"io/ioutil"
"os"
)

func open_file(file_path string)(string) {
	/*load data into current application*/

	f, err := os.Open(file_path)

	if err != nil {	// if an error is raised during opening
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// output the file content
	fmt.Printf("### File Content ###\n%s\n", string(c))
	return string(c)
}




