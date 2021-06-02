package main

import (
	"crypto/md5"
	b64 "encoding/base64"
	"io"
	"log"
	"os"
)

func calculateHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sum := h.Sum(nil)
	s := b64.StdEncoding.EncodeToString(sum)

	return s
}
