package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/pkg/errors"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Retrieve information for a single app",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("api_token") == "" {
			return errors.New("no api token specified")
		}
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		options := appetize.ShowOptions{
			PublicKey: args[0],
		}
		showResponse, err := client.ShowApp(options)
		if err != nil {
			return errors.Wrap(err, "failed to show app")
		}
		fmt.Printf("PublicKey:\t%s\n", showResponse.PublicKey)
		fmt.Printf("Created:\t%s\n", showResponse.Created)
		fmt.Printf("Updated:\t%s\n", showResponse.Updated)
		fmt.Printf("Disabled:\t%t\n", showResponse.Disabled)
		if showResponse.Timeout > 0 {
			fmt.Printf("Timeout:\t%d\n", showResponse.Timeout)
		}
		fmt.Printf("Platform:\t%s\n", showResponse.Platform)
		fmt.Printf("VersionCode:\t%d\n", showResponse.VersionCode)
		fmt.Printf("Bundle:\t%s\n", showResponse.Bundle)
		fmt.Printf("Name:\t%s\n", showResponse.Name)
		fmt.Printf("Note:\t%s\n", showResponse.Note)
		fmt.Printf("AppVersionName:\t%s\n", showResponse.AppVersionName)
		fmt.Printf("AppVersionCode:\t%s\n", showResponse.AppVersionCode)
		fmt.Printf("IconUrl:\t%s\n", showResponse.IconUrl)
		fmt.Printf("LaunchUrl:\t%s\n", showResponse.LaunchUrl)
		fmt.Printf("ViewUrl:\t%s\n", showResponse.ViewUrl())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
