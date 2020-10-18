package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang Intro",
		Description: "Golang intro repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"Golang Intro","description":"Golang intro repo","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`,
		string(bytes))
}
