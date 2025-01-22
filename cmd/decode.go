/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/izumiicxde/morse-code-decoder/service"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode the morse code",
	Long:  `Enter the morse code to decode it to normal string. each letter separated by space and word separated by /`,

	Run: func(cmd *cobra.Command, args []string) {
		text := service.Decode(args[0])
		fmt.Println("Decoded value: ", text)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
