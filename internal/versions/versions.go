/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package versions

import (
	"fmt"
	"strings"
)

// ValidK8sVersion represents any of the three currently supported k8s versions.
type ValidK8sVersion string

// NewValidK8sVersion validates the given string and produces a new ValidK8sVersion object.
func NewValidK8sVersion(k8sVersion string) (ValidK8sVersion, error) {
	if IsSupportedK8sVersion(k8sVersion) {
		return ValidK8sVersion(k8sVersion), nil
	}
	return "", fmt.Errorf("invalid k8sVersion supplied: %s", k8sVersion)
}

// IsSupportedK8sVersion checks if a given Kubernetes version is supported by Constellation.
func IsSupportedK8sVersion(version string) bool {
	switch version {
	case string(V1_23):
		return true
	case string(V1_24):
		return true
	case string(V1_25):
		return true
	default:
		return false
	}
}

// IsPreviewK8sVersion checks if a given Kubernetes version is still in preview and not fully supported.
func IsPreviewK8sVersion(version ValidK8sVersion) bool {
	return false
}

const (
	//
	// Constellation images.
	// These images are built in a way that they support all versions currently listed in VersionConfigs.
	//

	// KonnectivityAgentImage agent image for konnectivity service.
	KonnectivityAgentImage = "us.gcr.io/k8s-artifacts-prod/kas-network-proxy/proxy-agent:v0.0.33@sha256:48f2a4ec3e10553a81b8dd1c6fa5fe4bcc9617f78e71c1ca89c6921335e2d7da" // renovate:container
	// KonnectivityServerImage server image for konnectivity service.
	KonnectivityServerImage = "registry.k8s.io/kas-network-proxy/proxy-server:v0.0.33@sha256:2c111f004bec24888d8cfa2a812a38fb8341350abac67dcd0ac64e709dfe389c" // renovate:container
	// JoinImage image of Constellation join service.
	JoinImage = "ghcr.io/edgelesssys/constellation/join-service:v2.3.0-pre.0.20221109145754-0d12e37c9699@sha256:afe838fdf9753a6d50eef3a36a17d8993c1732397bb3f965dd25f291e7521acc" // renovate:container
	// KmsImage image of Constellation KMS server.
	KmsImage = "ghcr.io/edgelesssys/constellation/kmsserver:v2.3.0-pre.0.20221109145754-0d12e37c9699@sha256:bed58eff5ca1ad2bb0eddfdbb642a5dc5454bfd6a0248487ae8e2756227e0e80" // renovate:container
	// VerificationImage image of Constellation verification service.
	VerificationImage = "ghcr.io/edgelesssys/constellation/verification-service:v2.3.0-pre.0.20221109145754-0d12e37c9699@sha256:4c6887fb6d14cfa1456e95819bc6fff4a360f9c4aa6e57b1fe90ebd469c3a0d6" // renovate:container
	// GcpGuestImage image for GCP guest agent.
	// Check for new versions at https://github.com/GoogleCloudPlatform/guest-agent/releases and update in /.github/workflows/build-gcp-guest-agent.yml.
	GcpGuestImage = "ghcr.io/edgelesssys/gcp-guest-agent:20220927.00@sha256:3dea1ae3f162d2353e6584b325f0e325a39cda5f380f41e5a0ee43c6641d3905" // renovate:container
	// NodeOperatorCatalogImage image of node operator catalog image.
	NodeOperatorCatalogImage = "ghcr.io/edgelesssys/constellation/node-operator-catalog:v2.3.0-pre.0.20221109145754-0d12e37c9699@sha256:d11ec73033bbd698f7d614b836e3bfb1bbf647f2b972df952a0cf5d9c979f795" // renovate:container
	// NodeMaintenanceOperatorCatalogImage image of node maintenance operator catalog.
	// TODO: switch node maintenance operator catalog back to upstream quay.io/medik8s/node-maintenance-operator-catalog
	// once https://github.com/medik8s/node-maintenance-operator/issues/49 is resolved.
	NodeMaintenanceOperatorCatalogImage = "ghcr.io/edgelesssys/constellation/node-maintenance-operator-catalog:v0.13.1-alpha1@sha256:d382c3aaf9bc470cde6f6c05c2c6ff5c9dcfd90540d5b11f9cf69c4e1dd1ca9d" // renovate:container

	// QEMUMetadataImage image of QEMU metadata api service.
	QEMUMetadataImage = "ghcr.io/edgelesssys/constellation/qemu-metadata-api:v2.2.0@sha256:3c173639bbd258f56c7f4e97fa5dc7b7c63d7d45f96f7d7af5c43ed9eb2258ac" // renovate:container
	// LibvirtImage image that provides libvirt.
	LibvirtImage = "ghcr.io/edgelesssys/constellation/libvirt:v2.2.0@sha256:81ddc30cd679a95379e94e2f154861d9112bcabfffa96330c09a4917693f7cce" // renovate:container

	// ConstellationQEMUImageURL is the artifact URL for QEMU qcow2 images.
	ConstellationQEMUImageURL = "https://cdn.confidential.cloud/constellation/images/mini-constellation/v2.2.2/constellation.raw"

	// currently supported versions.
	//nolint:revive
	V1_23 ValidK8sVersion = "1.23"
	//nolint:revive
	V1_24 ValidK8sVersion = "1.24"
	//nolint:revive
	V1_25 ValidK8sVersion = "1.25"

	// Default k8s version deployed by Constellation.
	Default ValidK8sVersion = V1_24
)

