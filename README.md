# gpx-nominatim

Extract the address of the first trackpoint of a GPX using Nominatim [Reverse Geocoding] (http://wiki.openstreetmap.org/wiki/Nominatim#Reverse_Geocoding)

Usage:

    ~ go get github.com/ptrv/go-gpx/...
    ~ go build gpx-nominatim.go

    ~ ./gpx-nominatim {file}.gpx


Example:

    ~ gpx-nominatim ~/Documents/gpx-tracks/capanna-gnifetti.gpx
    Rifugio Città di Mantova, sentiero per Rifugio Gnifetti, Gressoney-La-Trinité, VDA, Italia