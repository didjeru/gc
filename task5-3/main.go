package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	Surname    string   `json:"surname"`
	Firstname  string   `json:"firstname"`
	Patronymic string   `json:"patronymic"`
	Address    *address `json:"address,omitempty"`
}

type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	csvFile, _ := os.Open("base1.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []person
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		people = append(people, person{
			Firstname:  line[1],
			Surname:    line[0],
			Patronymic: line[2],
			Address: &address{
				City:  line[3],
				State: line[4],
			},
		})
	}
	peopleJSON, _ := json.Marshal(people)
	fmt.Println(string(peopleJSON))
}
