package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func getRandomNumber() int {
	return rand.Intn(89139999999-89130000000) + 89130000000
}

func main() {
	rand.Seed(time.Now().UnixNano())

	addressBook := make(map[string]int)

	data, err := ioutil.ReadFile("data.json")

	if err == nil {
		err = json.Unmarshal(data, &addressBook)

		if err != nil {
			fmt.Println("Can not parse data.json")
		}
	} else {
		fmt.Println("Can not read the file data.json")
	}

	fmt.Println("Current data:", addressBook)

	addressBook["Alex"] = getRandomNumber()
	addressBook["Bob"] = getRandomNumber()

	fmt.Println("New data:", addressBook)

	data, err = json.Marshal(addressBook)

	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("data.json", data, 0644)

	if err != nil {
		fmt.Println("Can not write a file data.json", err)
	}
}
