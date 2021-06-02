package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Builds manifest from local directory (path, hash and size)
func buildManifest(dir string, fileNamePrefixExclude string) manifest {
	m := manifest{
		Files: []manifestFile{},
	}
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && !strings.HasPrefix(info.Name(), fileNamePrefixExclude) {
				m.Files = append(m.Files, manifestFile{
					Path: path,
					Size: info.Size(),
					Hash: calculateHash(path),
				})
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return m
}

const ()

// Determines which files to download by comparing remote and local files.
// Prints info per file.
func calculateDiffManifest(rm manifest) manifest {
	var dm manifest

	for _, rf := range rm.Files {
		s := "unchanged"

		// missing?
		if _, err := os.Stat(rf.Path); err != nil {
			s = "missing"
		} else {
			// modified?
			if calculateHash("./"+rf.Path) != rf.Hash {
				s = "outdated"
			}
		}

		fmt.Printf(" - [%v] %v (%v bytes)\n", s, rf.Path, rf.Size)

		if s != "unchanged" {
			dm.Files = append(dm.Files, rf)
		}
	}

	return dm
}

// Downloads each file in the manifest.
// Breaks if error
func downloadManifestFiles(m manifest, remoteUrl string) {
	for _, mf := range m.Files {
		fmt.Printf(" - downloading: %v (%v bytes)\n", mf.Path, mf.Size)
		u := urlJoin(remoteUrl, mf.Path)
		if err := downloadFile(u, mf.Path); err != nil {
			fmt.Printf("[error] Failed to download file '%v'. %v\n", u, err.Error())
			break
		}
	}
}
