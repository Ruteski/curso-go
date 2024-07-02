package main

import "apis/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(config.DBDriver)
}
