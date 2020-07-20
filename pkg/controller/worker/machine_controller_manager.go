package worker

import (
	"context"
	"fmt"
	"path/filepath"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/utils/chart"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/kubevirt"
)

var (
	mcmChart = &chart.Chart{
		Name:   kubevirt.MachineControllerManagerName,
		Path:   filepath.Join(kubevirt.InternalChartsPath, kubevirt.MachineControllerManagerName, "seed"),
		Images: []string{kubevirt.MachineControllerManagerImageName, kubevirt.MCMProviderKubeVirtImageName},
		Objects: []*chart.Object{
			{Type: &appsv1.Deployment{}, Name: kubevirt.MachineControllerManagerName},
			{Type: &corev1.Service{}, Name: kubevirt.MachineControllerManagerName},
			{Type: &corev1.ServiceAccount{}, Name: kubevirt.MachineControllerManagerName},
			{Type: &corev1.Secret{}, Name: kubevirt.MachineControllerManagerName},
			{Type: extensionscontroller.GetVerticalPodAutoscalerObject(), Name: kubevirt.MachineControllerManagerVpaName},
			{Type: &corev1.ConfigMap{}, Name: kubevirt.MachineControllerManagerMonitoringConfigName},
		},
	}

	mcmShootChart = &chart.Chart{
		Name: kubevirt.MachineControllerManagerName,
		Path: filepath.Join(kubevirt.InternalChartsPath, kubevirt.MachineControllerManagerName, "shoot"),
		Objects: []*chart.Object{
			{Type: &rbacv1.ClusterRole{}, Name: fmt.Sprintf("extensions.gardener.cloud:%s:%s", kubevirt.Name, kubevirt.MachineControllerManagerName)},
			{Type: &rbacv1.ClusterRoleBinding{}, Name: fmt.Sprintf("extensions.gardener.cloud:%s:%s", kubevirt.Name, kubevirt.MachineControllerManagerName)},
		},
	}
)

func (w *workerDelegate) GetMachineControllerManagerChartValues(ctx context.Context) (map[string]interface{}, error) {
	namespace := &corev1.Namespace{}
	if err := w.Client().Get(ctx, kutil.Key(w.worker.Namespace), namespace); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"providerName": kubevirt.Name,
		"namespace": map[string]interface{}{
			"uid": namespace.UID,
		},
		"podLabels": map[string]interface{}{
			v1beta1constants.LabelPodMaintenanceRestart: "true",
		},
	}, nil
}

func (w *workerDelegate) GetMachineControllerManagerShootChartValues(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"providerName": kubevirt.Name,
	}, nil
}
