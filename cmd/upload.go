package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/pkg/errors"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Direct file uploads",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("api_token") == "" {
			return errors.New("no api token specified")
		}
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		detectArg := func(args []string) (string, string) {
			if len(args) >= 2 {
				return args[0], args[1]
			} else {
				return "", args[0]
			}
		}
		publicKeyArg, filePathArg := detectArg(args)
		filePath, err := filepath.Abs(filePathArg)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to get absolute path for filepath=%s", filePath))
		}
		options := appetize.UploadOptions{
			Platform:    viper.GetString("platform"),
			AbsFilePath: filePath,
			PublicKey:   publicKeyArg,
		}
		uploadResponse, err := client.Upload(options)
		if err != nil {
			return errors.Wrap(err, "failed to upload file")
		}
		fmt.Printf("PublicKey:\t%s\n", uploadResponse.PublicKey)
		fmt.Printf("Created:\t%s\n", uploadResponse.Created)
		fmt.Printf("Updated:\t%s\n", uploadResponse.Updated)
		fmt.Printf("Platform:\t%s\n", uploadResponse.Platform)
		fmt.Printf("VersionCode:\t%d\n", uploadResponse.VersionCode)
		fmt.Printf("PublicUrl:\t%s\n", uploadResponse.PublicUrl)
		fmt.Printf("AppUrl:\t%s\n", uploadResponse.AppUrl)
		fmt.Printf("ManageUrl:\t%s\n", uploadResponse.ManageUrl)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// platform
	uploadCmd.PersistentFlags().String("platform", "", "'ios' or 'android'")
	viper.BindPFlag("platform", uploadCmd.PersistentFlags().Lookup("platform"))
}
