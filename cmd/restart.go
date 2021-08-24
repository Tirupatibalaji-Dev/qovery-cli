package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/qovery/qovery-cli/utils"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Print the restart of your application",
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

		req, err := http.NewRequest("POST", "https://api.qovery.com/application/"+string(application)+"/restart", nil)

		req.Header.Add("Authorization", string(bearer))
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		_ = resp
		if err != nil {
			utils.PrintlnError(err)
			os.Exit(0)
		}

		if resp.StatusCode == 400 {
			fmt.Println("Application " + name + " is already restarted or an optional is in progress")
			os.Exit(0)
		} else if resp.StatusCode >= 400 {
			utils.PrintlnError(errors.New("Received " + resp.Status + " response while listing organizations. "))
			os.Exit(0)
		}

		if resp.StatusCode == 202 {
			fmt.Println("Application " + name + " restart has been requested")
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
