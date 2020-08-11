package download

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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

func Do(trail string, folder string, count int) error {
	for i := 1; i <= count; i++ {
		section := fmt.Sprintf("%v_%v.gpx", trail, i)
		gpxFile := path.Join(folder, section)
		etagFile := gpxFile + ".etag"

		url := baseURL + section
		if !needsUpdate(gpxFile, etagFile, url, section) {
			log.Println("etag match", section)
			continue
		}

		// nolint:gosec // G107: Potential HTTP request made with variable url
		resp, err := http.Get(url)
		if err != nil {
			log.Println("get error", section, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println("http error", section, resp.StatusCode)
			continue
		}

		for k, v := range resp.Header {
			log.Println(k, v[0])
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("response read error", section, err)
			resp.Body.Close()

			continue
		}

		resp.Body.Close()

		err = ioutil.WriteFile(gpxFile, b, 0o600)
		if err != nil {
			log.Println("gpx write error", section, err)
			continue
		}

		etag := resp.Header.Get(etagHeader)
		if etag != "" {
			err = ioutil.WriteFile(etagFile, []byte(etag), 0o600)
			if err != nil {
				log.Println("etag write error", section, err)
				continue
			}
		} else {
			log.Println("no etag header", section)
		}
	}

	return nil
}
