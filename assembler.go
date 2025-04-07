package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbolTable = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

var nextVariableAddress = 16

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: assembler <file.asm>")
		return
	}

	inputFile := os.Args[1]
	outputFile := strings.TrimSuffix(inputFile, ".asm") + ".hack"

	// First pass: Build symbol table
	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines = firstPass(lines)

	// Second pass: Translate to binary
	binaryCode := secondPass(lines)

	// Write to output file
	err = writeLines(outputFile, binaryCode)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "//") {
			line = strings.Split(line, "//")[0] // Remove comments
			lines = append(lines, strings.TrimSpace(line))
		}
	}
	return lines, scanner.Err()
}

func firstPass(lines []string) []string {
	var result []string
	instructionAddress := 0

	for _, line := range lines {
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			// Label declaration (e.g., (LOOP))
			label := line[1 : len(line)-1]
			symbolTable[label] = instructionAddress
		} else {
			result = append(result, line)
			instructionAddress++
		}
	}
	return result
}

func secondPass(lines []string) []string {
	var result []string

	for _, line := range lines {
		if strings.HasPrefix(line, "@") {
			// A-instruction
			symbol := line[1:]
			address := getAddress(symbol)
			result = append(result, fmt.Sprintf("0%015b", address))
		} else {
			// C-instruction
			result = append(result, translateCInstruction(line))
		}
	}
	return result
}

func getAddress(symbol string) int {
	if address, ok := symbolTable[symbol]; ok {
		return address
	}
	if num, err := strconv.Atoi(symbol); err == nil {
		return num
	}
	symbolTable[symbol] = nextVariableAddress
	nextVariableAddress++
	return symbolTable[symbol]
}

func translateCInstruction(line string) string {
	dest, comp, jump := "", line, ""
	if strings.Contains(line, "=") {
		parts := strings.Split(line, "=")
		dest = parts[0]
		comp = parts[1]
	}
	if strings.Contains(comp, ";") {
		parts := strings.Split(comp, ";")
		comp = parts[0]
		jump = parts[1]
	}

	return "111" + compToBinary(comp) + destToBinary(dest) + jumpToBinary(jump)
}

func compToBinary(comp string) string {
	compTable := map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}
	return compTable[comp]
}

func destToBinary(dest string) string {
	destTable := map[string]string{
		"":    "000",
		"M":   "001",
		"D":   "010",
		"MD":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}
	return destTable[dest]
}

func jumpToBinary(jump string) string {
	jumpTable := map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}
	return jumpTable[jump]
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
