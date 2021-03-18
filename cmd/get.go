package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
  Use:   "get <master_password> <address>",
  Short: "Get credentials for desired address",
  Args: cobra.ExactArgs(2),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Credentials!")
  },
}