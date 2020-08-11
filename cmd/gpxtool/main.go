package main

import (
	"log"

	"github.com/larshelmer/gpxtool/internal/download"
	"github.com/larshelmer/gpxtool/internal/generation"
	"github.com/moorara/konfig"
)

func main() {
	Config := struct {
		Generate bool
		Get      bool
		File     string
		Folder   string
	}{
		Folder: ".",
	}

	err := konfig.Pick(&Config)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	if Config.Get {
		err := download.Do("SL1", Config.Folder, 1)
		if err != nil {
			log.Fatalf("get error: %v", err)
		}
	}

	if Config.Generate {
		err = generation.Swift(Config.File)
		if err != nil {
			log.Fatalf("generation error: %v", err)
		}
	}
}
