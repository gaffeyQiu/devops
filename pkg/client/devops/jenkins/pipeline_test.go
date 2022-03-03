package jenkins

import (
	devopsv1alpha1 "devops.io/devops/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestCreatePipeline(t *testing.T) {
	pipeline := &devopsv1alpha1.Pipeline{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pipeline",
			APIVersion: devopsv1alpha1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hello",
			Namespace: "default",
		},
		Spec: devopsv1alpha1.PipelineSpec{
			Action: devopsv1alpha1.PipelineCreate,
			Pipeline: &devopsv1alpha1.NoScmPipeline{
				Name: "hello",
				Discarder: &devopsv1alpha1.DiscarderProperty{
					DaysToKeep: "3",
					NumToKeep:  "3",
				},
				Jenkinsfile: `pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello World'
            }
        }
    }
}`,
			},
		},
	}

	cli := createJenkinsCli()
	err := cli.CreatePipeline(pipeline)
	assert.Nil(t, err)
}

func TestRunPipeline(t *testing.T) {
	pipeline := &devopsv1alpha1.Pipeline{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pipeline",
			APIVersion: devopsv1alpha1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hello",
			Namespace: "default",
		},
		Spec: devopsv1alpha1.PipelineSpec{
			Action: devopsv1alpha1.PipelineRun,
		},
	}

	cli := createJenkinsCli()
	number, err := cli.RunPipeline(pipeline)
	assert.Nil(t, err)
	assert.NotEmpty(t, number)
}
