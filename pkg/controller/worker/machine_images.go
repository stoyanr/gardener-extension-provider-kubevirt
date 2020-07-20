package worker

import (
	"context"

	apiskubevirt "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	kubevirtv1alpha1 "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (w *workerDelegate) GetMachineImages(ctx context.Context) (runtime.Object, error) {
	if w.machineImages == nil {
		if err := w.generateMachineConfig(ctx); err != nil {
			return nil, err
		}
	}

	workerStatus := apiskubevirt.WorkerStatus{
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
