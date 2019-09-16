package cmd

import (
	"os"

	"github.com/jacobkaufmann/mnemonic/cli"

	"github.com/jacobkaufmann/mnemonic/device/card"
	"github.com/spf13/cobra"
)

// MakeAddCmd returns the mnemonic add command.
func MakeAddCmd() *cobra.Command {
	cmdAdd := &cobra.Command{
		Use:   "add [type to add]",
		Short: "Add an instance of a given type",
	}

	cmdAdd.AddCommand(cmdAddDeck())

	return cmdAdd
}

// cmdAddDeck defines the add command for type card.
func cmdAddDeck() *cobra.Command {
	var name string

	cmdAdd := &cobra.Command{
		Use:   "deck",
		Short: "Add a deck",
		RunE: func(cmd *cobra.Command, args []string) error {
			deck := card.NewDeck(name)

			filename := cli.DeckDir() + name + fileExtensionJSON
			err := writeFileJSON(filename, &deck, os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmdAdd.Flags().StringVarP(&name, flagName, flagNameS, "", "name for deck (required)")
	cmdAdd.MarkFlagRequired(flagName)

	return cmdAdd
}
