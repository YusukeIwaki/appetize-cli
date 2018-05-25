package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/YusukeIwaki/appetize-cli/optional"
	"github.com/pkg/errors"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creating Apps",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("api_token") == "" {
			return errors.New("no api token specified")
		}
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		options := appetize.CreateOptions{
			Url:      args[0],
			Platform: viper.GetString("create.platform"),
		}
		flags := cmd.Flags()
		if flags.Changed("timeout") {
			if timeout, err := flags.GetInt("timeout"); err == nil {
				options.Timeout = optional.NewInt(timeout)
			}
		}
		if flags.Changed("note") {
			if note, err := flags.GetString("note"); err == nil {
				options.Note = optional.NewString(note)
			}
		}
		if flags.Changed("launch-url") {
			if launchUrl, err := flags.GetString("launch-url"); err == nil {
				options.LaunchUrl = optional.NewString(launchUrl)
			}
		}
		createResponse, err := client.CreateApp(options)
		if err != nil {
			return errors.Wrap(err, "failed to create app")
		}
		fmt.Printf("PublicKey:\t%s\n", createResponse.PublicKey)
		fmt.Printf("Created:\t%s\n", createResponse.Created)
		fmt.Printf("Updated:\t%s\n", createResponse.Updated)
		if createResponse.Timeout > 0 {
			fmt.Printf("Timeout:\t%d\n", createResponse.Timeout)
		}
		fmt.Printf("Platform:\t%s\n", createResponse.Platform)
		fmt.Printf("VersionCode:\t%d\n", createResponse.VersionCode)
		fmt.Printf("Note:\t%s\n", createResponse.Note)
		fmt.Printf("LaunchUrl:\t%s\n", createResponse.LaunchUrl)
		fmt.Printf("PublicUrl:\t%s\n", createResponse.PublicUrl)
		fmt.Printf("AppUrl:\t%s\n", createResponse.AppUrl)
		fmt.Printf("ManageUrl:\t%s\n", createResponse.ManageUrl)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String("platform", "", "'ios' or 'android'")
	viper.BindPFlag("create.platform", createCmd.PersistentFlags().Lookup("platform"))

	createCmd.Flags().String("launch-url", "", "specify a deep link to bring your users to a specific location when your app is launched.")
	createCmd.Flags().Int("timeout", 0, "the number of seconds to wait until automatically ending the session due to user inactivity. Must be 30, 60, 90, 120, 180, 300 or 600, default is 120")
	createCmd.Flags().String("note", "", "a note for your own purposes, will appear on your management dashboard. set 'null' to delete")
}
