package main

import (
	"fmt"
	"os"

	"github.com/jacobkaufmann/mnemonic/cli"
	"github.com/jacobkaufmann/mnemonic/cmd"

	"github.com/spf13/cobra"
)

func makeRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "mnemonic",
		Short: "Mnemonic is a general purpose recall tool",
	}

	rootCmd.AddCommand(cmd.MakeListCmd())
	rootCmd.AddCommand(cmd.MakeAddCmd())
	rootCmd.AddCommand(cmd.MakeImportCmd())
	rootCmd.AddCommand(cmd.MakeStudyCmd())

	return rootCmd, nil
}

func main() {
	cmd, err := makeRootCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cli.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
