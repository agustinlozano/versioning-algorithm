// Given an array of strings and a plain text file
// make a function to check the matching of the strings in the file
// and return the number of matches

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"reflect"
)

func main() {
	// Define an array of strings with five elements
	guids := []string{
		"0dLSQoDLzFMBd4W2uDyJKF",
		"1KBlFBdjz0Tf_GT2VxZ7_I",
		"36zqNfIRfCvQpZuJcz2xU7",
		"1KBlFBdjz0Tf_GUCtxZ7ti",
		"0wQknZQiXEOBcv3sHGzx_O"}



	// Open the file
	myFile, err := os.Open("./IFCs/myIFC.ifc")
	if err != nil {
		fmt.Println(err)
	}
	defer myFile.Close()

	checkMatches(myFile, guids)

	fmt.Println(reflect.TypeOf(myFile))
}

func checkMatches(ifcModel &os.file, guids []string) int {
	fmt.Println(guids)

	// Read the file
	scanner := bufio.NewScanner(ifcModel)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Close the file
	ifcModel.Close()

	// Check the matches
	matches := 0
	for _, line := range lines {
		if strings.Contains(line, "Hello") {
			matches++
		}
	}

	// Print the number of matches
	fmt.Println(matches)

	return 0
}