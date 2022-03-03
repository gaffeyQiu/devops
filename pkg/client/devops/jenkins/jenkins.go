package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
)

type jenkinsClient struct {
	j *gojenkins.Jenkins
}

func CreateJenkinsClient(url string, username string, password string) *jenkinsClient {
	jenkins := gojenkins.CreateJenkins(nil, url, username, password)
	_, err := jenkins.Poll(context.Background())
	if err != nil {
		panic(err)
	}
	return &jenkinsClient{j: jenkins}
}
