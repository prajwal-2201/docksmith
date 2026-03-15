package build

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func CollectFiles(pattern string) ([]string, error) {

	var files []string

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {

		info, err := os.Stat(match)
		if err != nil {
			return nil, err
		}

		if info.IsDir() {

			err := filepath.Walk(match, func(path string, info os.FileInfo, err error) error {

				if err != nil {
					return err
				}

				if !info.IsDir() {
					files = append(files, path)
				}

				return nil
			})

			if err != nil {
				return nil, err
			}

		} else {
			files = append(files, match)
		}
	}

	sort.Strings(files)

	return files, nil
}

func CreateLayer(files []string) ([]byte, error) {

	sort.Strings(files)

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	for _, file := range files {

		info, err := os.Stat(file)
		if err != nil {
			return nil, err
		}

		header := &tar.Header{
			Name:    file,
			Size:    info.Size(),
			Mode:    int64(info.Mode()),
			ModTime: time.Unix(0, 0), // critical for reproducible builds
		}

		if err := tw.WriteHeader(header); err != nil {
			return nil, err
		}

		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(tw, f); err != nil {
			f.Close()
			return nil, err
		}

		f.Close()
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
