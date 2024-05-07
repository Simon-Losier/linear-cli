package cmd

import (
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

var setTask = &cobra.Command{
	Use:   "done",
	Short: "Set task as done",
	Long: `Sets a task using the identifyer (eg: ABC-21) as done,
	
	Example
	linear-cli done ABC-01
	`,
	Run: setIssueDone,
}

func init() {
	rootCmd.AddCommand(setTask)
}

type getTaskRespnse struct {
	Issues struct {
		Nodes []struct {
			Id         string
			Title      string
			Identifier string
		}
	}
}

func setIssueDone(cmd *cobra.Command, args []string) {
	fmt.Println(args[0])
	client := graphql.NewClient()
}
