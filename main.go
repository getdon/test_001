package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	if f, err := os.Open(".//test.txt"); err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()

		buff := make([]byte, 32)
		if _, err2 := f.Read(buff); err2 != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(buff))
		}
	}
}