var (
	// NodeOperatorVersion version of node operator.
	NodeOperatorVersion = versionFromDockerImage(NodeOperatorCatalogImage)
	// NodeMaintenanceOperatorVersion version of node maintenance operator.
	NodeMaintenanceOperatorVersion = versionFromDockerImage(NodeMaintenanceOperatorCatalogImage)
)

// Regenerate the hashes by running go generate.
// To add another Kubernetes version, add a new entry to the VersionConfigs map below and fill the Hash field with an empty string.
//go:generate go run generateHashes.go

// VersionConfigs holds download URLs for all required kubernetes components for every supported version.
var VersionConfigs = map[ValidK8sVersion]KubernetesVersion{
	V1_23: {
		PatchVersion: "v1.23.17", // renovate:kubernetes-release
		CNIPlugins: ArtifactVersion{
			URL:  "https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz", // renovate:cni-plugins-release
			Hash: "sha256:b275772da4026d2161bf8a8b41ed4786754c8a93ebfb6564006d5da7f23831e5",
		},
		Crictl: ArtifactVersion{
			URL:  "https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.25.0/crictl-v1.25.0-linux-amd64.tar.gz", // renovate:crictl-release
			Hash: "sha256:86ab210c007f521ac4cdcbcf0ae3fb2e10923e65f16de83e0e1db191a07f0235",
		},
		Kubelet: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.23.17/bin/linux/amd64/kubelet", // renovate:kubernetes-release
			Hash: "sha256:f2bef00508790f632d035a6cfdd31539115611bfc93c5a3266ceb95bb2f27b76",
		},
		Kubeadm: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.23.17/bin/linux/amd64/kubeadm", // renovate:kubernetes-release
			Hash: "sha256:46c847e2699839b9ccf6673f0b946c4778a3a2e8e463d15854ba30d3f0cbd87a",
		},
		Kubectl: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.23.17/bin/linux/amd64/kubectl", // renovate:kubernetes-release
			Hash: "sha256:13ce4b18ba6e15d5d259249c530637dd7fb9722d121df022099f3ed5f2bd74cd",
		},
		// CloudControllerManagerImageAWS is the CCM image used on AWS.
		CloudControllerManagerImageAWS: "registry.k8s.io/provider-aws/cloud-controller-manager:v1.23.17@sha256:adcb84e52d1eceb3dd0090f145f0d934cf915babbfe22be7383e52a9e36ad512", // renovate:container
		// CloudControllerManagerImageGCP is the CCM image used on GCP.
		// TODO: use newer "cloud-provider-gcp" from https://github.com/kubernetes/cloud-provider-gcp when newer releases are available.
		CloudControllerManagerImageGCP: "ghcr.io/edgelesssys/cloud-provider-gcp:v23.0.0@sha256:bf54ecb58fef5b1358d1dd25b1068598a74adbc7e7622b42a2708d1ed4bdc4bc", // renovate:container
		// CloudControllerManagerImageAzure is the CCM image used on Azure.
		CloudControllerManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-controller-manager:v1.23.30@sha256:b73c41016aff20868c0b4ac77ddfe77c41ac5a7ad99a3137dc883646e92f93ea", // renovate:container
		// CloudNodeManagerImageAzure is the cloud-node-manager image used on Azure.
		CloudNodeManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.23.30@sha256:c94217a20d054de8178f7fae3756f3b32c3c4ee9441f676972245d35c041d7fe", // renovate:container
		// External service image. Depends on k8s version.
		ClusterAutoscalerImage: "registry.k8s.io/autoscaling/cluster-autoscaler:v1.23.1@sha256:cd2101ba67f3d6ec719f7792d4bdaa3a50e1b716f3a9ccee8931086496c655b7", // renovate:container
	},
	V1_24: {
		PatchVersion: "v1.24.17", // renovate:kubernetes-release
		CNIPlugins: ArtifactVersion{
			URL:  "https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz", // renovate:cni-plugins-release
			Hash: "sha256:b275772da4026d2161bf8a8b41ed4786754c8a93ebfb6564006d5da7f23831e5",
		},
		Crictl: ArtifactVersion{
			URL:  "https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.25.0/crictl-v1.25.0-linux-amd64.tar.gz", // renovate:crictl-release
			Hash: "sha256:86ab210c007f521ac4cdcbcf0ae3fb2e10923e65f16de83e0e1db191a07f0235",
		},
		Kubelet: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.24.17/bin/linux/amd64/kubelet", // renovate:kubernetes-release
			Hash: "sha256:2da0b93857cf352bff5d1eb42e34d398a5971b63a53d8687b45179a78540d6d6",
		},
		Kubeadm: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.24.17/bin/linux/amd64/kubeadm", // renovate:kubernetes-release
			Hash: "sha256:9fea42b4fb5eb2da638d20710ebb791dde221e6477793d3de70134ac058c4cc7",
		},
		Kubectl: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.24.17/bin/linux/amd64/kubectl", // renovate:kubernetes-release
			Hash: "sha256:f93c18751ec715b4d4437e7ece18fe91948c71be1f24ab02a2dde150f5449855",
		},
		// CloudControllerManagerImageAWS is the CCM image used on AWS.
		CloudControllerManagerImageAWS: "registry.k8s.io/provider-aws/cloud-controller-manager:v1.24.15@sha256:08a0d97934206334ff503212b56be2dfb40b235d4547a5db81a036c20d9c581f", // renovate:container
		// CloudControllerManagerImageGCP is the CCM image used on GCP.
		// TODO: use newer "cloud-provider-gcp" from https://github.com/kubernetes/cloud-provider-gcp when newer releases are available.
		CloudControllerManagerImageGCP: "ghcr.io/edgelesssys/cloud-provider-gcp:v24.0.0@sha256:80e2910509ccb4d99b2e08182c2101fbed64f0663194adae08fc1cf878ecc58b", // renovate:container
		// CloudControllerManagerImageAzure is the CCM image used on Azure.
		CloudControllerManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-controller-manager:v1.24.22@sha256:ec6f40834af6e344d372ef27567ff21b83a1ed4988a351698a8758dbcea5c69e", // renovate:container
		// CloudNodeManagerImageAzure is the cloud-node-manager image used on Azure.
		CloudNodeManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.24.22@sha256:eed5da9077e1e7a4e479bf0c29887bc365b57d17baad7afb9106906964205e32", // renovate:container
		// External service image. Depends on k8s version.
		ClusterAutoscalerImage: "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.24.1@sha256:59c8411f69e8cc240568067e864ecc9cbe84338eb2a48512937c5d4e632602f4", // renovate:container
	},
	V1_25: {
		PatchVersion: "v1.25.14", // renovate:kubernetes-release
		CNIPlugins: ArtifactVersion{
			URL:  "https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz", // renovate:cni-plugins-release
			Hash: "sha256:b275772da4026d2161bf8a8b41ed4786754c8a93ebfb6564006d5da7f23831e5",
		},
		Crictl: ArtifactVersion{
			URL:  "https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.25.0/crictl-v1.25.0-linux-amd64.tar.gz", // renovate:crictl-release
			Hash: "sha256:86ab210c007f521ac4cdcbcf0ae3fb2e10923e65f16de83e0e1db191a07f0235",
		},
		Kubelet: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.25.14/bin/linux/amd64/kubelet", // renovate:kubernetes-release
			Hash: "sha256:7f7437e361f829967ee02e30026d7e85219693432ac5e930cc98dd9c7ddb2fac",
		},
		Kubeadm: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.25.14/bin/linux/amd64/kubeadm", // renovate:kubernetes-release
			Hash: "sha256:b8a6119d2a3a7c6add43dcf8f920436bf7fe71a77a086e96e40aa9d6f70be826",
		},
		Kubectl: ArtifactVersion{
			URL:  "https://storage.googleapis.com/kubernetes-release/release/v1.25.14/bin/linux/amd64/kubectl", // renovate:kubernetes-release
			Hash: "sha256:e4e569249798a09f37e31b8b33571970fcfbdecdd99b1b81108adc93ca74b522",
		},
		// CloudControllerManagerImageAWS is the CCM image used on AWS.
		CloudControllerManagerImageAWS: "registry.k8s.io/provider-aws/cloud-controller-manager:v1.25.3@sha256:47eb1c1e6a3bd6d0fb44ac4992885b6218f1448ea339de778d8b703df463c06f", // renovate:container
		// CloudControllerManagerImageGCP is the CCM image used on GCP.
		// TODO: use newer "cloud-provider-gcp" from https://github.com/kubernetes/cloud-provider-gcp when newer releases are available.
		CloudControllerManagerImageGCP: "ghcr.io/edgelesssys/cloud-provider-gcp:v25.2.0@sha256:86fa9d31ed0b3d0d8806f13d6e7debd3471028b2cb7cca3a876d8a31612a7ba5", // renovate:container
		// CloudControllerManagerImageAzure is the CCM image used on Azure.
		// Check for newer versions at https://github.com/kubernetes-sigs/cloud-provider-azure/blob/master/README.md.
		CloudControllerManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-controller-manager:v1.25.19@sha256:9bc7284fe6d55f2095400a44120314d03d306afdac20a54dc908328aa8527424", // renovate:container
		// CloudNodeManagerImageAzure is the cloud-node-manager image used on Azure.
		// Check for newer versions at https://github.com/kubernetes-sigs/cloud-provider-azure/blob/master/README.md.
		CloudNodeManagerImageAzure: "mcr.microsoft.com/oss/kubernetes/azure-cloud-node-manager:v1.25.19@sha256:b38e4ad314b2b004aaac8cce8cc451756b1d6a215dcdb3989d6865d9cdfde3cc", // renovate:container
		// External service image. Depends on k8s version.
		// Check for new versions at https://github.com/kubernetes/autoscaler/releases.
		ClusterAutoscalerImage: "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.25.1@sha256:cfdfd46b726a2d5488d4df9c41b636acbab0dff4c01c8c6dbe3d057539f2f2af", // renovate:container
	},
}

