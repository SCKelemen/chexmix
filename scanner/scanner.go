package scanner

/*
import (
	"os"
	"path/filepath"
	//"filepath"
) */

type Scanner struct {
	Path  string
	files []string
}

/*
func New(path string) (Scanner, error) {

	abs, err := filepath.Abs(path)
	if err != nil {
		return Scanner{}, err
	}
	fi, err := os.Stat(abs)
	if err != nil {
		return Scanner{}, err
	}

	if fi.IsDir() {
		for _, item := range fi.
	}
	return Scanner{Path: path}, nil
}

func (s Scanner) Scan() {

}
*/
