/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/42nerds/configuru/internal/app/configuru"
	"gitlab.com/42nerds/configuru/internal/app/files"
	"gitlab.com/42nerds/configuru/internal/app/questions"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [source] [destination]",
	Short: "Initialize a new project from a template",
	Long:  `Initialize a new project from a template`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			fmt.Printf("Missing Argument expected 2\nPlease make sure you provided the template source and its destination.\n")
			os.Exit(1)
		}

		outputDir, err := files.FormatDir(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			log.Fatalln(err)
		}
		_, fileErr := os.Stat(outputDir + "configuru.yaml")
		if fileErr == nil && !force {
			fmt.Printf("Output directory is already a template. Use\n\n\t configuru init --force\n\n to overwrite it.\n If you want to update use\n\n\tconfiguru update\n\ninstead.\n")
			os.Exit(1)
		}

		configuration, err := configuru.UnmarshalconfiguruFile(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		configuration.Answers, err = questions.PromptQuestions(configuration.Questions)
		if err != nil {
			log.Fatalln(err)
		}

		err = configuru.MarshalconfiguruFile(args[1], configuration)
		if err != nil {
			log.Fatalln(err)
		}

		var renderConfig configuru.RenderConfiguration
		err = renderConfig.Fill(configuration)
		if err != nil {
			log.Fatalln(err)
		}
		err = renderConfig.AddTemplates(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		err = renderConfig.Execute(args[1])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("force", "f", false, "Forces the recreation of the templated structure. Ignores all differences in Template Files")
}
