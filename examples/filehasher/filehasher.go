package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
)

func main() {
	fileHash, err := hashFile("filehasher.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", fileHash)
}

func hashFile(path string) ([]byte, error) {
	// Read File
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Hash it
	h := fnv.New128()

	h.Write(bts)

	fileHash := h.Sum(nil)

	return fileHash, nil
}
