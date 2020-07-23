package worker

import (
	"context"
	"fmt"
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
	const kubeconfigKey = "kubeconfig"

	secret, err := extensionscontroller.GetSecretByReference(ctx, w.Client(), &w.worker.Spec.SecretRef)
	if err != nil {
		return nil, err
	}

	if secret.Data == nil {
		return nil, fmt.Errorf("secret does not contain any data")
	}

	kubeconfig, ok := secret.Data[kubeconfigKey]
	if !ok {
		return nil, fmt.Errorf("missing %q field in secret", kubeconfigKey)
	}

	return map[string][]byte{
		"kubeconfig": kubeconfig,
	}, nil
}

func (w *workerDelegate) generateMachineConfig(ctx context.Context) error {
	var (
		machineDeployments = worker.MachineDeployments{}
		machineClasses     []map[string]interface{}
		machineImages      []apiskubevirt.MachineImage
	)

	secretData, err := w.generateMachineClassSecretData(ctx)
	if err != nil {
		return err
	}

	for _, pool := range w.worker.Spec.Pools {

		// hardcoded for now as we don't support zones yet
		zoneIdx := int32(0)
		zoneLen := int32(1)

		workerPoolHash, err := worker.WorkerPoolHash(pool, w.cluster)
		if err != nil {
			return err
		}

		imageSourceURL, err := w.findMachineImage(pool.MachineImage.Name, pool.MachineImage.Version)
		if err != nil {
			return err
		}
		machineImages = appendMachineImage(machineImages, apiskubevirt.MachineImage{
			Name:    pool.MachineImage.Name,
			Version: pool.MachineImage.Version,
			SourceURL: imageSourceURL,
		})

		deploymentName := fmt.Sprintf("%s-%s-z", w.worker.Namespace, pool.Name)
		className      := fmt.Sprintf("%s-%s", deploymentName, workerPoolHash)

		machineClasses = append(machineClasses, map[string]interface{}{
			"name": className,
			"storageClassName": "standard",
			"pvcSize": "10Gi",
			"sourceURL": imageSourceURL,
			"cpus": "100m",
			"memory": "4096M",
			"namespace": "default",
			"tags": map[string]string{
				"mcm.gardener.cloud/cluster": w.worker.Namespace,
				"mcm.gardener.cloud/role": "node",
			},
			"secret": map[string]interface{}{
				"cloudConfig": string(pool.UserData),
				"kubeconfig": string(secretData["kubeconfig"]),
			},
		})

		machineDeployments = append(machineDeployments, worker.MachineDeployment{
			Name:           deploymentName,
			ClassName:      className,
			SecretName:     className,
			Minimum:        worker.DistributeOverZones(zoneIdx, pool.Minimum, zoneLen),
			Maximum:        worker.DistributeOverZones(zoneIdx, pool.Maximum, zoneLen),
			MaxSurge:       worker.DistributePositiveIntOrPercent(zoneIdx, pool.MaxSurge, zoneLen, pool.Maximum),
			MaxUnavailable: worker.DistributePositiveIntOrPercent(zoneIdx, pool.MaxUnavailable, zoneLen, pool.Minimum),
			Labels:         pool.Labels,
			Annotations:    pool.Annotations,
			Taints:         pool.Taints,
		})
	}

	w.machineDeployments = machineDeployments
	w.machineClasses = machineClasses
	w.machineImages = machineImages

	return nil
}
