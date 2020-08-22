package generation

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path"
	"strconv"

	"github.com/larshelmer/gpxtool/internal/resource"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Trail struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var trails = []Trail{
	{
		"SL1",
		"Kust till kust",
	},
}

type Coordinate struct {
	Section   string  `json:"section"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func readGPXFile2(name string, folder string) ([]TrackSegment, error) {
	file := path.Join(folder, name)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return []TrackSegment{}, err
	}

	var gpx GPX

	if err = xml.Unmarshal(b, &gpx); err != nil {
		var e *xml.SyntaxError
		if errors.As(err, &e) { // assume it is a unmarked ISO-8859-1 file
			r := transform.NewReader(bytes.NewReader(b), charmap.ISO8859_1.NewDecoder())
			d := xml.NewDecoder(r)
			if err := d.Decode(&gpx); err != nil {
				log.Println(file)
				return []TrackSegment{}, err
			}
		} else {
			log.Println(file)
			return []TrackSegment{}, err
		}
	}

	return gpx.Track.TrackSegments, nil
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func distanceBetweenCoordinates(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	earthRadiusKm := 6371.0

	dLat := degreesToRadians(lat2 - lat1)
	dLon := degreesToRadians(lon2 - lon1)

	lat1 = degreesToRadians(lat1)
	lat2 = degreesToRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c
}

func Json(folder string) error {
	coordinates := make(map[string][]Coordinate)
	for _, section := range resource.Sections {
		segments, err := readGPXFile2(section.File, folder)
		if err != nil {
			return err
		}
		var plat, plon float64
		for _, segment := range segments {
			for _, point := range segment.TrackPoints {
				lat, err := strconv.ParseFloat(point.Latitude, 64)
				if err != nil {
					return err
				}
				lon, err := strconv.ParseFloat(point.Longitude, 64)
				if err != nil {
					return err
				}
				if plat != 0.0 && plon != 0.0 {
					section.GPSLength += distanceBetweenCoordinates(plat, plon, lat, lon)
				}
				plat = lat
				plon = lon
				coordinates[section.Trail] = append(coordinates[section.Trail], Coordinate{
					section.Code,
					lat,
					lon,
				})
			}
		}
	}

	trailsJson, err := json.MarshalIndent(trails, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(folder, "trails.json"), trailsJson, 0o600)
	if err != nil {
		return err
	}

	sectionsJson, err := json.MarshalIndent(resource.Sections, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(folder, "sections.json"), sectionsJson, 0o600)
	if err != nil {
		return err
	}

	for k, v := range coordinates {
		coordinatesJson, err := json.MarshalIndent(v, "", "	")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path.Join(folder, fmt.Sprintf("%v_coordinates.json", k)), coordinatesJson, 0o600)
		if err != nil {
			return err
		}
	}

	return nil
}
