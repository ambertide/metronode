package common

import "structs"

type Edge struct {
	A      *Station
	B      *Station
	Active bool
	layout structs.HostLayout
}

type Station struct {
	Arriving  *Edge
	Departing *Edge
	Active    bool
	Name      string
	layout    structs.HostLayout
}

type Line struct {
	Start  *Station
	Stop   *Station
	Name   string
	layout structs.HostLayout
}

func (l *Line) GetLineData() string {
	repr := ""
	currentStation := l.Start
	for currentStation != nil {
		repr += currentStation.Name
		if currentStation.Departing == nil {
			return repr
		}
		currentStation = currentStation.Departing.B
	}
	return repr
}
