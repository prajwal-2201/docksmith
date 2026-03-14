package build

import (
	"docksmith/parser"
	"fmt"
)

type Builder struct{}

func (b *Builder) Build(instr []parser.Instruction) error {

	state := NewState()

	for _, inst := range instr {

		switch inst.Type {

		case "FROM":
			// load base image

		case "COPY":
			// create layer

		case "RUN":
			// execute command

		case "WORKDIR":
			// update state

		case "ENV":
			// update env

		case "CMD":
			// store cmd

		default:
			return fmt.Errorf("unknown instruction at line %d", inst.Line)

		}
	}

	return nil
}
