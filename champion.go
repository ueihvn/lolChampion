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

type Aatrox Champion

type Champion struct {
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

func (s *Spell) SpellPrint() {
	fmt.Printf("%s: %s\n", s.ID, s.Name)
}

func (bs *BasicStat) BasicStatPrint() {
	fmt.Printf(
		"HP/RegenHP: %12d/%d\nMP/RegenMP: %12d/%d\nArmor: %12d\nAttack damage: %12d\nAttack speed: %12d\n",
		bs.HP, bs.HPregen,
		bs.MP, bs.MPregen,
		bs.Armor,
		bs.Attackdamage,
		bs.Attackspeed)
}

func (c *Champion) ChampionPrint() {
	fmt.Printf("%s - %v", c.ID, c.Title)
	c.Stats.BasicStatPrint()

	for _, v := range c.Spells {
		v.SpellPrint()
	}
}

func EncodingChampionJSON() {
	response, err := http.Get("https://ddragon.leagueoflegends.com/cdn/11.6.1/data/en_US/champion/Aatrox.json")
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

	responeObject.Data.Aatrox.Stats.BasicStatPrint()
}
