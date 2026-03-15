package build

import (
	"docksmith/parser"
	"docksmith/utils"
	"fmt"
)

type Builder struct{}

func (b *Builder) Build(instructions []parser.Instruction) error {

	state := NewState()

	for _, inst := range instructions {

		switch inst.Type {

		case "FROM":
			// handled later

		case "COPY":

			if len(inst.Args) < 2 {
				return fmt.Errorf("COPY requires source and destination at line %d", inst.Line)
			}

			src := inst.Args[0]

			files, err := CollectFiles(src)
			if err != nil {
				return err
			}

			layer, err := CreateLayer(files)
			if err != nil {
				return err
			}

			digest := utils.ComputeDigest(layer)

			state.Layers = append(state.Layers, digest)
			state.PrevLayer = digest

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
