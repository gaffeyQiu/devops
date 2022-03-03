package jenkins

import (
	"testing"
)

func createJenkinsCli() *jenkinsClient {
	return CreateJenkinsClient("http://localhost:8080", "admin", "B0J9N43z5MlijWzjKAaSem")
}

func TestCreateJenkinsClient(t *testing.T) {
	createJenkinsCli()
}
