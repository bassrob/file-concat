package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	inputFiles := os.Args[1:]
	outputFile := getFileName()

	reader, err := createReader(inputFiles)
	if err != nil {
		printAndHold(fmt.Sprintf("An error occurred during read: %s", err.Error()))
		return
	}

	writer, err := createWriter(outputFile)
	if err != nil {
		printAndHold(fmt.Sprintf("An error occurred during write: %s", err.Error()))
		return
	}

	err = pipe(reader, writer)
	if err != nil {
		printAndHold(fmt.Sprintf("An error occurred during pipe: %s", err.Error()))
	}
}

func createReader(filePaths []string) (reader io.Reader, err error) {
	readers := []io.Reader{}
	for _, filePath := range filePaths {
		inputFile, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		readers = append(readers, inputFile)
		readers = append(readers, newLineReader())
	}

	return io.MultiReader(readers...), nil
}

func createWriter(filePath string) (writer *bufio.Writer, err error) {
	outputFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return bufio.NewWriter(outputFile), nil
}

func pipe(reader io.Reader, writer *bufio.Writer) (err error) {
	_, err = writer.ReadFrom(reader)
	if err != nil {
		return
	}

	err = writer.Flush()
	if err != nil {
		return
	}

	return
}

func newLineReader() io.Reader {
	newLine := []byte("\r\n")
	return bytes.NewReader(newLine)
}

func getFileName() string {
	h, m, s := time.Now().Clock()
	return fmt.Sprintf("%d%d%d.txt", h, m, s)
}

func printAndHold(msg string) {
	fmt.Println(msg)
	fmt.Scan()
}
