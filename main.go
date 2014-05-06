package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/garslo/app-bot/config"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "app-bot.conf", "location of the config file")
}

func main() {
	flag.Parse()
	config, err := config.LoadConfig(configFile)
	DieIf(err)
	fmt.Println(config)
}

func DieIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problems: %v\n", err)
		os.Exit(1)
	}
}
