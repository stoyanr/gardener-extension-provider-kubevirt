package worker

import (
	"context"
	"path/filepath"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/worker"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"

	apiskubevirt "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/kubevirt"
	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
)

// MachineClassKind yields the name of the KubeVirt machine class.
func (w *workerDelegate) MachineClassKind() string {
	return "MachineClass"
}

// MachineClassList yields a newly initialized KubeVirtMachineClassList object.
func (w *workerDelegate) MachineClassList() runtime.Object {
	return &machinev1alpha1.MachineClassList{}
}

// DeployMachineClasses generates and creates the KubeVirt specific machine classes.
func (w *workerDelegate) DeployMachineClasses(ctx context.Context) error {
	if w.machineClasses == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return err
		}
	}
	return w.seedChartApplier.Apply(
		ctx, filepath.Join(kubevirt.InternalChartsPath, "machine-class"), w.worker.Namespace, "machineclass",
		kubernetes.Values(map[string]interface{}{"machineClasses": w.machineClasses}),
	)
}

// GenerateMachineDeployments generates the configuration for the desired machine deployments.
func (w *workerDelegate) GenerateMachineDeployments(ctx context.Context) (worker.MachineDeployments, error) {
	if w.machineDeployments == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return nil, err
		}
	}
	return w.machineDeployments, nil
}

func (w *workerDelegate) generateMachineClassSecretData(ctx context.Context) (map[string][]byte, error) {
	//secret, err := extensionscontroller.GetSecretByReference(ctx, w.Client(), &w.worker.Spec.SecretRef)
	_, err := extensionscontroller.GetSecretByReference(ctx, w.Client(), &w.worker.Spec.SecretRef)
	if err != nil {
		return nil, err
	}

	return map[string][]byte{
		// TODO: pass kubeconfig here?
	}, nil
}

func (w *workerDelegate) generateMachineConfig(ctx context.Context) error {
	var (
		machineDeployments = worker.MachineDeployments{}
		machineClasses     []map[string]interface{}
		machineImages      []apiskubevirt.MachineImage
	)

	// TODO: check zones
	// TODO: create machine deployment and classes

	w.machineDeployments = machineDeployments
	w.machineClasses = machineClasses
	w.machineImages = machineImages

	return nil
}
