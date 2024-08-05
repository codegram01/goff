package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	// user enter command : goff [old path] [new path]
	oldRootPath := flag.Arg(0)
	if oldRootPath == "" {
		fmt.Print(helpMess)
		return
	}
	isOldRootPathDir, err := IsDir(oldRootPath)
	if err != nil {
		log.Fatal(err)
	}

	newRootPath := flag.Arg(1)
	if newRootPath == "" {
		fmt.Print(helpMess)
		return
	}
	isNewRootPathDir, err := IsDir(newRootPath)
	if err != nil {
		log.Fatal(err)
	}

	// both compare is folder
	if isOldRootPathDir && isNewRootPathDir {
		compare, err := DoCompare(oldRootPath, newRootPath)
		if err != nil {
			log.Fatal(err)
		}
		LogCompare(compare)
	}

	// both compare is files
	if !isOldRootPathDir && !isNewRootPathDir {
		changed, err := DoCompareFile(oldRootPath, newRootPath)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(Yellow + "## Changed:" + Reset)
		fmt.Printf("%s", changed)
	}

	if isOldRootPathDir && !isNewRootPathDir || !isOldRootPathDir && isNewRootPathDir {
		log.Fatal("Only can compare file vs file , folder vs folder")
	}

}
