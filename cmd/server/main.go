package main

import (
	"github.com/Scrowszinho/api-go/tree/master/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	println(config.DBDriver)
}
