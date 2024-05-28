package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	byteCountPtr := flag.Bool("c", false, "Returns byte count for file")
	lineCountPtr := flag.Bool("l", false, "Returns line count for file")
	wordCountPtr := flag.Bool("w", false, "Returns word count for file")
	charCountPtr := flag.Bool("m", false, "Returns character count for file")

	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	if !*byteCountPtr && !*lineCountPtr && !*wordCountPtr && !*charCountPtr {
		byteCount := getByteCount(data)
		lineCount := getLineCount(data)
		wordCount := getWordCount(data)

		fmt.Printf("%d %d %d %s", lineCount, wordCount, byteCount, filename)
	}

	if *byteCountPtr {
		byteCount := getByteCount(data)

		fmt.Printf("%d %s", byteCount, filename)
	}

	if *lineCountPtr {
		lineCount := getLineCount(data)

		fmt.Printf("%d %s", lineCount, filename)
	}

	if *wordCountPtr {
		wordCount := getWordCount(data)

		fmt.Printf("%d %s", wordCount, filename)
	}

	if *charCountPtr {
		stringData := getDataAsString(data)

		fmt.Println(len([]rune(stringData)), filename)
	}
}

func getDataAsString(data []byte) string {
	return string(data)
}

func getByteCount(data []byte) int {
	return len(data)
}

func getLineCount(data []byte) int {
	stringData := getDataAsString(data)
	lines := strings.Split(stringData, "\n")

	return len(lines) - 1
}

func getWordCount(data []byte) int {
	stringData := getDataAsString(data)
	words := strings.Fields(stringData)

	return len(words)
}