// KubernetesVersion bundles download URLs to all version-releated binaries necessary for installing/deploying a particular Kubernetes version.
type KubernetesVersion struct {
	PatchVersion                     string
	CNIPlugins                       ArtifactVersion // No k8s version dependency.
	Crictl                           ArtifactVersion // k8s version dependency.
	Kubelet                          ArtifactVersion // k8s version dependency.
	Kubeadm                          ArtifactVersion // k8s version dependency.
	Kubectl                          ArtifactVersion // k8s version dependency.
	CloudControllerManagerImageAWS   string          // k8s version dependency.
	CloudControllerManagerImageGCP   string          // Using self-built image until resolved: https://github.com/kubernetes/cloud-provider-gcp/issues/289
	CloudControllerManagerImageAzure string          // k8s version dependency.
	CloudNodeManagerImageAzure       string          // k8s version dependency. Same version as above.
	ClusterAutoscalerImage           string          // Matches k8s versioning scheme.
}

// ArtifactVersion is a version of a particular artifact.
type ArtifactVersion struct {
	URL  string
	Hash string
}

// versionFromDockerImage returns the version tag from the image name, e.g. "v1.22.2" from "foocr.io/org/repo:v1.22.2@sha256:3009fj0...".
func versionFromDockerImage(imageName string) string {
	beforeAt, _, _ := strings.Cut(imageName, "@")
	_, version, ok := strings.Cut(beforeAt, ":")
	if !ok {
		panic(fmt.Errorf("failed to extract version from image name, no ':' found in %s", imageName))
	}
	return version
}
