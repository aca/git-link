package main

import (
	"log"
)

type Link struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Size        int64  `json:"size"`
	XXH64       string `json:"xxh_64"`
}

const version = "v0"

func main() {
	log.SetPrefix("[git-link] ")
	err := cmdMain()
	if err != nil {
		log.Fatal(err)
	}
}
