package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

	// return

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
	scanner.Scan()
	scanner.Scan()

	for {
		line := scanner.Text()
		if line == "" {
			break
		}

		println("Parsing Number:", line)
		widthStr := strings.Split(line, " ")[0]
		heightStr := strings.Split(line, " ")[1]
		width, err := strconv.Atoi(widthStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		arrbase := make([][]int, height)
		for i := range arrbase {
			arrbase[i] = make([]int, width)
		}

		x, y := 0, 0
		for {
			scanner.Scan()
			line = scanner.Text()
			if line == "" {
				break
			}
			if line[0] == '.' || line[0] == 'X' {
				x = 0
				for _, c := range line {
					if c == '.' {
						arrbase[y][x] = 1
					} else if c == 'X' {
						arrbase[y][x] = 0
					}
					x++
				}
				y++
			} else {
				break
			}
		}

		fmt.Println("Input:")
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				fmt.Printf("%d ", arrbase[i][j])
			}
			fmt.Println()
		}

		if !BackTrack(0, 0, arrbase, "", outFile, width, height) {
			panic("Error: No path found")
		}
	}
}

func PathIsValid(x int, y int, arr [][]int, width int, height int) bool {
	if x < 0 || y < 0 || x >= width || y >= height {
		return false
	}
	if arr[y][x] == 1 {
		return true
	}
	return false
}

func BackTrack(x int, y int, arr [][]int, str string, outFile *os.File, width int, height int) bool {
	if !PathIsValid(x, y, arr, width, height) {
		return false
	}

	if arr[y][x] == 1 {
		arr[y][x] = 2
		if BackTrack(x+1, y, arr, str+"D", outFile, width, height) {
			return true
		}
		if BackTrack(x, y+1, arr, str+"S", outFile, width, height) {
			return true
		}
		if BackTrack(x-1, y, arr, str+"A", outFile, width, height) {
			return true
		}
		if BackTrack(x, y-1, arr, str+"W", outFile, width, height) {
			return true
		}

		wrong := false
		for xx := 0; xx < width; xx++ {
			for yy := 0; yy < height; yy++ {
				if arr[yy][xx] != 2 && arr[yy][xx] != 0 {
					wrong = true
					// fmt.Println("------------")
					// for i := 0; i < width; i++ {
					// 	for j := 0; j < height; j++ {
					// 		fmt.Printf("%d ", arr[j][i])
					// 	}
					// 	fmt.Println()
					// }
					break
				}
			}
			if wrong {
				break
			}
		}
		if !wrong {
			fmt.Println("Solution!")
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					fmt.Printf("%d ", arr[j][i])
				}
				fmt.Println()
			}
			fmt.Fprintln(outFile, str)
			return true
		}
		arr[y][x] = 1
	}

	return false
}
