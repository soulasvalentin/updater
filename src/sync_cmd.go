package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func sync(c config) {
	// download & parse remote manifest
	fmt.Println("Downloading remote manifest..")
	manifestUrl := urlJoin(c.RemoteUrl, ManifestFileName)
	rmBytes, rmErr := downloadFileToMemory(manifestUrl)
	if rmErr != nil {
		log.Fatal("Failed to download remote manifest ", rmErr)
	}
	var rm manifest
	if err := json.Unmarshal(rmBytes, &rm); err != nil {
		log.Fatal("Failed to parse remote manifest ", err)
	}

	// calculate diff manifest
	fmt.Println("Calculating diff manifest..")
	dm := calculateDiffManifest(rm)

	if len(dm.Files) > 0 {
		// download diff manifest
		fmt.Println("Downloading files..")
		downloadManifestFiles(dm, c.RemoteUrl)
	}
}
