package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jacobkaufmann/mnemonic/cli"
	"github.com/spf13/cobra"
)

const (
	listTitleFmtStr = "%s:\n"
	listItemIndent  = "  "
	listItemFmtStr  = listItemIndent + "- %s"
)

// MakeListCmd returns the mnemonic list command.
func MakeListCmd() *cobra.Command {
	cmdList := &cobra.Command{
		Use:   "list [type to list]",
		Short: "List content of a given type",
	}

	cmdList.AddCommand(cmdListDeck())

	return cmdList
}

func cmdListDeck() *cobra.Command {
	cmdList := &cobra.Command{
		Use:   "deck",
		Short: "List decks",
		RunE: func(cmd *cobra.Command, args []string) error {
			files, err := ioutil.ReadDir(cli.DeckDir())
			if err != nil {
				return err
			}

			fileList := make([]interface{}, len(files))
			for i := range files {
				fileList[i] = files[i]
			}

			fn := func(file interface{}) string {
				f, _ := file.(os.FileInfo)
				return strings.Split(f.Name(), ".")[0]
			}
			cli.DisplayMessage(msgListFn("Decks", fileList, fn))

			return nil
		},
	}

	return cmdList
}

// msgList is a helper function to format a list message from a slice of type
// string.
func msgList(title string, list []string) string {
	var msg strings.Builder

	msg.WriteString(fmt.Sprintf(listTitleFmtStr, title))
	for i, item := range list {
		msg.WriteString(fmt.Sprintf(listItemFmtStr, item))
		if i+1 != len(list) {
			msg.WriteString("\n")
		}
	}

	return msg.String()
}

// msgListFn is a helper function to format a list message from a slice of
// type interface{} where fn maps the interface{} value to its corresponding
// string value.
func msgListFn(title string, list []interface{}, fn func(interface{}) string) string {
	var msg strings.Builder

	msg.WriteString(fmt.Sprintf(listTitleFmtStr, title))
	for i, item := range list {
		msg.WriteString(fmt.Sprintf(listItemFmtStr, fn(item)))
		if i+1 != len(list) {
			msg.WriteString("\n")
		}
	}

	return msg.String()
}
