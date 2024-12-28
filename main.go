package main

import "fmt"

type totp struct {
	Name string
	Key  string
}

func main() {

	var totps []totp

	// создаем слайс структур
	totps = append(totps, totp{Name: "hello", Key: "world"})
	totps = append(totps, totp{Name: "nohello", Key: "noworld"})
	fmt.Println(totps)

	newTotps := &totps

	// берем индекс от родительского слайса
	fmt.Println((*newTotps)[0])

	/*
	   	# command-line-arguments
	     ./main.go:21:23: invalid operation: cannot index newTotps (variable of type *[]totp)
	*/

}
