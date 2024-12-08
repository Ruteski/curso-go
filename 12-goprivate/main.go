package main

import (
	"fmt"

	"github.com/ruteski/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
