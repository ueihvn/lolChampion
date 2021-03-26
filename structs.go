package main

import "fmt"

type Response struct {
	Data    Data   `json:"data"`
	Version string `json:"version"`
}

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
	fmt.Printf("%s - %v\n", c.ID, c.Title)
	c.Stats.BasicStatPrint()

	for _, v := range c.Spells {
		v.SpellPrint()
	}
}
