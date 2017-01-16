package main

import (
	"fmt"

	"github.com/waterdudu/cycler/hub"
)

func main() {

	hub := hub.NewHub()

	fmt.Printf("Cycler hub : %p", hub)
}
