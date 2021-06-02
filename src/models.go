package main

type config struct {
	RemoteUrl string `json:"remoteUrl"`
}

type manifestFile struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
	Size int64  `json:"size"`
}

type manifest struct {
	Files []manifestFile `json:"files"`
}
