package helper

import (
	"fmt"

	api "github.com/gardener/gardener-extension-provider-kubevirt/pkg/apis/kubevirt"
)

// FindMachineImage takes a list of machine images and tries to find the first entry
// whose name, version, and zone matches with the given name, version, and zone. If no such entry is
// found then an error will be returned.
func FindMachineImage(configImages []api.MachineImage, imageName, imageVersion string) (*api.MachineImage, error) {
	for _, machineImage := range configImages {
		if machineImage.Name == imageName && machineImage.Version == imageVersion {
			return &machineImage, nil
		}
	}
	return nil, fmt.Errorf("no machine image with name %q, version %q found", imageName, imageVersion)
}

// FindImage takes a list of machine images, and the desired image name and version. It tries
// to find the image with the given name and version. If it cannot be found then an error
// is returned.
func FindImage(profileImages []api.MachineImages, imageName, imageVersion string) (string, error) {
	for _, machineImage := range profileImages {
		if machineImage.Name == imageName {
			for _, version := range machineImage.Versions {
				if imageVersion == version.Version {
					sourceURL := ""
					if version.SourceURL != "" {
						sourceURL = version.SourceURL
					}
					return sourceURL, nil
				}
			}
		}
	}

	return "", fmt.Errorf("could not find an image for name %q in version %q", imageName, imageVersion)
}
