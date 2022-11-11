package main

import (
	"fmt"
	"goTour/greetings"
	"log"
)

func main () {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	 // A slice of names.
	 names := []string{"Bintu", "Rugy", "Danke"}

	 // Request greeting messages for the names.
	 messages, err := greetings.Hellos(names)


	if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}


