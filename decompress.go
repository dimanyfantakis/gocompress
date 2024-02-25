package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func Decompress(filename string) string {

	freqMap := make(map[rune]int)
	encoded, _ := readHeader(filename, freqMap)

	root := buildHuffmanTree(freqMap)

	var result string
	currentNode := root

	for _, bit := range encoded {
		if bit == '0' {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		if currentNode.Char != 0 {
			result += string(currentNode.Char)
			currentNode = root
		}
	}

	fmt.Println(result)
	return result
}

// Utility functions.

func readHeader(filename string, freqMap map[rune]int) (string, error) {
	file := readFile(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Empty line indicates end of header
			break
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line format: %s", line)
		}

		char := []rune(parts[0])[0] // Extract the character from string
		freq := 0
		fmt.Sscanf(parts[1], "%d", &freq)
		freqMap[char] = freq
	}

	// Sort the keys of the frequency map for deterministic order
	var keys []rune
	for char := range freqMap {
		keys = append(keys, char)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	encodedMessage := ""
	// Read the remaining part of the file as the encoded message
	for scanner.Scan() {
		encodedMessage += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return encodedMessage, nil
}
