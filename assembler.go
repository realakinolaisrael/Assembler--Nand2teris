package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if the user provided an input file (e.g., go run main.go Prog.asm)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go Prog.asm")
		return
	}

	inputFile := os.Args[1]
	outputFile := strings.Replace(inputFile, ".asm", ".hack", 1)

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the output file
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cleanLine := clean(line)

		if cleanLine == "" {
			continue
		}

		binary, err := translate(cleanLine)
		if err != nil {
			fmt.Println("Error translating line:", err)
			return
		}

		_, err = out.WriteString(binary + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Assembly complete. Output saved to", outputFile)
}

// clean removes whitespace and comments
func clean(line string) string {
	// Remove comments
	if idx := strings.Index(line, "//"); idx != -1 {
		line = line[:idx]
	}
	return strings.TrimSpace(line)
}

// translate converts a single line of Hack assembly into binary
func translate(line string) (string, error) {
	if strings.HasPrefix(line, "@") {
		// A-instruction: @value
		number := line[1:]
		value, err := strconv.Atoi(number)
		if err != nil {
			return "", fmt.Errorf("invalid A-instruction: %s", line)
		}
		return fmt.Sprintf("%016b", value), nil
	}

	// C-instruction placeholder (e.g., D=A, 0;JMP)
	// This is simplified and always returns a dummy binary
	return "1110000000000000", nil
}
