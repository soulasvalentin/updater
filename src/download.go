package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

// Downloads text file to memory as []byte.
// Returns error if unsuccessful http.
func downloadFileToMemory(url string) ([]byte, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("StatusCode " + string(resp.Status))
	}

	// read content
	b, err := ioutil.ReadAll(resp.Body)
	return b, err
}

// Downloads url to specified path.
// Returns error if unsuccessful http.
// Creates necessary folders.
func downloadFile(url string, localPath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("StatusCode " + string(resp.Status))
	}

	// make sure directory exist
	localPath = strings.ReplaceAll(localPath, "\\", "/")
	dir, _ := path.Split(localPath)
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}

	// create the file
	out, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
