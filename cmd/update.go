/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/42nerds/configuru/internal/app/configuru"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [destination]",
	Short: "Update a already provided template.",
	Long:  `Update a already provided template.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := configuru.UnmarshalconfiguruFile(args[0])
		if err != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
