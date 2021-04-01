package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var version string

var rootCmd = &cobra.Command{
	Version: version,
	Use:     "configuru COMMAND",
	Short:   "A brief description of your application",
	Long: `	     __           __    __     __     __   
	|__|  _)  |\ |   |_    |__)   |  \   (_    
	   | /__  | \|.  |__.  | \ .  |__/.  __).  
													  
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.configuru.yaml)")
}

// initConfig TODO: TBD
func initConfig() {

}
