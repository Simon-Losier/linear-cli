package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

type createTaskSuccess struct {
	IssueCreate struct {
		Success bool
		Issue   struct {
			Id         string
			Title      string
			Identifier string
		}
	}
}

var createTask = &cobra.Command{
	Use:   "new",
	Short: "Create new task",
	Long: `
	Create a new task

	Example:
	linear-cli new -p exampleProject -c 6
	`,
	Run: createNewTask,
}

func init() {
	rootCmd.AddCommand(createTask)
	createTask.Flags().StringP("title", "t", "$nil", "Name of Task")
	createTask.Flags().StringP("description", "d", "", "Description of task")
	createTask.Flags().StringP("project", "p", "$nil", "Name of the project")
	createTask.Flags().StringP("team", "T", "$nil", "Name of the team")
}

func createNewTask(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	description, _ := cmd.Flags().GetString("description")
	project, _ := cmd.Flags().GetString("project")
	team, _ := cmd.Flags().GetString("team")
	if title == "$nil" {
		fmt.Println("No title set.")
		return
	} //
	client := graphql.NewClient("https://api.linear.app/graphql")
	query := fmt.Sprintf(`
	mutation IssueCreate {
		issueCreate(
		  input: {
			title: "%s"
			description: "%s"
			projectId: "%s"
			teamId: "%s"
		  }
		) {
		  success
		  issue {
			id
			title
			identifier
		  }
		}
	  }
	`, title, description, getProjectID(project), getTeamID(team))
	request := graphql.NewRequest(query)
	if os.Getenv("apiKey") == "" {
		fmt.Println("apiKey empty")
		return
	}
	request.Header.Add("Authorization", os.Getenv("apiKey"))
	var response createTaskSuccess
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	if !response.IssueCreate.Success {
		fmt.Println("Task creation failure")
	}
	fmt.Println()
	issue := response.IssueCreate.Issue
	fmt.Println("[" + issue.Identifier + "] " + issue.Title)
}

// Get Team ID ------
type teamIdResonseStruct struct {
	Teams struct {
		Nodes []struct {
			Id string
		}
	}
}

type teamIdFilterStruct struct {
	Name struct {
		Contains string `json:"contains"`
	} `json:"name"`
}

func getTeamID(teamName string) string {
	client := graphql.NewClient("https://api.linear.app/graphql")
	query := `
	query Query($filter: TeamFilter) {
		teams(filter: $filter) {
		  nodes {
			id
		  }
		}
	  }
	`
	request := graphql.NewRequest(query)
	if os.Getenv("apiKey") == "" {
		fmt.Println("Please enter api key to `apiKey` env variable")
		panic("No ENV api")
	}
	request.Header.Add("Authorization", os.Getenv("apiKey"))
	filter := teamIdFilterStruct{}
	filter.Name.Contains = teamName
	request.Var("filter", filter)

	var response teamIdResonseStruct
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	return response.Teams.Nodes[0].Id
}

// Get Project ID --------
type projectIdFilterStruct struct {
	Name struct {
		Contains string `json:"contains"`
	} `json:"name"`
}
type projectIdResponseStruct struct {
	Projects struct {
		Nodes []struct {
			Id string
		}
	}
}

func getProjectID(projectName string) string {
	client := graphql.NewClient("https://api.linear.app/graphql")
	query := `
	query Project($filter: ProjectFilter) {

		projects(filter: $filter) {
		  nodes {
			id
		  }
		}
	  }
	`
	request := graphql.NewRequest(query)
	if os.Getenv("apiKey") == "" {
		fmt.Println("Please enter api key to `apiKey` env variable")
		panic("No ENV api")
	}
	request.Header.Add("Authorization", os.Getenv("apiKey"))
	filter := projectIdFilterStruct{}
	filter.Name.Contains = projectName
	request.Var("filter", filter)

	var response projectIdResponseStruct
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	if len(response.Projects.Nodes) == 0 {
		panic("No teams found, do note search is case sensitive")
	}
	return response.Projects.Nodes[0].Id
}
