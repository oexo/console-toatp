package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type toatp struct {
	name string
	key  string
}

func main() {
	fmt.Println("vim-go")
	content, err := ioutil.ReadFile("./keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println(string(content))
}
