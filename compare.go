package main

import "fmt"

var (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Reset   = "\033[0m"
	Magenta = "\033[35m"
)

type Compare struct {
	RemovedFiles []string
	CreatedFiles []string
	Changed      [][]byte
}

// compare folder
func DoCompare(oldRootPath, newRootPath string) (*Compare, error) {
	compare := &Compare{}

	oldTree, err := GetTree(oldRootPath)
	if err != nil {
		return nil, err
	}

	newTree, err := GetTree(newRootPath)
	if err != nil {
		return nil, err
	}

	var sameTree = make(Tree)
	for path := range newTree {
		var _, isExit = oldTree[path]
		if isExit {
			sameTree[path] = ""

			delete(newTree, path)
			delete(oldTree, path)
		}
	}

	for path := range oldTree {
		compare.RemovedFiles = append(compare.RemovedFiles, path)
	}

	for path := range newTree {
		compare.CreatedFiles = append(compare.CreatedFiles, path)
	}

	for path := range sameTree {
		pathOldFile := oldRootPath + path
		pathNewFile := newRootPath + path

		diff, err := DoCompareFile(pathOldFile, pathNewFile)
		if err != nil {
			return nil, err
		}
		if diff != nil {
			compare.Changed = append(compare.Changed, diff)
		}
	}

	return compare, nil
}

// compare file
func DoCompareFile(pathOldFile, pathNewFile string) ([]byte, error) {
	oldFile, err := GetFileInfo(pathOldFile)
	if err != nil {
		return nil, err
	}
	if oldFile.IsDir {
		return nil, nil
	}

	newFile, err := GetFileInfo(pathNewFile)
	if err != nil {
		return nil, err
	}
	if newFile.IsDir {
		return nil, nil
	}

	diff := Diff(pathOldFile, oldFile.Data, pathNewFile, newFile.Data)
	return diff, nil
}

// log compare
func LogCompare(compare *Compare) {
	if compare.RemovedFiles != nil {
		fmt.Println(Red + "## Removed:" + Reset)
		for _, path := range compare.RemovedFiles {
			fmt.Printf("+ %s\n", path)
		}
	}

	if compare.CreatedFiles != nil {
		fmt.Println()
		fmt.Println(Green + "## Created:" + Reset)
		for _, path := range compare.CreatedFiles {
			fmt.Printf("+ %s\n", path)
		}
	}

	if compare.Changed != nil {
		fmt.Println()
		fmt.Println(Yellow + "## Changed:" + Reset)
		for _, change := range compare.Changed {
			fmt.Printf("%s", change)
		}
	}
}
