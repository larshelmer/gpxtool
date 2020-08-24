package generation

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Track   Track    `xml:"trk"`
	Route   Route    `xml:"rte"`
}

type Track struct {
	XMLName       xml.Name       `xml:"trk"`
	TrackSegments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	XMLName     xml.Name     `xml:"trkseg"`
	TrackPoints []TrackPoint `xml:"trkpt"`
	Name        string
}

type TrackPoint struct {
	XMLName   xml.Name `xml:"trkpt"`
	Latitude  string   `xml:"lat,attr"`
	Longitude string   `xml:"lon,attr"`
}

type Route struct {
	XMLName     xml.Name     `xml:"rte"`
	RoutePoints []RoutePoint `xml:"rtept"`
}

type RoutePoint struct {
	XMLName   xml.Name `xml:"rtept"`
	Latitude  string   `xml:"lat,attr"`
	Longitude string   `xml:"lon,attr"`
}

func readGPXFile(trail string, folder string, section int) (TrackSegment, error) {
	file := path.Join(folder, fmt.Sprintf("%v_%v.gpx", trail, section))
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return TrackSegment{}, err
	}

	var gpx GPX

	err = xml.Unmarshal(b, &gpx)
	if err != nil {
		return TrackSegment{}, err
	}

	if len(gpx.Track.TrackSegments) != 1 {
		log.Println("unexpected segment count", len(gpx.Track.TrackSegments), trail, section)
	}

	gpx.Track.TrackSegments[0].Name = fmt.Sprintf("%v-%v", trail, section)

	return gpx.Track.TrackSegments[0], nil
}

func Swift(trail string, folder string, count int) error {
	var segments []TrackSegment
	for i := 1; i <= count; i++ {
		s, err := readGPXFile(trail, folder, i)
		if err != nil {
			return err
		}
		segments = append(segments, s)
	}

	file := path.Join(folder, fmt.Sprintf("%v.swift", trail))

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	templ, err := template.New("swift_trails.tmpl").Funcs(template.FuncMap{
		"name": func() string { return trail },
	}).ParseFiles("templates/swift_trails.tmpl")
	if err != nil {
		return err
	}

	return templ.Execute(f, segments)

	/*	return template.Must(template.ParseFiles("templates/swift_trails.tmpl")).Funcs(template.FuncMap{
		"name": func() string { return trail },
	}).Execute(f, segments) */
}
