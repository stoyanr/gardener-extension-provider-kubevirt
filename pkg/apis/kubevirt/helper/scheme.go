package helper

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/gardener/gardener/extensions/pkg/controller"

	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
	"github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt/install"
)

var (
	// Scheme is a scheme with the types relevant for KubeVirt actuators.
	Scheme *runtime.Scheme

	decoder runtime.Decoder
)

func init() {
	Scheme = runtime.NewScheme()
	utilruntime.Must(install.AddToScheme(Scheme))

	decoder = serializer.NewCodecFactory(Scheme).UniversalDecoder()
}

func GetControlPlaneConfig(cluster *controller.Cluster) (*kubevirt.ControlPlaneConfig, error) {
	cpConfig := &kubevirt.ControlPlaneConfig{}
	if cluster.Shoot.Spec.Provider.ControlPlaneConfig != nil {
		if _, _, err := decoder.Decode(cluster.Shoot.Spec.Provider.ControlPlaneConfig.Raw, nil, cpConfig); err != nil {
			return nil, errors.Wrapf(err, "could not decode providerConfig of controlplane '%s'", cluster.Shoot.Name)
		}
	}
	return cpConfig, nil
}
