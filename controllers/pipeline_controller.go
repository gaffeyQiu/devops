/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	devopsClient "devops.io/devops/pkg/client/devops"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	devopsv1alpha1 "devops.io/devops/api/v1alpha1"
)

// PipelineReconciler reconciles a Pipeline object
type PipelineReconciler struct {
	client.Client
	Scheme       *runtime.Scheme
	DevopsClient devopsClient.Interface
}

//+kubebuilder:rbac:groups=devops.devops.io,resources=pipelines,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=devops.devops.io,resources=pipelines/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=devops.devops.io,resources=pipelines/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pipeline object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *PipelineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// 1. 获取到 cr 结构体
	var pipeline devopsv1alpha1.Pipeline
	if err := r.Get(ctx, req.NamespacedName, &pipeline); err != nil {
		if errors.IsNotFound(err) {
			klog.Errorf("未获取到 pipeline 资源，%s", err.Error())
		}

		klog.Errorf("发生错误 %s", err.Error())
		return ctrl.Result{}, nil
	}

	var devopsClientErr error
	var build *devopsv1alpha1.Build
	switch pipeline.Spec.Action {
	case devopsv1alpha1.PipelineCreate:
		devopsClientErr = r.DevopsClient.CreatePipeline(&pipeline)
	case devopsv1alpha1.PipelineUpdate:
		// todo()
	case devopsv1alpha1.PipelineDelete:
		// todo()
	case devopsv1alpha1.PipelineRun:
		build, devopsClientErr = r.DevopsClient.RunPipeline(&pipeline)
	default:
		klog.Errorf("未能识别的 pipeline action: %s", pipeline.Spec.Action)
		return ctrl.Result{}, nil
	}

	if devopsClientErr != nil {
		klog.Errorf("devopsclient has err: %s", devopsClientErr.Error())
		return ctrl.Result{}, nil
	}

	if build != nil {
		pipeline.Status.LastBuild = build
		if err := r.Update(ctx, &pipeline); err != nil {
			klog.Errorf("Update Pipeline Status Build error: %s", err.Error())
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PipelineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsv1alpha1.Pipeline{}).
		Complete(r)
}
