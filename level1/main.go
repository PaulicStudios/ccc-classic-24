package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	exampleFileName, err := filepath.Glob("./in/level*_example.in")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	inputFile := exampleFileName[0]
	ExecFile(inputFile)
	exampleOut, err := filepath.Glob("./in/level*_example.out")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	ownExampleOut, err := filepath.Glob("./out/level*_example.out")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	exampleFile, err := os.Open(exampleOut[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	scanner := bufio.NewScanner(exampleFile)
	ownExampleFile, err := os.Open(ownExampleOut[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	scannerOwn := bufio.NewScanner(ownExampleFile)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		scannerOwn.Scan()
		lineOwn := scannerOwn.Text()
		if len(line) != len(lineOwn) {
			fmt.Println("Error: example output files do not match in line: ", count)
			fmt.Println("Expected:", line)
			fmt.Println("Got:", lineOwn)
		}
		count++
	}

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
