package cli

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner *bufio.Scanner
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

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

// DisplayMessage writes msg to the console.
func DisplayMessage(msg string) {
	fmt.Println(msg)
}

// PromptUser writes prompt to the console and returns the repsonse.
func PromptUser(prompt string) (string, error) {
	_, err := fmt.Print(prompt)
	if err != nil {
		return "", err
	}

	scanner.Scan()
	if err = scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
