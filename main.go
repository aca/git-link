package main

import (
	"log"
)

type Link struct {
	Source      string
	Destination string
	Size        int64
	XXH64       string
}

const version = "v0"

func main() {
	log.SetPrefix("[git-link] ")
	err := cmdMain()
	if err != nil {
		log.Fatal(err)
	}
}
