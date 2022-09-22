/*
Copyright © 2022 Raylin Gloor <raylin366@hotmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/R4yl1n/getUUIDwin/sendMail"
	"github.com/spf13/cobra"
)

// UUIDCmd represents the UUID command
var UUIDCmd = &cobra.Command{
	Use:   "UUID",
	Short: "Example: get UUID Raylin366@hotmail.com",
	Long: `first call the command and add as a Param your Email adress

	please check if you wrote it right if it doesn't work call 0763970026`,

	Run: func(cmd *cobra.Command, args []string) {

		Out, err := exec.Command("powershell", "wmic path win32_computersystemproduct get uuid").Output()
		if err != nil {
			log.Fatal("command command invalid")
		}

		if len(args) >= 1 && args[0] != "" {
			mail := args[0]

			sendMail.Sendto(mail, string(Out))
			fmt.Println(string(Out))

		} else {
			fmt.Println("please insert your email if you want to get it on your mail example: get UUID raylin.gloor@sunrise.net")
			fmt.Println(string(Out))
		}
	},
}

func init() {
	rootCmd.AddCommand(UUIDCmd)
}

