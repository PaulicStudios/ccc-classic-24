package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFiles, err := filepath.Glob("./in/level*_*.in")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, inputFile := range inputFiles {
		ExecFile(inputFile)
	}
}

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return file
}

func OutFile(fileName string) *os.File {
	outputFile := strings.Replace(fileName, "in", "out", 2)
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return outFile
}

func ExecFile(fileName string) {
	fmt.Println("Processing file:", fileName)
	file := OpenFile(fileName)
	outFile := OutFile(fileName)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Fprintln(outFile, line+"test")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			break
		}
	}
}
