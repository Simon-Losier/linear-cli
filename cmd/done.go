package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

var setTaskDone = &cobra.Command{
	Use:   "done",
	Short: "Set task as done",
	Long: `Sets a task using the identifyer (eg: ABC-21) as done,
	
	Example
	linear-cli done ABC-01
	`,
	Run: setIssueDone,
}

func init() {
	rootCmd.AddCommand(setTaskDone)
}

type getTaskRespnse struct {
	IssueUpdate struct {
		Success bool
		Issue   struct {
			Id         string
			Title      string
			Identifier string
			State      struct {
				Id   string
				Name string
			}
		}
	}
}

func setIssueDone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter task identifier eg: ABC-12")
		return
	}
	stateId := "7c4d9600-7c90-4f61-8bd3-004f1a89d489"
	client := graphql.NewClient("https://api.linear.app/graphql")
	query := fmt.Sprintf(`
	mutation IssueUpdate {
		issueUpdate(
		  id: "%s"
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
	`, args[0], stateId)
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
