package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
)

func saveJsonToFile(obj interface{}, path string) {
	file, _ := json.MarshalIndent(obj, "", "  ")
	err := ioutil.WriteFile(path, file, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func urlJoin(a string, b string) string {
	b = strings.ReplaceAll(b, "\\", "/")
	u, _ := url.Parse(a)
	u.Path = path.Join(u.Path, b)
	return u.String()
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
