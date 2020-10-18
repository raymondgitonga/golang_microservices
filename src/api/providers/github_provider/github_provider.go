package github_provider

import (
	"fmt"
	"github.com/raymondgitonga/golang_microservices/src/api/dormain/github"
)

const (
	headerAuthorization       = "Authorizarion"
	headerAuthorizationFormat = "token %s"
)

func getAuthorizationHeader(acessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, acessToken)
}

func CreateRepo(request github.CreateRepoRequest) (github.CreateRepoResponse, github.GithubErrorResponse) {
	header := getAuthorizationHeader("4e35daaf4d8f571b981767a939bd390f605e9ef6")

	return github.CreateRepoResponse{}, github.GithubErrorResponse{}
}
