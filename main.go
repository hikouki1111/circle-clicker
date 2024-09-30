package main

import (
	"circle-clicker/game"
)

func main() {
	c := make(chan struct{})
	game.Start()
	<-c
}
