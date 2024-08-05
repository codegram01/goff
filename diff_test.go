package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDiff(t *testing.T) {
	pathOldFile := "testdata/old/mess.txt"
	oldFile, err := os.ReadFile(pathOldFile)
	if err != nil {
		log.Fatal(err)
	}

	pathNewFile := "testdata/new/mess.txt"
	newFile, err := os.ReadFile(pathNewFile)
	if err != nil {
		log.Fatal(err)
	}

	// i use path for name 
	compare := Diff(pathOldFile, oldFile, pathNewFile, newFile)

	fmt.Printf("%s", compare)
}