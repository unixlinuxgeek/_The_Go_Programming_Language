package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Вывод (Вывод главной go-подпрограмме)
	for {
		fmt.Println(<-squares)
	}
}
