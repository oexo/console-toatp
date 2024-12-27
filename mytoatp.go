package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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
	return "zero", errors.New("toatp key " + name + " was not found")
}

func getAllToatps(sl *[]toatp) {
	for _, v := range *sl {
		otp, expiredTimestamp := gotp.NewDefaultTOTP(string(v.Key)).NowWithExpiration()
		fmt.Println("ET(sec):", expiredTimestamp-time.Now().Unix(), "- OTP:", otp, "- OTP Name:", v.Name)
	}
}

func addNewTotp(name string, key string, sl *[]toatp) {
	newTotp := toatp{Name: name, Key: key}
	*sl = append(*sl, newTotp)
	jsonData, err := json.Marshal(*sl)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("/Users/dg/t/golearn/toatp/keys.json", jsonData, 0644)
}

func main() {
	content, err := ioutil.ReadFile("/Users/dg/t/golearn/toatp/keys.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var toatps []toatp
	json.Unmarshal([]byte(content), &toatps)

	switch len(os.Args) {
	case 2:
		key, err := getKeyByName(os.Args[1], &toatps)
		if err != nil {
			log.Fatal(err)
		}

		otp, expiredTimestamp := gotp.NewDefaultTOTP(key).NowWithExpiration()
		currentTimestamp := time.Now().Unix()
		l := log.New(os.Stderr, "", 0)
		l.Println("ET(sec):", expiredTimestamp-currentTimestamp)
		fmt.Printf(otp)
	case 4:
		addNewTotp(os.Args[2], os.Args[3], &toatps)
	default:
		getAllToatps(&toatps)
	}

}
