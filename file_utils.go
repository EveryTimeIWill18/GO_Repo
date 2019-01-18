package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println(checkExt(".eml"))
}

func checkExt(ext string) []string {

	os.Chdir("C:/Some/Dir")
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

	system_utils
	============
	Package is used to speed up the loading
	of files to the hadoop linux server.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func fileCount(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// an error has occured
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}

	return i, nil
}

// load the files of a certain extension
func getFiles(path string, ext_type string) []string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		// an error has occured
		panic(err)
	}

	var ext_files []string // create a slice to hold the files
	//var i int = 1			 // create an index for the array
	var ext_sum int = 0 // count the number of files of the ext_type

	for _, file := range files {
		if strings.Contains(string(file.Name()), ext_type) {
			ext_files = append(ext_files, file)
			ext_sum++
		}
	}
	fmt.Println("Number of files added to ext_files: ", ext_sum)

	// return the slice
	return ext_files
}

func main() {

	//files, _ := ioutil.ReadDir("Z:/WinRisk/PC_BusinessAnalytics/SA_Claims/SA_2")
	//fmt.Println(len(files))
	var files []string

	root := "Z:/WinRisk/PC_BusinessAnalytics/SA_Claims/SA_2"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {

	}

}
