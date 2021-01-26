package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ptrv/go-gpx"
)

// Nominatim ...
type Nominatim struct {
	Address struct {
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		County      string `json:"county"`
		Road        string `json:"road"`
		State       string `json:"state"`
		Suburb      string `json:"suburb"`
		Village     string `json:"village"`
	} `json:"address"`
	Boundingbox []string `json:"boundingbox"`
	DisplayName string   `json:"display_name"`
	Lat         string   `json:"lat"`
	Licence     string   `json:"licence"`
	Lon         string   `json:"lon"`
	OsmID       string   `json:"osm_id"`
	OsmType     string   `json:"osm_type"`
	PlaceID     string   `json:"place_id"`
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide a GPX file path!")
		return
	}

	gpxfile := args[0]
	tracks, err := gpx.ParseFile(gpxfile)

	if err != nil {
		fmt.Println("Error opening gpx file: ", err)
		return
	}

	lat := tracks.Tracks[0].Segments[0].Waypoints[0].Lat
	lon := tracks.Tracks[0].Segments[0].Waypoints[0].Lon

	reverseAPI := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%f&lon=%f&zoom=18&addressdetails=1", lat, lon)

	resp, err := http.Get(reverseAPI)
	if err != nil {
		log.Fatal("error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var place Nominatim
	json.Unmarshal(body, &place)
	fmt.Println(place.DisplayName)
}
