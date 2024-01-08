package main

import (
	"fmt"
	"os"
	"strings"

	utils "moviedb/internal/utils"
)

const (
	modifyList  = "modifyList"
	createMovie = "createMovie"
)

var testsToRun []func()

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("Goodbye!")

	tests := os.Getenv("TESTS")
	fmt.Println("gathering tests...")
	for _, test := range strings.Split(tests, ",") {
		switch test {
		case modifyList:
			testsToRun = append(testsToRun, utils.TestModifyList)
		case createMovie:
			testsToRun = append(testsToRun, utils.TestCreateMovie)
		}
		fmt.Println("complete")
	}

	if testsToRun != nil {
		fmt.Println("running tests")
		for _, test := range testsToRun {
			test()
		}
		fmt.Println("complete")
	}

}
