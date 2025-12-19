package main

import (
	"ambertide/metronode/internal/pkg/metronode/common"
	"ambertide/metronode/pkg/metronode"
	"unsafe"
)

type System struct {
	Line *common.Line
}

var system System

//go:wasmexport extractMetronode
func ExtractMetronode(
	systemName string,
	systemDescription string,
) unsafe.Pointer {
	line, err := metronode.ExtractMetronodes(
		systemName,
		systemDescription,
	)

	if err != nil {
		return nil
	}

	return unsafe.Pointer(line)
}

func main() {

}
