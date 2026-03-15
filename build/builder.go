package build

import "docksmith/parser"

type Builder struct{}

func (b *Builder) Build(instructions []parser.Instruction) error {

	state := NewState()

	for _, inst := range instructions {

		switch inst.Type {

		case "FROM":
			// handled later

		case "COPY":
			// create layer later

		case "RUN":
			// run command later

		case "WORKDIR":
			state.WorkDir = inst.Args[0]

		case "ENV":
			// set env later

		case "CMD":
			// store command later

		}
	}

	return nil
}
