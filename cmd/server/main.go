package main

import (
	"github.com/Scrowszinho/full-cycle-go/tree/master/apis/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	println(config.DBDriver)
}
