package cya_game

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RollCYAGame() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fileName := flag.String("file", "./cya-game/gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	//Opening the file
	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
