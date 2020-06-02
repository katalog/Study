package main

import (
	"context"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

// Response of APi
type Response struct {
	License struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"license"`
}

func main() {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request to GitHub API
	req := graphql.NewRequest(`
	query {
		viewer {
		  name
		   repositories(last: 4) {
			 nodes {
			   name
			 }
		   }
		 }
	  }
`)
	//ea09c2500977a89476914e9486912ff5f2b1f5f5
	//var GithubToken = os.Getenv("GITHUB_TOKEN")
	var GithubToken = os.Getenv("ea09c2500977a89476914e9486912ff5f2b1f5f5")
	req.Header.Add("Authorization", "bearer "+GithubToken)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData.License.Description)
}
