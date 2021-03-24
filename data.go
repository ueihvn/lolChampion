package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseChampions struct {
	Data map[string]interface{} `json:"data"`
}

func GetListChampion() {
	response, err := http.Get("https://ddragon.leagueoflegends.com/cdn/11.6.1/data/en_US/champion.json")
	if err != nil {
		log.Fatalf("%+v\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	var dataChampions ResponseChampions
	json.Unmarshal(data, &dataChampions)

	for k, _ := range dataChampions.Data {
		fmt.Println(k)
	}

}
