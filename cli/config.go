package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	cfgFilename = "config.yaml"
	cfgType     = "yaml"
	cfgDir      = ".mnemonic"

	deckDirKey = "DeckDir"
	noteDirKey = "NoteDir"
)

var (
	home = os.Getenv("HOME")

	cfgPath = filepath.Join(home, cfgDir)
	cfgFile = filepath.Join(cfgPath, cfgFilename)

	defaultDeckDir = filepath.Join(cfgPath, "decks")
	defaultNoteDir = filepath.Join(cfgPath, "notes")
)

func init() {
	viper.SetConfigType(cfgType)
	viper.SetConfigName(cfgFilename)
	viper.AddConfigPath(cfgPath)
}

// initConfig initializes the mnemonic configuration.
func initConfig() error {
	err := makeConfig()
	if err != nil {
		return err
	}

	viper.Set(deckDirKey, defaultDeckDir)
	viper.Set(noteDirKey, defaultNoteDir)
	err = viper.WriteConfigAs(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

// loadConfig loads the mnemonic configuration.
func loadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}

// makeConfig creates the configuration directory and files.
func makeConfig() error {
	var err error

	err = makeDir(cfgPath, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Create(cfgFile)
	if err != nil {
		return err
	}
	err = os.Chmod(cfgFile, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// DeckDir returns the configured deck directory.
func DeckDir() string {
	return viper.GetString(deckDirKey)
}

// NoteDir returns the configured note directory.
func NoteDir() string {
	return viper.GetString(noteDirKey)
}
