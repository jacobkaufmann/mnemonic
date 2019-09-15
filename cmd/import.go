package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jacobkaufmann/mnemonic/cli"

	"github.com/jacobkaufmann/mnemonic/device/card"
	"github.com/spf13/cobra"
)

// MakeImportCmd returns the mnemonic import command.
func MakeImportCmd() *cobra.Command {
	cmdImport := &cobra.Command{
		Use:   "import [type to import]",
		Short: "Import content of a given type",
	}

	cmdImport.PersistentFlags().StringP(
		flagFilename,
		flagFilenameS,
		"",
		"filename of content (required)",
	)
	cmdImport.MarkFlagRequired(flagFilename)

	cmdImport.AddCommand(cmdImportDeck())

	return cmdImport
}

func cmdImportDeck() *cobra.Command {
	cmdImport := &cobra.Command{
		Use:   "deck",
		Short: "Import a deck of cards",
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, err := cmd.Flags().GetString(flagFilename)
			if err != nil {
				return err
			}

			var deck card.Deck
			err = readFileJSON(filename, &deck)
			if err != nil {
				return err
			}

			// Assign deck name as filename.
			deck.Name = strings.Split(filepath.Base(filename), ".")[0]

			importPath := filepath.Join(cli.DeckDir(), deck.Name+fileExtensionJSON)
			err = writeFileJSON(importPath, &deck, os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmdImport
}
