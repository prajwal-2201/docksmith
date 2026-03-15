package build

type BuildState struct {
	Layers    []string
	WorkDir   string
	Env       map[string]string
	PrevLayer string
	Cmd       []string
}

func NewState() *BuildState {

	return &BuildState{
		Layers:    []string{},
		WorkDir:   "",
		Env:       make(map[string]string),
		PrevLayer: "",
	}
}
