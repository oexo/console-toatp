package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type toatp struct {
	Name string
	Key  string
}

func getKeyByName(name string, sl []toatp) {
	for k, _ := range sl {
		if name == sl[k].Name {
			fmt.Println(sl[k].Key)
		}
	}
}

func testPointers(sl *[]toatp) {
	//fmt.Println(*sl)
	fmt.Println(*sl)
	for _, v := range *sl {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("vim-go")
	content, err := ioutil.ReadFile("./keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println(string(content))
	var toatps []toatp
	json.Unmarshal([]byte(content), &toatps)
	fmt.Printf("Birds : %+v \n", toatps)
	fmt.Println(toatps[0].Name)
	getKeyByName("test 0", toatps)

	testPointers(&toatps)

}
