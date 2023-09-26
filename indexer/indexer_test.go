package indexer

import (
	"path/filepath"
	"slices"
	"testing"
)

func TestIndexFiles(t *testing.T) {
	root := filepath.Join("..", "test-folders")
	root = filepath.Clean(root)
	files := []string{
		filepath.Join(root, "foo", "foo.txt"),
		filepath.Join(root, "foo", "bar", "bar.txt"),
		filepath.Join(root, "foo", "suu", "sa", "ru.md"),
		filepath.Join(root, "foo", "suu", "sa", "sa.txt"),
	}

	folders := []string{
		filepath.Join(root),
		filepath.Join(root, "foo"),
		filepath.Join(root, "foo", "bar"),
		filepath.Join(root, "foo", "suu"),
		filepath.Join(root, "foo", "suu", "sa"),
	}

	indexedFiles, err := IndexFiles(root)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	t.Run("Amount of Elements", func(t *testing.T) {
		if len(indexedFiles.Files) != len(files) {
			t.Fatalf("Amount of files is %d, expected %d\n", len(indexedFiles.Files), len(files))
		}
		if len(indexedFiles.Folders) != len(folders) {
			t.Fatalf("Amount of folders is %d, expected %d\n", len(indexedFiles.Folders), len(folders))
		}
	})

	t.Run("Files", func(t *testing.T) {
		for _, file := range indexedFiles.Files {
			if !slices.Contains(files, file) {
				t.Fatalf("File %s not found in files\n", file)
			}
		}
	})

	t.Run("Folders", func(t *testing.T) {
		for _, folder := range indexedFiles.Folders {
			if !slices.Contains(folders, folder) {
				t.Fatalf("Folder %s not found in folders\n", folder)
			}
		}
	})
}
