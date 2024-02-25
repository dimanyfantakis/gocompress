package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func Compress(filename string) string {
	path, _ := filepath.Abs(filename)
	readfile := readFile(path)
	defer readfile.Close()

	message := ""
	freqMap := make(map[rune]int)

	// Create the frequence table.
	// Read the message to be encoded.
	scanner := bufio.NewScanner(readfile)
	for scanner.Scan() {
		line := scanner.Text()
		message += line
		for _, char := range line {
			freqMap[char]++
		}
	}

	// printFreqTable(freqMap)

	root := buildHuffmanTree(freqMap)

	printCodes(root, "")

	codes := make(map[rune]string)
	createCodes(root, "", codes)

	encodedMessage := ""
	for _, char := range message {
		encodedMessage += codes[char]
	}

	fn, _ := writeEncodedFile(message, filename, freqMap, encodedMessage)
	fmt.Println(encodedMessage)
	return fn
}

func writeEncodedFile(message string, filename string, freqMap map[rune]int, encodedMessage string) (string, error) {
	dir, filed := filepath.Split(filename)
	fn := dir + "encoded_" + filed

	file := createFile(fn)
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Sort the keys of the frequency map for deterministic order
	var keys []rune
	for char := range freqMap {
		keys = append(keys, char)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// The header if the frequence table.
	for char, freq := range freqMap {
		_, err := fmt.Fprintf(writer, "%c:%d\n", char, freq)
		if err != nil {
			return "", err
		}
	}

	// Use a new line as a separator between the
	// header and the encoded message.
	_, err := writer.WriteString("\n")
	if err != nil {
		return "", err
	}

	// Write the encoded message to the file.
	_, err = writer.WriteString(encodedMessage)
	if err != nil {
		return "", err
	}

	return fn, nil
}

// Utility functions.

func createCodes(root *Node, code string, codes map[rune]string) {
	if root == nil {
		return
	}

	// Leaf node (contains a character).
	if root.Char != 0 {
		codes[root.Char] = code
	}

	// Traverse left with code '0' and right with code '1'.
	createCodes(root.Left, code+"0", codes)
	createCodes(root.Right, code+"1", codes)
}

func readFile(path string) *os.File {
	readfile, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}
	return readfile
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		os.Exit(1)
	}
	return file
}

func printCodes(root *Node, code string) {
	if root == nil {
		return
	}

	// Leaf node (contains a character)
	if root.Char != 0 {
		fmt.Printf("Character: %c, Code: %s\n", root.Char, code)
	}

	// Traverse left with code '0' and right with code '1'
	printCodes(root.Left, code+"0")
	printCodes(root.Right, code+"1")
}

func printFreqTable(freqMap map[rune]int) {
	fmt.Println("Character Frequencies:")
	for char, freq := range freqMap {
		fmt.Printf("%c: %d\n", char, freq)
	}
}
