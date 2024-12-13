package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	return "zero", errors.New("toatp key was not found")
}

func main() {
	content, err := ioutil.ReadFile("/Users/dg/t/golearn/toatp/keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var toatps []toatp
	json.Unmarshal([]byte(content), &toatps)
	key, err := getKeyByName(os.Args[1], &toatps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gotp.NewDefaultTOTP(key).Now())

}
