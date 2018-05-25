package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/YusukeIwaki/appetize-cli/optional"
	"github.com/pkg/errors"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update app settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("api_token") == "" {
			return errors.New("no api token specified")
		}
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		updateForm := appetize.UpdateForm{}
		flags := cmd.Flags()
		if flags.Changed("disabled") {
			if disabled, err := flags.GetBool("disabled"); err == nil {
				updateForm.Disabled = optional.NewBool(disabled)
			}
		}
		if flags.Changed("timeout") {
			if timeout, err := flags.GetInt("timeout"); err == nil {
				updateForm.Timeout = optional.NewInt(timeout)
			}
		}
		if flags.Changed("note") {
			if note, err := flags.GetString("note"); err == nil {
				updateForm.Note = optional.NewString(note)
			}
		}
		if flags.Changed("launch-url") {
			if launchUrl, err := flags.GetString("launch-url"); err == nil {
				updateForm.LaunchUrl = optional.NewString(launchUrl)
			}
		}
		options := appetize.UpdateOptions{
			PublicKey:  args[0],
			UpdateForm: updateForm,
		}
		updateResponse, err := client.UpdateApp(options)
		if err != nil {
			return errors.Wrap(err, "failed to update app")
		}
		fmt.Printf("PublicKey:\t%s\n", updateResponse.PublicKey)
		fmt.Printf("Created:\t%s\n", updateResponse.Created)
		fmt.Printf("Updated:\t%s\n", updateResponse.Updated)
		fmt.Printf("Disabled:\t%t\n", updateResponse.Disabled)
		if updateResponse.Timeout > 0 {
			fmt.Printf("Timeout:\t%t\n", updateResponse.Timeout)
		}
		fmt.Printf("Platform:\t%s\n", updateResponse.Platform)
		fmt.Printf("VersionCode:\t%d\n", updateResponse.VersionCode)
		fmt.Printf("Bundle:\t%s\n", updateResponse.Bundle)
		fmt.Printf("Name:\t%s\n", updateResponse.Name)
		fmt.Printf("Note:\t%s\n", updateResponse.Note)
		fmt.Printf("AppVersionName:\t%s\n", updateResponse.AppVersionName)
		fmt.Printf("AppVersionCode:\t%s\n", updateResponse.AppVersionCode)
		fmt.Printf("IconUrl:\t%s\n", updateResponse.IconUrl)
		fmt.Printf("LaunchUrl:\t%s\n", updateResponse.LaunchUrl)
		fmt.Printf("ViewUrl:\t%s\n", updateResponse.ViewUrl())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("launch-url", "", "specify a deep link to bring your users to a specific location when your app is launched.")
	updateCmd.Flags().Bool("disabled", false, "disables streaming for this app. 'true' or 'false'. (Since appetize API is buggy, we have to always set disabled=true for keeping an app to be disabled.)")
	updateCmd.Flags().Int("timeout", 0, "the number of seconds to wait until automatically ending the session due to user inactivity. Must be 30, 60, 90, 120, 180, 300 or 600, default is 120")
	updateCmd.Flags().String("note", "", "a note for your own purposes, will appear on your management dashboard. set 'null' to delete")
}
