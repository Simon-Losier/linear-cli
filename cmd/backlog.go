package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

var setTaskBacklog = &cobra.Command{
	Use:   "backlog",
	Short: "Set task as Backlog",
	Long: `Sets a task using the identifyer (eg: ABC-21) as todo,
	
	Example
	linear-cli todo ABC-01
	`,
	Run: setIssueBacklog,
}

func init() {
	rootCmd.AddCommand(setTaskBacklog)
}

func setIssueBacklog(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter task identifier eg: ABC-12")
		return
	}
	stateId := "f73e1695-349f-44d2-8d11-3f1acdea7d2b"
	client := graphql.NewClient("https://api.linear.app/graphql")
	query := fmt.Sprintf(`
	mutation IssueUpdate {
		issueUpdate(
		  id: "SIM-91"
		  input: {
			stateId: "%s"
		  }
		) {
		  success
		  issue {
			id
			title
			identifier
			state {
			  id
			  name
			}
		  }
		}
	  }
	`, stateId)
	request := graphql.NewRequest(query)
	if os.Getenv("apiKey") == "" {
		fmt.Println("Please enter api key to `apiKey` env variable")
	}
	request.Header.Add("Authorization", os.Getenv("apiKey"))
	var response getTaskRespnse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	if !response.IssueUpdate.Success {
		fmt.Println("Invalid query")
	}
	issue := response.IssueUpdate.Issue
	fmt.Println("[" + issue.Identifier + "] " + "(" + issue.State.Name + ") " + issue.Title)
}
