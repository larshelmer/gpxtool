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
		Folder   string
	}{
		Folder: "gpx",
	}

	err := konfig.Pick(&Config)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	if Config.Get {
		if err := download.Do(Config.Folder); err != nil {
			log.Println("download error", err)
		}
	}

	if Config.Generate {
		if err = generation.Json(Config.Folder); err != nil {
			log.Println("generation error", err)
		}
	}
}
