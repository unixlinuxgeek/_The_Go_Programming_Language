// Create if file not exist and add text.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	txt_data := "Lorem ipsum..."

	name := "file.txt"
	file, err := os.Open(name)
	if err != nil {
		file, err = os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.WriteFile(name, []byte(txt_data), 0755)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 8)
	cnt, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", cnt, data[:cnt])
}
