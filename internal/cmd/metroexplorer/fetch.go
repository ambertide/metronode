package metroexplorer

import (
	metronode "ambertide/metronode/internal/pkg/metronode"
	io "io"
	"log"
	http "net/http"
)

// Get the data from sources using net/http
// we can't use net/http in  wasm so we will assume
// the JSON gets to the package somehow.

// But for our flows we should use net/http

func FetchMetroSystemData(systemName string) string {
	sourceURL := metronode.Sources[systemName]
	if sourceURL == "" {
		log.Println("Could not find source for: ", systemName)
		return ""
	}
	resp, err := http.Get(sourceURL)
	if err != nil {
		log.Println("Error occurred in getting source: ", err.Error())
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		return string(body)
	}
	log.Println("Error occurred reading the response body: ", err.Error())
	return ""
}
