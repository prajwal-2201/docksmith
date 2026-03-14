package interfaces

type LayerStore interface {
	SaveLayer(digest string, data []byte) error
	LayerExists(digest string) bool
}

type ImageStore interface {
	LoadBaseImage(name string) ([]string, error)
	SaveManifest(name string, manifest []byte) error
}

type CacheStore interface {
	Get(key string) (string, bool)
	Set(key string, digest string) error
}
