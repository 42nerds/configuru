package commands

import (
	"fmt"
	"log"
	"os/user"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// doctorCmd TODO: tbd
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "this helps to debug issues with your configuru",
	Long: `TODO: A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cyan")
		color.Cyan("Prints text in cyan.")
		user, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}
		homeDirectory := user.HomeDir
		fmt.Printf("User: %s\n", user.Username)
		fmt.Printf("Home Directory: %s\n", homeDirectory)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
