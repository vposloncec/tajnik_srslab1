package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
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
    fmt.Println("We put something in")
  },
}