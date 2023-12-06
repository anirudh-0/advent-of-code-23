package trebuchet

import (
	"io"
	"os"
)

func fetchInput() (io.Reader, error) {
	file, err := os.Open("./1-trebuchet/lines.txt")
	if err != nil {
		return nil, err
	}
	return file, nil
}
