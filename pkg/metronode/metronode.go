package metronode

import "ambertide/metromap/internal/pkg/metronode"

func ExtractMetronodes(name string) *metronode.Line {
	line := new(metronode.Line)
	line.Name = name
	start := new(metronode.Station)
	start.Name = "Hilal"
	stop := new(metronode.Station)
	stop.Name = "Alsancak"
	edge := new(metronode.Edge)
	edge.A = start
	edge.B = stop
	start.Departing = edge
	stop.Arriving = edge
	line.Start = start
	line.Stop = stop
	return line
}
