package parser

import "strings"

type Instruction struct {
	Type string
	Args []string
	Raw  string
	Line int
}

func Parse(lines []string) ([]Instruction, error) {

	instructions := []Instruction{}

	for i, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		instr := Instruction{
			Type: strings.ToUpper(parts[0]),
			Args: parts[1:],
			Raw:  line,
			Line: i + 1,
		}

		instructions = append(instructions, instr)
	}

	return instructions, nil
}
