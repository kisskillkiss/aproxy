package main

import (
	"aproxy/config"
	"flag"
	"log"
)

var configPath string

func init() {
	log.SetFlags(log.LstdFlags)
	flag.StringVar(&configPath, "c", "config.txt", "filepath of config")
	flag.Parse()
}

func main() {
	c := config.OpenConfig(configPath)
	if c == nil {
		return
	}
	c.Run()
}
