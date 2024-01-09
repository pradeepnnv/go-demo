package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type MyConfig struct {
	Port     string `env:"PORT, default=5555"`
	Username string `env:"USERNAME, required"`
}

func main() {
	ctx := context.Background()

	var c MyConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
