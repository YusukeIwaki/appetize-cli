package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Direct file uploads",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
		fmt.Println(args)
		fmt.Println(viper.GetString("platform"))
		fmt.Println(viper.GetString("api_token"))
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// platform
	uploadCmd.PersistentFlags().String("platform", "", "'ios' or 'android'")
	viper.BindPFlag("platform", uploadCmd.PersistentFlags().Lookup("platform"))
}
