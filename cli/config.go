package cli

import (
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	cfgType = "yaml"
	cfgFile = "config"
	cfgPath = "$HOME/.mnemonic/"

	appDir = "mnemonic"

	deckDirKey = "DeckDir"
	noteDirKey = "NoteDir"

	defaultDeckDir = "decks"
	defaultNoteDir = "notes"
)

// Init initializes the mnemonic configuration.
func Init() {
	viper.SetConfigType(cfgType)
	viper.SetConfigName(cfgFile)
	viper.AddConfigPath(cfgPath)
	viper.SetDefault(deckDirKey, defaultDeckDir)
	viper.SetDefault(noteDirKey, defaultNoteDir)
	viper.WriteConfig()
	viper.WatchConfig()
}

// Load loads the mnemonic configuration. An error is returned if there is an
// error retrieving the configuration.
func Load() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}

// DeckDir returns the configured deck directory.
func DeckDir() string {
	return filepath.Join(appDir, viper.GetString(deckDirKey))
}

// NoteDir returns the configured note directory.
func NoteDir() string {
	return filepath.Join(appDir, viper.GetString(noteDirKey))
}
