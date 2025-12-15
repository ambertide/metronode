package metronode

import (
	"ambertide/metromap/internal/pkg/metronode"
	common "ambertide/metromap/internal/pkg/metronode/common"
)

func ExtractMetronodes_(name string) *common.Line {
	line := new(common.Line)
	line.Name = name
	start := new(common.Station)
	start.Name = "Hilal"
	stop := new(common.Station)
	stop.Name = "Alsancak"
	edge := new(common.Edge)
	edge.A = start
	edge.B = stop
	start.Departing = edge
	stop.Arriving = edge
	line.Start = start
	line.Stop = stop
	return line
}

// Given the name of system and JSON description of the system
// return a system struct which contains the parsed nodes of the
// subway system.
func ExtractMetronodes(
	systemName string,
	systemDescription string,
) (*common.Line, error) {
	return metronode.ExtractSystemFromString(
		systemName,
		systemDescription,
	)
}
