package main

import (
	"fmt"
	"os"
)

func init() {
	user := os.Getenv("USER")
	if user == "" {
		panic("$USER is not set!!!")
	}
	fmt.Println("user:", user)
}

func init() {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostname:", h)
}

func main() {

}
