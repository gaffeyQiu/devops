package jenkins

import (
	"context"
	devopsv1alpha1 "devops.io/devops/api/v1alpha1"
	"strconv"
	"time"
)

func (cli *jenkinsClient) CreatePipeline(p *devopsv1alpha1.Pipeline) error {
	config, err := createPipelineConfigXml(p.Spec.Pipeline)
	if err != nil {
		return err
	}
	_, err = cli.j.CreateJob(context.Background(), config, p.Name)
	return err
}

func (cli *jenkinsClient) RunPipeline(p *devopsv1alpha1.Pipeline) (*devopsv1alpha1.Build, error) {
	ctx := context.Background()
	queueid, err := cli.j.BuildJob(ctx, p.Name, nil)
	if err != nil {
		return nil, err
	}

	build, err := cli.j.GetBuildFromQueueID(ctx, queueid)
	if err != nil {
		return nil, err
	}

	// Wait for build to finish
	for build.IsRunning(ctx) {
		time.Sleep(1 * time.Second)
		build.Poll(ctx)
	}

	return &devopsv1alpha1.Build{
		Number: strconv.Itoa(int(build.GetBuildNumber())),
		Result: build.GetResult(),
	}, nil
}
