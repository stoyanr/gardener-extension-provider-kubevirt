package main

import (

	"github.com/gardener/gardener-extension-provider-kubevirt/cmd/gardener-extension-provider-kubevirt/app"

	"github.com/gardener/gardener-resource-manager/pkg/log"
	"github.com/gardener/gardener/extensions/pkg/controller"
	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	// TODO: change logger to be more flexible for development purposes
	runtimelog.SetLogger(log.ZapLogger(false))
	cmd := app.NewControllerManagerCommand(controller.SetupSignalHandlerContext())
	if err := cmd.Execute(); err != nil {
		controllercmd.LogErrAndExit(err, "error while starting KubeVirt extension")
	}
}
