package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"moviedb/internal/db"
	"moviedb/internal/server"
	utils "moviedb/internal/utils"
)

const (
	modifyList  = "modifyList"
	createMovie = "createMovie"
)

var testsToRun []func(utils.InitConfig)

func testRunner(cfg utils.InitConfig) {
	fmt.Println("Hello!")
	defer fmt.Println("Test runner out!")

	tests := os.Getenv("TESTS")
	fmt.Println("gathering tests...")
	fmt.Printf("there are %v tests to run\n", len(strings.Split(tests, ","))-1)
	for _, test := range strings.Split(tests, ",") {
		switch test {
		case modifyList:
			testsToRun = append(testsToRun, utils.TestModifyList)
		case createMovie:
			testsToRun = append(testsToRun, utils.TestCreateMovie)
		}
		fmt.Printf("complete: %v\n", testsToRun)
	}

	if len(testsToRun) != 0 {
		fmt.Println("running tests")
		for _, test := range testsToRun {
			test(cfg)
		}
		fmt.Println("complete")
	}

}

func main() {
	cfg := utils.LoadEnv()
	db.Init(cfg.PostgresURI, cfg.Debug)

	if cfg.RunTests {
		testRunner(cfg)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go server.StartServer()

	<-signalChan
	fmt.Println("\nshut down complete")
}
