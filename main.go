package main

import (
	"flag"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var accessToken string
var owner string

func main() {
	token := flag.String("t", "", "access token")
	milestone := flag.String("m", "", "milestone")
	owner := flag.String("o", "", "owner")
	repo := flag.String("r", "", "repository")
	flag.Parse()

	if token == nil || milestone == nil || owner == nil || repo == nil {
		log.Fatalln("The parameter is not enough")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	opt := github.IssueListByRepoOptions{
		Milestone: *milestone,
		State:     "closed",
	}

	issues, _, err := client.Issues.ListByRepo(*owner, *repo, &opt)
	if err != nil {
		log.Fatalln(err)
	}

	for _, issue := range issues {
		log.Printf("[%s](%s)", *issue.Title, *issue.HTMLURL)
	}
}
