package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	Version string = "11.6.1"
	Url     string = "https://ddragon.leagueoflegends.com/"
)

type ResponseChampions struct {
	Data map[string]interface{} `json:"data"`
}

func GetListChampions() {
	getListChampionsUrl := Url + "cdn/" + Version + "/data/en_US/champion.json"

	response, err := http.Get(getListChampionsUrl)
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

	for k := range dataChampions.Data {
		fmt.Println(k)
	}
}

func GetVersion() string {
	versionUrl := Url + "realms/na.json"

	response, err := http.Get(versionUrl)
	if err != nil {
		log.Fatalf("%+v\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	var dataVersion map[string]interface{}
	json.Unmarshal(data, &dataVersion)

	return dataVersion["v"].(string)
}
