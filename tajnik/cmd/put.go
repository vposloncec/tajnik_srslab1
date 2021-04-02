package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vposloncec/lab1-srs/persistance"
)

func init() {
	rootCmd.AddCommand(putCmd)
}

var putCmd = &cobra.Command{
	Use:   "put <master_password> <address> <password>",
	Short: "Put (add) password to the manager's database",
	Long: `Tajnik stores one's password for a desired address using a master password. 
Password for that address can later be queried using the get command.`,

	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		mfPath := viper.GetString("master_file")
		mf, err := os.OpenFile(mfPath, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		data, err := persistance.LoadDecrypt(args[0], mf)
		if err != nil {
			log.Fatalln(err)
		}
		mf.Close()

		data[args[1]] = persistance.PassWithPadding(args[2])

		mf, err = os.OpenFile(mfPath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer mf.Close()
		if err != nil {
			log.Fatalln(err)
		}
		err = data.SaveEncrypt(args[0], mf)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
