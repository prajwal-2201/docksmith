package build

import (
	"docksmith/cache"
	"docksmith/parser"
	"docksmith/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Builder struct {
	Cache   *cache.Cache
	NoCache bool
}

func (b *Builder) Build(instructions []parser.Instruction) (*BuildState, error) {

	state := NewState()
	cacheBroken := b.NoCache

	for _, inst := range instructions {

		switch inst.Type {

		case "FROM":

			if len(inst.Args) != 1 {
				return nil, fmt.Errorf("FROM requires exactly one argument at line %d", inst.Line)
			}

			base := inst.Args[0]

			// temporary behavior (until image store from Person A)
			fmt.Println("Using base image:", base)

			state.PrevLayer = ""

		case "COPY":

			if len(inst.Args) < 2 {
				return nil, fmt.Errorf("COPY requires source and destination at line %d", inst.Line)
			}

			src := inst.Args[0]

			files, err := CollectFiles(src)
			if err != nil {
				return nil, err
			}

			key := cache.ComputeKey(
				state.PrevLayer,
				inst.Raw,
				state.WorkDir,
				state.Env,
				files,
			)

			if !cacheBroken {
				if digest, ok := b.Cache.Lookup(key); ok {

					fmt.Println(inst.Raw, "[CACHE HIT]")

					state.Layers = append(state.Layers, digest)
					state.PrevLayer = digest

					break
				}
			}

			fmt.Println(inst.Raw, "[CACHE MISS]")

			layer, err := CreateLayer(files)
			if err != nil {
				return nil, err
			}

			digest := utils.ComputeDigest(layer)

			b.Cache.Store(key, digest)

			state.Layers = append(state.Layers, digest)
			state.PrevLayer = digest

			cacheBroken = true

		case "RUN":

			if len(inst.Args) == 0 {
				return nil, fmt.Errorf("RUN requires a command at line %d", inst.Line)
			}

			key := cache.ComputeKey(
				state.PrevLayer,
				inst.Raw,
				state.WorkDir,
				state.Env,
				nil,
			)

			if !b.NoCache && !cacheBroken {
				if digest, ok := b.Cache.Lookup(key); ok {

					fmt.Println(inst.Raw, "[CACHE HIT]")

					state.Layers = append(state.Layers, digest)
					state.PrevLayer = digest

					break
				}
			}

			fmt.Println(inst.Raw, "[CACHE MISS]")

			cmd := exec.Command(inst.Args[0], inst.Args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				return nil, err
			}

			digest := state.PrevLayer

			if !b.NoCache {
				b.Cache.Store(key, digest)
			}

			cacheBroken = true

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
