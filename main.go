package main

import (
	"fmt"
	"log"

	"github.com/huangmatty/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.SetUser("matt"); err != nil {
		log.Fatalf("error setting current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.DBUrl)
	fmt.Println(cfg.CurrentUsername)
}
