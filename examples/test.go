package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type networkInformation struct {
	Networks	[]string
	Ports		[]string
}

func main() {
	var config networkInformation
	if _, err := toml.DecodeFile("example.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ranges to scan %s on ports %s\n", config.Networks, config.Ports)
}
