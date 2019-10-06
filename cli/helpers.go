package cli

import (
	"os"
)

// makeDir is a helper function to create directory dir with permissions perm.
func makeDir(dir string, perm os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModeDir)
	}

	err := os.Chmod(dir, perm)
	if err != nil {
		return err
	}

	return nil
}

// makeDirs is a helper function to create multiple directories with
// permissions perm.
func makeDirs(dirs []string, perm os.FileMode) error {
	for _, dir := range dirs {
		err := makeDir(dir, perm)
		if err != nil {
			return err
		}
	}
	return nil
}
