package main

import "apis/configs"

func main() {
	config := configs.NewConfig()
	println(config.GetDBDriver())
}
