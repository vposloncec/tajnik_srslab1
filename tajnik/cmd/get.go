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
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get <master_password> <address>",
	Short: "Get credentials for desired address",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mf := viper.GetString("master_file")
		r, err := os.OpenFile(mf, os.O_RDONLY, 0644)
		defer r.Close()
		if err != nil {
			log.Fatalln(err)
		}
		data, err := persistance.LoadDecrypt(args[0], r)
		if err != nil {
			log.Fatalln(err)
		}
		password, ok := data[args[1]]
		if !ok {
			fmt.Println("No passwords found for desired address: ", args[1])
			return
		}
		fmt.Println(password.Get())
	},
}
