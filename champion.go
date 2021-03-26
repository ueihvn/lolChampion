package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetChampion(champion string) {
	championUrl := Url + "cdn/" + Version + "/data/en_US/champion/" + champion + ".json"

	response, err := http.Get(championUrl)
	if err != nil {
		log.Fatalf("%+v\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	var responeObject Response
	json.Unmarshal(data, &responeObject)

	responeObject.Data.Aatrox.ChampionPrint()
}
