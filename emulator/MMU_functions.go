package main

import (
	"bufio"
	"os"
)

func ReadGBFile(filename string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		return nil
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes
}
