package devops

import devopsv1alpha1 "devops.io/devops/api/v1alpha1"

type PipelineOperator interface {
	CreatePipeline(p *devopsv1alpha1.Pipeline) error

	RunPipeline(p *devopsv1alpha1.Pipeline) (*devopsv1alpha1.Build, error)
}
