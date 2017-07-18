package main

import (
	"fmt"
	"io/ioutil"

	"github.com/barcatfigaro/Vysoify/vysoify"
)

func main() {
	bytes, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
		return
	}
	vysoify.FixVysoFakes(bytes)
	fmt.Println(string(bytes))
	err = ioutil.WriteFile("done.txt", bytes, 0644)
	if err != nil {
		fmt.Printf("could not make file: %v\n", err)
		return
	}
}
