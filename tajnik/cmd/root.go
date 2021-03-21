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
		Use:   "tajnik",
		Short: "Simple password (credentials) manager",
		Version: "0.0.1",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "mfile", "", "Master file where data is stored(default is $HOME/.config/tajnik.yaml)")
	viper.SetDefault("author", "Viktor Posloncec <viktor.posloncec@fer.hr>")
	viper.SetDefault("license", "MIT")

	// Create .tajnik directory if it does not exist
	home, err := os.UserHomeDir()
	if err != nil{
		log.Fatalln(err)
	}
	appHomeDir := filepath.Join(home, ".tajnik")
	if _, err := os.Stat(appHomeDir); os.IsNotExist(err) {
	    err := os.Mkdir(appHomeDir, 0755)
	    if err != nil{
	    	log.Fatalln(err)
		}
	    fmt.Println("Application root directory not found, created", filepath.Dir(appHomeDir))
	}
	viper.SetDefault("master_file", filepath.Join(appHomeDir,"master_file"))

}
