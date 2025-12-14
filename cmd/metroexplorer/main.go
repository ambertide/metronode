package main

import (
	"ambertide/metromap/internal/cmd/metroexplorer"
	"ambertide/metromap/pkg/metronode"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags)
	line := metronode.ExtractMetronodes("M1")
	data := metroexplorer.FetchMetroSystemData("IZBAN")
	println(data)
	print(line.GetLineData())
}
