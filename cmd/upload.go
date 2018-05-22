package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YusukeIwaki/appetize-cli/appetize"
	"github.com/pkg/errors"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Direct file uploads",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := appetize.Client{
			ApiToken: viper.GetString("api_token"),
		}
		filePath, err := filepath.Abs(strings.Join(args, " "))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to get absolute path for filepath=%s", filePath))
		}
		options := appetize.UploadOptions{
			Platform:    viper.GetString("platform"),
			AbsFilePath: filePath,
		}
		uploadResponse, err := client.Upload(options)
		if err != nil {
			return errors.Wrap(err, "failed to upload file")
		}
		fmt.Println(uploadResponse.PublicKey)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// platform
	uploadCmd.PersistentFlags().String("platform", "", "'ios' or 'android'")
	viper.BindPFlag("platform", uploadCmd.PersistentFlags().Lookup("platform"))
}
