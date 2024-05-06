package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var connect = &cobra.Command{
	Use:   "connect",
	Short: "Enter linear.app API key",
	Long: `Enter linear.app cli key Use:
	
	linear-cli connect -k <LINEA.APP API KEY> 
	`,
	Run: addAPIKey,
}

func init() {
	rootCmd.AddCommand(connect)
	connect.Flags().StringP("key", "k", "$nil", "API Key")
}

func addAPIKey(cmd *cobra.Command, args []string) {
	key, _ := cmd.Flags().GetString("key")
	fmt.Println("Set key: " + key)
	os.Setenv("apiKey", key)
	fmt.Println("Key is: " + os.Getenv("apiKey"))
}
