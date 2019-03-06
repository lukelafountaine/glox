package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"./scan"
)

func main() {
	if len(os.Args) > 1 {
		program, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		run(string(program))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
			run(scanner.Text())
		}
	}
}

func run(source string) {
	scanner := scan.NewScanner(source)
	scanner.ScanTokens()
}
