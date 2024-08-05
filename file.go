package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	IsDir bool
	Data  []byte
}

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

func GetFileInfo(path string) (*FileInfo, error) {
	isDir, err := IsDir(path)
	if err != nil {
		return nil, err
	}
	if isDir {
		return &FileInfo{
			IsDir: true,
			Data:  nil,
		}, nil
	}

	file, err := os.ReadFile(path)
	return &FileInfo{
		IsDir: false,
		Data:  file,
	}, err
}

type Tree map[string]string

// this function must return map
func GetTree(rootPath string) (Tree, error) {
	var tree = make(Tree)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if path == rootPath {
			return nil
		}

		path = strings.TrimPrefix(path, rootPath)

		tree[path] = ""

		return nil
	})

	return tree, err
}
