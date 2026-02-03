package main

import (
	"fmt"

	"github.com/CybrRonin/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	cfg.SetUser("alex")

	cfg, err = config.Read()
	fmt.Println(cfg)
}
