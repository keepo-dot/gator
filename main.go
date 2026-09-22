package main

import (
	"fmt"

	"github.com/keepo-dot/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %s\n", err)
		return
	}
	if err := cfg.SetUser("keepo-dot"); err != nil {
		fmt.Printf("error setting user: %s", err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config: %s\n", err)
		return
	}
	fmt.Println(cfg)
}
