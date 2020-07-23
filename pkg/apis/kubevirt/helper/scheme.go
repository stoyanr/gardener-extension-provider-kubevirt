package helper

import (
	"fmt"
	"github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

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

func GetCloudProfileConfigFromProfile(profile *gardencorev1beta1.CloudProfile) (*kubevirt.CloudProfileConfig, error) {
	var cloudProfileConfig *kubevirt.CloudProfileConfig
	if profile.Spec.ProviderConfig != nil && profile.Spec.ProviderConfig.Raw != nil {
		cloudProfileConfig = &kubevirt.CloudProfileConfig{}
		if _, _, err := decoder.Decode(profile.Spec.ProviderConfig.Raw, nil, cloudProfileConfig); err != nil {
			return nil, errors.Wrapf(err, "could not decode providerConfig of cloudProfile")
		}
	}
	return cloudProfileConfig, nil
}

func GetCloudProfileConfig(cluster *controller.Cluster) (*kubevirt.CloudProfileConfig, error) {
	if cluster == nil {
		return nil, nil
	}
	if cluster.CloudProfile == nil {
		return nil, fmt.Errorf("missing cluster cloud profile")
	}
	cloudProfileConfig, err := GetCloudProfileConfigFromProfile(cluster.CloudProfile)
	if err != nil {
		return nil, errors.Wrapf(err, "shoot '%s'", cluster.Shoot.Name)
	}
	return cloudProfileConfig, nil
}

