package interfaces

type Runtime interface {
	Run(rootfs string, cmd []string, env map[string]string) error
}
