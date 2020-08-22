package download

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/larshelmer/gpxtool/internal/resource"
)

const (
	baseURL    = "https://cdn.hoodin.com/skaneleden/assets/geomanager/geofiles/"
	etagHeader = "Etag"
)

func needsUpdate(gpxFile string, etagFile string, url string, section string) bool {
	if _, err := os.Stat(gpxFile); err != nil {
		return true
	}

	if _, err := os.Stat(etagFile); err != nil {
		return true
	}

	b, err := ioutil.ReadFile(etagFile)
	if err != nil {
		log.Println("etag read error", section, err)
		return true
	}

	etag := string(b)
	if etag == "" {
		return true
	}

	// nolint:gosec // G107: Potential HTTP request made with variable url
	resp, err := http.Head(url)
	if err != nil {
		log.Println("head error", section, err)
		return true
	}

	resp.Body.Close()

	return resp.Header.Get(etagHeader) != etag
}

func Do(folder string) error {
	for _, section := range resource.Sections {
		gpxFile := path.Join(folder, section.File)
		etagFile := gpxFile + ".etag"

		url := baseURL + section.File
		if !needsUpdate(gpxFile, etagFile, url, section.File) {
			log.Println("etag match", section.File)
			continue
		}

		// nolint:gosec // G107: Potential HTTP request made with variable url
		resp, err := http.Get(url)
		if err != nil {
			log.Println("get error", section.File, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println("http error", section.File, resp.StatusCode)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("response read error", section.File, err)
			resp.Body.Close()

			continue
		}

		resp.Body.Close()

		err = ioutil.WriteFile(gpxFile, b, 0o600)
		if err != nil {
			log.Println("gpx write error", section.File, err)
			continue
		}

		etag := resp.Header.Get(etagHeader)
		if etag != "" {
			err = ioutil.WriteFile(etagFile, []byte(etag), 0o600)
			if err != nil {
				log.Println("etag write error", section.File, err)
				continue
			}
		} else {
			log.Println("no etag header", section.File)
		}
	}

	return nil
}
