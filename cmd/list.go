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
	Long:  "Listing apps\n\nREMARK: handling hasMore/nextKey is not supported yet :(",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		options := appetize.ListOptions{}
		listResponse, err := client.ListApps(options)
		if err != nil {
			return errors.Wrap(err, "failed to list apps")
		}
		for _, data := range listResponse.Data {
			fmt.Printf("data:\t%s\t%s\n", data.PublicKey, data.Created)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
