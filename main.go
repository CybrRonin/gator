package main

import (
	"fmt"

	"github.com/CybrRonin/Gator/internal/config"
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
