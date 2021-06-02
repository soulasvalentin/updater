package main

// Build & save local manifest
func build(c config) {
	m := buildManifest(DefaultFilesDir, "updater")
	saveJsonToFile(m, ManifestFileName)
	// todo: output files and totals
}
