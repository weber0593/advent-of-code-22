package file_reader

import (
	"bufio"
	"os"
)

func ReadFile(Path string) ([]string, error) {
	var data []string
	f, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
