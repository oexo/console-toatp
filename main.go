package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xlzd/gotp"
)

type toatp struct {
	Name string
	Key  string
}

func getKeyByName(name string, sl *[]toatp) (string, error) {
	for _, v := range *sl {
		if name == v.Name {
			return v.Key, nil
		}
	}
	return "", errors.New("TOATP Not find")
}

func main() {
	fmt.Println("vim-go")
	content, err := ioutil.ReadFile("./keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// fmt.Println(string(content))
	var toatps []toatp
	json.Unmarshal([]byte(content), &toatps)
	//	fmt.Printf("Birds : %+v \n", toatps)
	//	fmt.Println(toatps[0].Name)
	key, err := getKeyByName("test 0", &toatps)
	if err != nil {
		fmt.Println(err)
	}

	// testPointers(&toatps)
	fmt.Println("Current OTP is", gotp.NewDefaultTOTP(key).Now())

}
