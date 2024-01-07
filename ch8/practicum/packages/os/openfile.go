package main

import (
	"fmt"
	"os"
)

func main() {
	name := "dias.txt"
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		file, err = os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		fmt.Println("File open")
	} else {
		fmt.Println("File created")
	}
}