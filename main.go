package main

import (
	"flag"
	"game_platform"
)

func main() {
	port := flag.Int("port", 8080, "TCP or UDP")
	game_platform.Init(*port)
}
