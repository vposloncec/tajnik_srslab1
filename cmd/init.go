package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
  Use:   "init <master_password>",
  Short: "Initializes tajnik using the provided master password",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Tajnik initialized using provided password!")
  },
}