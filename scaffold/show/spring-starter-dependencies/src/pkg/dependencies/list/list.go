package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	listDependenciesURL = "https://start.spring.io/metadata/client"
)

type Inputs struct {
}

type Result struct {
	Dependencies Dependencies
}

type Dependencies struct {
	Type   string
	Values []TypeDependency
}

type TypeDependency struct {
	Name   string
	Values []Dependency
}

type Dependency struct {
	Id          string
	Name        string
	Description string
}

func (in Inputs) Run() {
	log.Println("list dependencies")

	resp, err := http.Get(listDependenciesURL)
	if err != nil {
		log.Fatal("Failed get list dependencies", err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Failed get list dependencies! Response Status Code: ", resp.Status)

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data Result
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Failed unmarshall", err)
	}

	for _, t := range data.Dependencies.Values {
		fmt.Printf("Type name: %s \n\n", t.Name)
		for _, d := range t.Values {
			fmt.Printf("\tId: %v\n", d.Id)
			fmt.Printf("\tName: %v\n", d.Name)
			fmt.Printf("\tDescription: %v\n\n", d.Description)
		}
	}
}
