package generation

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Track   Track    `xml:"trk"`
}

type Track struct {
	XMLName       xml.Name       `xml:"trk"`
	TrackSegments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	XMLName     xml.Name     `xml:"trkseg"`
	TrackPoints []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	XMLName   xml.Name `xml:"trkpt"`
	Latitude  string   `xml:"lat,attr"`
	Longitude string   `xml:"lon,attr"`
}

func Swift(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var gpx GPX

	err = xml.Unmarshal(b, &gpx)
	if err != nil {
		return err
	}

	d := filepath.Base(file) + ".swift"

	f, err := os.Create(d)
	if err != nil {
		return err
	}

	defer f.Close()

	return template.Must(template.ParseFiles("templates/swift_trails.tmpl")).Execute(f, gpx.Track.TrackSegments[0])
}
