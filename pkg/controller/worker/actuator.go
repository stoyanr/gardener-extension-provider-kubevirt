package worker

import (
	"context"
	apiskubevirt "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/helper"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/common"
	"github.com/gardener/gardener/extensions/pkg/controller/worker"
	"github.com/gardener/gardener/extensions/pkg/controller/worker/genericactuator"
	"github.com/gardener/gardener/extensions/pkg/util"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gardener "github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/go-logr/logr"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/imagevector"
	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/kubevirt"
)

// NewActuator creates a new Actuator that updates the status of the handled WorkerPoolConfigs.
func NewActuator() worker.Actuator {
	delegateFactory := &delegateFactory{
		logger: log.Log.WithName("worker-actuator"),
	}

	return genericactuator.NewActuator(
		log.Log.WithName("kubevirt-worker-actuator"),
		delegateFactory,
		kubevirt.MachineControllerManagerName,
		mcmChart,
		mcmShootChart,
		imagevector.ImageVector(),
		extensionscontroller.ChartRendererFactoryFunc(util.NewChartRendererForShoot),
	)
}

type delegateFactory struct {
	logger logr.Logger
	common.RESTConfigContext
}

func (d *delegateFactory) WorkerDelegate(ctx context.Context, worker *extensionsv1alpha1.Worker, cluster *extensionscontroller.Cluster) (genericactuator.WorkerDelegate, error) {
	clientset, err := kubernetes.NewForConfig(d.RESTConfig())
	if err != nil {
		return nil, err
	}

	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}

	seedChartApplier, err := gardener.NewChartApplierForConfig(d.RESTConfig())
	if err != nil {
		return nil, err
	}

	return NewWorkerDelegate(
		d.ClientContext,
		seedChartApplier,
		serverVersion.GitVersion,

		worker,
		cluster,
	)
}

type workerDelegate struct {
	common.ClientContext

	seedChartApplier gardener.ChartApplier
	serverVersion    string

	cloudProfileConfig *apiskubevirt.CloudProfileConfig
	cluster            *extensionscontroller.Cluster
	worker             *extensionsv1alpha1.Worker

	machineClasses     []map[string]interface{}
	machineDeployments worker.MachineDeployments
	machineImages      []apiskubevirt.MachineImage
}

// NewWorkerDelegate creates a new context for a worker reconciliation.
func NewWorkerDelegate(
	clientContext common.ClientContext,

	seedChartApplier gardener.ChartApplier,
	serverVersion string,

	worker *extensionsv1alpha1.Worker,
	cluster *extensionscontroller.Cluster,
) (genericactuator.WorkerDelegate, error) {
	config, err := helper.GetCloudProfileConfig(cluster)
	if err != nil {
		return nil, err
	}
	return &workerDelegate{
		ClientContext: clientContext,

		seedChartApplier: seedChartApplier,
		serverVersion:    serverVersion,

		cloudProfileConfig: config,
		cluster:            cluster,
		worker:             worker,
	}, nil
}
