/*
Copyright © 2023 Ben Lechlitner <otherland@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd *cobra.Command

func NewImportCommand() *cobra.Command {

	importCmd = &cobra.Command{
		Use:   "import",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("import called")
		},
	}

	return importCmd
}
