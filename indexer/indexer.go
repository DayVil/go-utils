package indexer

import (
	"io/fs"
	"path/filepath"
)

type FilesAndFolders struct {
	Files   []string
	Folders []string
}

func IndexFiles(root string) (*FilesAndFolders, error) {
	fAndF := FilesAndFolders{}
	root = filepath.Clean(root)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fAndF.Folders = append(fAndF.Folders, path)
		} else {
			fAndF.Files = append(fAndF.Files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &fAndF, nil
}
