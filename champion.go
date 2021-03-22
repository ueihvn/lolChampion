package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Data    Data   `json:"data"`
	Version string `json:"version"`
}

type Data struct {
	Aatrox Aatrox `json:"Aatrox"`
}

type Aatrox struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Stats  BasicStat `json:"stats"`
	Spells []Spell   `json:"spells"`
}

type BasicStat struct {
	HP           int `json:"hp"`
	MP           int `json:"mp"`
	Armor        int `json:"armor"`
	HPregen      int `json:"hpregen"`
	MPregen      int `json:"mpregen"`
	Attackdamage int `json:"attackdamage"`
	Attackspeed  int `json:"attackspeed"`
}

type Spell struct {
	ID   string
	Name string
}

func EncodingChampionJSON() {
	response, err := http.Get("https://ddragon.leagueoflegends.com/cdn/11.6.1/data/vn_VN/champion/Aatrox.json")
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

	fmt.Println(responeObject.Data.Aatrox)
}
