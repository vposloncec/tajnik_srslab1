package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vposloncec/lab1-srs/persistance"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init <master_password>",
	Short: "Initializes tajnik using the provided master password",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tajnik initialized using provided password!")
		mfPath := viper.GetString("master_file")
		if _, err := os.Stat(mfPath); err == nil {
			fmt.Println("master_file already exists! use put or get commands instead")
		}

		mf, err := os.Create(mfPath)
		defer mf.Close()
		if err != nil {
			log.Fatalln(err)
		}
		data := persistance.Storage{}
		err = data.SaveEncrypt(args[0], mf)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
