package main

import (
	"os"
)

func isExist(dir *string) bool {
	path := *dir
	_, err := os.Stat(path)

	return os.IsNotExist(err)
}
