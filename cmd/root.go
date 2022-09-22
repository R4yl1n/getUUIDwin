/*
Copyright © 2022 Raylin Gloor <Raylin366@hotmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "getUUIDwin",
	Short: "Example: get UUID Raylin366@hotmail.com",
	Long: `first call the command and add as a Param your Email adress
	
	please check if you wrote it right if it doesn't work call 0763970026`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
