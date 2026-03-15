package build

import (
	"docksmith/parser"
	"docksmith/utils"
	"fmt"
	"strings"
)

type Builder struct{}

func (b *Builder) Build(instructions []parser.Instruction) (*BuildState, error) {

	state := NewState()

	for _, inst := range instructions {

		switch inst.Type {

		case "FROM":
			// handled later

		case "COPY":

			if len(inst.Args) < 2 {
				return nil, fmt.Errorf("COPY requires source and destination at line %d", inst.Line)
			}

			src := inst.Args[0]

			files, err := CollectFiles(src)
			if err != nil {
				return nil, err
			}

			layer, err := CreateLayer(files)
			if err != nil {
				return nil, err
			}

			digest := utils.ComputeDigest(layer)

			state.Layers = append(state.Layers, digest)
			state.PrevLayer = digest

		case "RUN":
			// run command later

		case "WORKDIR":

			if len(inst.Args) != 1 {
				return nil, fmt.Errorf("WORKDIR requires exactly one argument at line %d", inst.Line)
			}

			state.WorkDir = inst.Args[0]

		case "ENV":

			if len(inst.Args) != 1 {
				return nil, fmt.Errorf("ENV must be KEY=value format at line %d", inst.Line)
			}

			parts := strings.SplitN(inst.Args[0], "=", 2)

			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid ENV format at line %d", inst.Line)
			}

			key := parts[0]
			value := parts[1]

			state.Env[key] = value

		case "CMD":
			state.Cmd = inst.Args

		}
	}

	return state, nil
}
