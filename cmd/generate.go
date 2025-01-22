package cmd

import (
	"fmt"

	"github.com/izumiicxde/morse-code-decoder/service"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "morse code generator",
	Long:  `Generate a morse code for the given text`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("The morse code for %s is %s", args[0], service.Generate(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
