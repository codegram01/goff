package main

import (
	"fmt"
	"log"
	"testing"
)

func TestIsDir(t *testing.T) {
	isDir, err := IsDir("testdata/new")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isDir)
}