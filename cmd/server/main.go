package main

import (
	"github.com/Scrowszinho/api-go/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	println(config.DBDriver)
}
