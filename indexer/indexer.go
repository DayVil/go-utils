package indexer

import (
	"io/fs"
	"path/filepath"
)

type FilesAndFolders struct {
	Files   []string
	Folders []string
}

// IsForbiddenDir checks if the directory is forbidden
func IsForbiddenDir(name string) bool {
	return name == "System Volume Information" || name == "$RECYCLE.BIN" || name == "RECYCLER" || name == ".Trash-1000" || name == ".git" || name == ".svn" || name == ".idea" || name == ".vscode" || name == "node_modules" || name == "vendor" || name == "." || name == ".."
}

// IndexFiles indexes all files and folders in the given directory except the forbidden ones
func IndexFiles(root string) (*FilesAndFolders, error) {
	fAndF := FilesAndFolders{}
	root = filepath.Clean(root)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if IsForbiddenDir(info.Name()) {
			return filepath.SkipDir
		}

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
		return &fAndF, err
	}
	return &fAndF, nil
}

// IndexFiles indexes all files and folders in the given directory
func IndexAllFiles(root string) (*FilesAndFolders, error) {
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
		return &fAndF, err
	}
	return &fAndF, nil
}
