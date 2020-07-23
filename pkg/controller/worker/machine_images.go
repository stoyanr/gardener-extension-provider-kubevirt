package worker

import (
	"context"
	"fmt"

	"github.com/gardener/gardener/extensions/pkg/util"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	apiskubevirt "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	apiskubevirthelper "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/helper"
	kubevirtv1alpha1 "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/v1alpha1"
)

func (w *workerDelegate) GetMachineImages(ctx context.Context) (runtime.Object, error) {
	if w.machineImages == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return nil, err
		}
	}

	workerStatus := &apiskubevirt.WorkerStatus{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiskubevirt.SchemeGroupVersion.String(),
			Kind:       "WorkerStatus",
		},
		MachineImages: w.machineImages,
	}

	workerStatusV1alpha1 := &kubevirtv1alpha1.WorkerStatus{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kubevirtv1alpha1.SchemeGroupVersion.String(),
			Kind:       "WorkerStatus",
		},
	}

	if err := w.Scheme().Convert(workerStatus, workerStatusV1alpha1, nil); err != nil {
		return nil, err
	}
	return workerStatusV1alpha1, nil
}

func (w *workerDelegate) findMachineImage(name, version string) (string, error) {
	if w.cloudProfileConfig != nil {
		sourceURL, err := apiskubevirthelper.FindImage(w.cloudProfileConfig.MachineImages, name, version)
		if err == nil {
			return sourceURL, nil
		}
	}

	// Try to look up machine image in worker provider status as it was not found in componentconfig.
	if providerStatus := w.worker.Status.ProviderStatus; providerStatus != nil {
		workerStatus := &apiskubevirt.WorkerStatus{}
		if _, _, err := w.Decoder().Decode(providerStatus.Raw, nil, workerStatus); err != nil {
			return "", errors.Wrapf(err, "could not decode worker status of worker '%s'", util.ObjectName(w.worker))
		}

		machineImage, err := apiskubevirthelper.FindMachineImage(workerStatus.MachineImages, name, version)
		if err != nil {
			return "", errorMachineImageNotFound(name, version)
		}

		return machineImage.SourceURL, nil
	}

	return "", errorMachineImageNotFound(name, version)
}

func errorMachineImageNotFound(name, version string) error {
	return fmt.Errorf("could not find machine image for %s/%s neither in componentconfig nor in worker status", name, version)
}

func appendMachineImage(machineImages []apiskubevirt.MachineImage, machineImage apiskubevirt.MachineImage) []apiskubevirt.MachineImage {
	if _, err := apiskubevirthelper.FindMachineImage(machineImages, machineImage.Name, machineImage.Version); err != nil {
		return append(machineImages, machineImage)
	}
	return machineImages
}
