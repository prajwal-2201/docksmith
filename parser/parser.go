package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	Type string
	Args []string
	Raw  string
	Line int
}

var allowed = map[string]bool{
	"FROM":    true,
	"COPY":    true,
	"RUN":     true,
	"WORKDIR": true,
	"ENV":     true,
	"CMD":     true,
}

func ParseFile(path string) ([]Instruction, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {

		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0])

		if !allowed[cmd] {
			return nil, fmt.Errorf(
				"unknown instruction at line %d: %s",
				lineNumber,
				cmd,
			)
		}

		inst := Instruction{
			Type: cmd,
			Args: parts[1:],
			Raw:  line,
			Line: lineNumber,
		}

		instructions = append(instructions, inst)
	}

	return instructions, nil
}
