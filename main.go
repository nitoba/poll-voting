package main

import configs "github.com/nitoba/poll-voting/config"

func main() {
	configs.LoadConfig()

	configs.BeforeAll()
}
