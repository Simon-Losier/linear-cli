package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

var getTasks = &cobra.Command{
	Use:   "tasks",
	Short: "Get tasks",
	Long: `Get all tasks.
	
	Example:
	linear-cli tasks -s Todo -p "Project"
	`,
	Run: getAllIssues,
}

type GetAllTasksResponse struct {
	Issues struct {
		Nodes []struct {
			Id         string
			Title      string
			Identifier string
			Project    struct {
				Name string
			}
			State struct {
				Name string
			}
		}
	}
}

type filterStruct struct {
	Project struct {
		Name struct {
			Contains string `json:"contains"`
		} `json:"name"`
	} `json:"project"`
	Team struct {
		Name struct {
			Contains string `json:"contains"`
		} `json:"name"`
	} `json:"team"`
	State struct {
		Name struct {
			Contains string `json:"contains"`
		} `json:"name"`
	} `json:"state"`
}

func init() {
	rootCmd.AddCommand(getTasks)
	getTasks.Flags().StringP("status", "s", "$nil", "The status")
	getTasks.Flags().StringP("project", "p", "$nil", "The project")
}

func getAllIssues(cmd *cobra.Command, args []string) {
	status, _ := cmd.Flags().GetString("status")
	project, _ := cmd.Flags().GetString("project")
	filter := filterStruct{}
	if status != "$nil" {
		fmt.Println("Status: " + status)
		filter.State.Name.Contains = status
	}
	if project != "$nil" {
		fmt.Println("Project: " + project)
		filter.Project.Name.Contains = project
	}

	client := graphql.NewClient("https://api.linear.app/graphql")
	query := `
	query($issuesFilter2: IssueFilter) {
		issues(filter: $issuesFilter2) {
		  nodes {
			  id
			  title
			  identifier
			  project {
				  name
			  }
			  state {
				  name
			  }
		  }
	  }
  }
	`
	request := graphql.NewRequest(query)
	if os.Getenv("apiKey") == "" {
		panic("API Key empty, please add you API key to `apiKey` env varable")
	}
	request.Header.Add("Authorization", os.Getenv("apiKey"))
	request.Var("issuesFilter2", filter)

	var response GetAllTasksResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	for _, el := range response.Issues.Nodes {
		fmt.Println("[" + el.Identifier + "] " + "(" + el.State.Name + ") " + el.Title)
	}
}
