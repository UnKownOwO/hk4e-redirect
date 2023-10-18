package main

import (
	"context"
	"flag"
	"fmt"
	"hk4e-redirect/app"
	"os"
)

var (
	config = flag.String("config", "application.toml", "config file")
)

func main() {
	flag.Parse()
	err := app.Run(context.TODO(), *config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
