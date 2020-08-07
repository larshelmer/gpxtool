package main

import (
	"log"

	"github.com/larshelmer/gpxtool/internal/generation"

	"github.com/moorara/konfig"
)

func main() {
	Config := struct {
		File string
	}{}

	err := konfig.Pick(&Config)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	err = generation.Swift(Config.File)
	if err != nil {
		log.Fatalf("generation error: %v", err)
	}
}
