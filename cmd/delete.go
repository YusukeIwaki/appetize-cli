package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/pkg/errors"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleting apps",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("api_token") == "" {
			return errors.New("no api token specified")
		}
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		for _, arg := range args {
			options := appetize.DeleteOptions{
				PublicKey: arg,
			}
			deleteResponse, err := client.DeleteApp(options)
			if err != nil {
				fmt.Printf("%s\tERROR:%s\n", arg, err.Error())
			} else if deleteResponse.Body != "OK" {
				fmt.Printf("%s\terror:%s\n", arg, deleteResponse.Body)
			} else {
				fmt.Printf("%s\tOK\n", arg)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
