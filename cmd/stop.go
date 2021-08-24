package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/qovery/qovery-cli/utils"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Print the stop of your application",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Capture(cmd)
		token, err := utils.GetAccessToken()
		if err != nil {
			utils.PrintlnError(err)
			os.Exit(0)
		}
		application, name, err := utils.CurrentApplication()
		if err != nil {
			utils.PrintlnError(err)
			os.Exit(0)
		}

		var bearer = "Bearer " + token

		req, err := http.NewRequest("POST", "https://api.qovery.com/application/"+string(application)+"/stop", nil)

		req.Header.Add("Authorization", string(bearer))
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		_ = resp
		if err != nil {
			utils.PrintlnError(err)
			os.Exit(0)
		}

		if resp.StatusCode == 202 {
			fmt.Println("Application " + name + " stop has been requested")
		} else {
			fmt.Println("Application " + name + " is already stopped or an optional is in progress")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
