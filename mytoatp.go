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

func getAllToatps(sl *[]toatp) {
	for _, v := range *sl {
		otp := gotp.NewDefaultTOTP(string(v.Key))
		fmt.Println(string(v.Name) + " - one-time password of timestamp 0 is: " + string(otp.At(0)) + " - Current one-time password is: " + string(gotp.NewDefaultTOTP(v.Key).Now()))
	}
}

func main() {
	content, err := ioutil.ReadFile("/Users/dg/t/golearn/toatp/keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var toatps []toatp
	json.Unmarshal([]byte(content), &toatps)

	var ttp string

	if len(os.Args) > 1 {
		ttp = os.Args[1]
	} else {
		ttp = "all"
	}

	if ttp == "all" {
		getAllToatps(&toatps)
	} else {
		key, err := getKeyByName(os.Args[1], &toatps)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(gotp.NewDefaultTOTP(key).Now())
	}

}
