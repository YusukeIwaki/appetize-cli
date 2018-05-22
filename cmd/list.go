package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/pkg/errors"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing apps",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		optionsKey, _ := cmd.Flags().GetString("key")
		options := appetize.ListOptions{
			Key: optionsKey,
		}
		listResponse, err := client.ListApps(options)
		if err != nil {
			return errors.Wrap(err, "failed to list apps")
		}
		fmt.Println(listResponse)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("key", "", "", "nextKey from previous request.")
}
