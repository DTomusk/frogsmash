package main

import (
	"fmt"
	"frogsmash/internal/game"
)

func main() {
	fmt.Println("Hello, World!")
	items := game.GenerateItems()
	game.Run(items)
}
