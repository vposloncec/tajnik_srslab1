package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:     "tajnik",
		Short:   "Simple password (credentials) manager",
		Version: "0.3.0",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "mfile", "", "Master file where data is stored(default is $HOME/.tajnik/master_file)")
	viper.SetDefault("author", "Viktor Posloncec <viktor.posloncec@fer.hr>")
	viper.SetDefault("license", "MIT")

	// Create .tajnik directory if it does not exist
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	appHome := filepath.Join(home, ".tajnik")
	if _, err := os.Stat(appHome); os.IsNotExist(err) {
		err := os.Mkdir(appHome, 0755)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Application root directory not found, created", appHome)
	}
	viper.SetDefault("master_file", filepath.Join(appHome, "master_file"))

}
