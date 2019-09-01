package cli

import (
	"os"
)

// Init initializes the mnemonic cli.
func Init() error {
	err := loadConfig()
	if err != nil {
		err = initConfig()
		if err != nil {
			return err
		}

		var dirs = []string{
			DeckDir(),
			NoteDir(),
		}
		err = makeDirs(dirs, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
