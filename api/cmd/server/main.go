package main

import (
	"fmt"
	"frogsmash/internal/app"
	"frogsmash/internal/delivery/http"
)

func main() {
	fmt.Println("Hello, World!")
	items := app.GenerateItems()
	app.Run(items)

	r := http.SetupRoutes()
	r.Run(":8080")
}
