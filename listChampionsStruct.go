package main

import (
	"log"
	"os"
)

type Data struct {
	Aatrox Champion `json:"Aatrox"`
}

func FileExists() bool {
	if _, err := os.Stat("listChampions.txt"); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.Fatalf("%/+v", err)
	}
	return false
}
