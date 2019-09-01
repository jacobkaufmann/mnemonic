package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jacobkaufmann/mnemonic/device/card"

	"github.com/jacobkaufmann/mnemonic/cli"

	"github.com/spf13/cobra"
)

const (
	flagName  = "name"
	flagNameS = "n"

	prefixQuestion = "Question:"
	prefixAnswer   = "Answer:"
	suffixQuestion = "[press enter for answer]"
)

var (
	// ErrContentNotFound indicates the content requested to study was not
	// found.
	ErrContentNotFound = errors.New("study content not found")
)

// MakeStudyCmd returns the mnemonic study command.
func MakeStudyCmd() *cobra.Command {
	cmdStudy := &cobra.Command{
		Use:   "study [type to study]",
		Short: "study content of a given type",
	}

	cmdStudy.PersistentFlags().StringP(
		flagName,
		flagNameS,
		"",
		"name of content (required)",
	)
	cmdStudy.MarkFlagRequired(flagName)

	cmdStudy.AddCommand(cmdStudyDeck())

	return cmdStudy
}

// cmdStudyDeck defines the study command for type deck.
func cmdStudyDeck() *cobra.Command {
	cmdStudy := &cobra.Command{
		Use:   "deck",
		Short: "study a deck",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}

			var deck card.Deck
			path := filepath.Join(cli.DeckDir(), name+fileExtensionJSON)
			err = unmarshalContentJSON(path, &deck)
			if err != nil {
				return err
			}

			toStudy := deck.Study(card.FilterKeepAll, true)
			for _, v := range toStudy {
				q, a := v.Query()

				_, err := cli.PromptUser(questionPrompt(q))
				if err != nil {
					return err
				}
				cli.DisplayMessage(fmt.Sprintf("%s %s\n", prefixAnswer, a))
			}

			return nil
		},
	}

	return cmdStudy
}

func questionPrompt(q string) string {
	return fmt.Sprintf("%s %s %s", prefixQuestion, q, suffixQuestion)
}

func unmarshalContentJSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
