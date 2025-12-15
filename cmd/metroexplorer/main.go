package main

import (
	"ambertide/metromap/internal/cmd/metroexplorer"
	"ambertide/metromap/pkg/metronode"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags)
	data := metroexplorer.FetchMetroSystemData("IZBAN")
	graph, _ := metronode.ExtractMetronodes("IZBAN", data)
	currentNode := graph.Start
	for currentNode != nil {
		println(currentNode.Name)
		if currentNode.Departing == nil {
			break
		}
		currentNode = currentNode.Departing.B
	}
}
