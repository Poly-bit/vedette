package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
)

type networkInformation struct {
	Networks []string
	Ports    []string
}

func main() {
	fmt.Println("Vedette - A certificate scanner")

	configFlag := flag.String("config", "example.toml", "Path to a toml file")
	flag.Parse()

	var config networkInformation
	if _, err := toml.DecodeFile(*configFlag, &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ranges to scan %s on ports %s\n", config.Networks, config.Ports)
}
