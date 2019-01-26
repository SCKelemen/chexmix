package scanner

import (
	"fmt"
	"os"
	"path/filepath"
)

type Scanner struct {
	Path  string
	files []string
}

func New(path string) Scanner {
	return Scanner{Path: path}
}

func (s Scanner) Scan() {
	err := filepath.Walk(s.Path, func(path string, info os.FileInfo, err error) error {
		s.files = append(s.files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range s.files {
		fmt.Println(file)
	}
}
