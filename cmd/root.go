package cmd

import (
	"io/fs"
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
	viper.SetDefault("master_file", "$HOME/.tajnik/master_file")

	master_file_path := viper.GetString("master_file")
	if _, err := os.Stat(master_file_path); os.IsNotExist(err) {
	    os.Mkdir(filepath.Dir(master_file_path), fs.ModeDir)
	}

}
