// Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента
// по одному аргументу в строке.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	sep = "\n"
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + " " + os.Args[i] + sep
	}
	fmt.Println(s)
}
