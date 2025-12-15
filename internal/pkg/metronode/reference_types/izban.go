package metronode

import (
	metronode "ambertide/metromap/internal/pkg/metronode/common"
	"cmp"
	"encoding/json"
	"slices"
)

type IzbanStation struct {
	Boylam         string
	IstasyonID     int
	Enlem          string
	IstasyonSirasi int
	IstasyonAdi    string
}

// Extract the station line for Izban after sorting
func ExtractIzban(description string) *metronode.Line {
	var unmarshalled []IzbanStation
	json.Unmarshal([]byte(description), &unmarshalled)
	slices.SortFunc(unmarshalled, func(a, b IzbanStation) int {
		return cmp.Compare(a.IstasyonSirasi, b.IstasyonSirasi)
	})
	line := new(metronode.Line)
	line.Name = "IZBAN"
	for _, element := range unmarshalled {
		station := new(metronode.Station)
		station.Name = element.IstasyonAdi
		if line.Start == nil {
			// If no starting station is set yet, set it to this
			line.Start = station
		}
		lastStation := line.Stop
		// First of all we need to have the arriving edge
		if lastStation != nil {
			// Which is the departing edge of the last station if one exists
			station.Arriving = lastStation.Departing
			station.Arriving.B = station
		}
		// Handle the departing
		station.Departing = new(metronode.Edge)
		station.Departing.A = station
		line.Stop = station
	}
	return line
}
