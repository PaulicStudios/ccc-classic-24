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

	for scanner.Scan() {
		line := scanner.Text()

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
		arrbase := make([][]int, width)
		for i := range arrbase {
			arrbase[i] = make([]int, height)
		}

		x, y := 0, -1
		for {
			y++
			scanner.Scan()
			line = scanner.Text()
			if line[0] == '.' || line[0] == 'X' {
				x = 0
				for _, c := range line {
					if c == '.' {
						arrbase[x][y] = 1
					} else if c == 'X' {
						arrbase[x][y] = 0
					}
					x++
				}
			} else {
				break
			}
		}

		println("Processing:", line)
		var wrong bool = false
		for yStart := 0; yStart < height; yStart++ {
			for xStart := 0; xStart < width; xStart++ {
				arr := make([][]int, len(arrbase))
				for i := range arrbase {
					arr[i] = make([]int, len(arrbase[i]))
					copy(arr[i], arrbase[i][:])
				}
				wrong = false
				x, y = xStart, yStart
				if arr[x][y] == 0 {
					wrong = true
					continue
				}
				arr[x][y] = 2
				for _, c := range line {
					switch c {
					case 'W':
						y--
					case 'A':
						x--
					case 'S':
						y++
					case 'D':
						x++
					default:
						println("Error: Invalid character in path")
					}

					if x < 0 || x >= width || y < 0 || y >= height {
						wrong = true
						break
					}
					if arr[x][y] == 0 || arr[x][y] == 2 {
						wrong = true
						break
					}
					arr[x][y] = 2
				}

				if wrong {
					// fmt.Println("Error: Invalid path asd")
					// for i := 0; i < width; i++ {
					// 	for j := 0; j < height; j++ {
					// 		fmt.Printf("%d ", arr[i][j])
					// 	}
					// 	fmt.Println()
					// }
					continue
				}

				for x := 0; x < width; x++ {
					for y := 0; y < height; y++ {
						if arr[x][y] != 2 && arr[x][y] != 0 {
							wrong = true
							// fmt.Println("Error: Invalid path")
							// for i := 0; i < width; i++ {
							// 	for j := 0; j < height; j++ {
							// 		fmt.Printf("%d ", arr[i][j])
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
					break
				}
			}
			if !wrong {
				break
			}
		}
		if wrong {
			fmt.Fprintln(outFile, "INVALID")
		}
		if !wrong {
			fmt.Fprintln(outFile, "VALID")
		}
	}
}
