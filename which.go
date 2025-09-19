package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please, provide an argument")
		return
	}
	for _, file := range arguments {
		path := os.Getenv("PATH")
		pathSplit := filepath.SplitList(path)
		for _, directory := range pathSplit {
			fullpath := filepath.Join(directory, file)

			fileInfo, err := os.Stat(fullpath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullpath)
						break
					}
				}
			}
		}
	}
}
