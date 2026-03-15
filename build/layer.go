package build

import (
	"path/filepath"
	"sort"
)

func CollectFiles(pattern string) ([]string, error) {

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// deterministic order
	sort.Strings(matches)

	return matches, nil
}
