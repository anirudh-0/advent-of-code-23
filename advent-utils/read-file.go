package adventutils

import (
	"io"
	"os"
)

func FetchInput(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
