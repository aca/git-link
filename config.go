package main

type Config struct {
	Version string `json:"version"`
	Links   []Link `json:"links"`
}
