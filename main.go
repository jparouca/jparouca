package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/guptarohit/asciigraph"
)

const (
	height     = 15
	width      = 100
	lpFile     = "lp.txt"
	readmeFile = "README.md"
)

func main() {
	lpValues, err := readLPValues("lp.txt")
	if err != nil {
		log.Fatalf("lp values %v", err)
	}

	graph := asciigraph.Plot(lpValues, asciigraph.Height(15), asciigraph.Width(100))
	content := fmt.Sprintf(`# ♟︎ LoL Ratings Chart #
LP History:
%s

`, graph)
	err = os.WriteFile("README.md", []byte(content), 0644)
	if err != nil {
		log.Fatalf("oh I love golang erros xD: %v", err)
	}

	fmt.Println("uptaed")
}

func readLPValues(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lpValues []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lp, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing lp: %v", err)
		}
		lpValues = append(lpValues, lp)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lpValues, nil
}
