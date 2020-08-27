<p>Packages:</p>
<ul>
<li>
<a href="#kubevirt.provider.extensions.gardener.cloud%2fv1alpha1">kubevirt.provider.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1">kubevirt.provider.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the kubevirt provider API resources.</p>
</p>
Resource Types:
<ul><li>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>
</li><li>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>
</li><li>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig</a>
</li><li>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus</a>
</li></ul>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig
</h3>
<p>
<p>CloudProfileConfig contains provider-specific configuration that is embedded into Gardener&rsquo;s <code>CloudProfile</code>
resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
kubevirt.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>CloudProfileConfig</code></td>
</tr>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImages">
[]MachineImages
</a>
</em>
</td>
<td>
<p>MachineImages is the list of machine images that are understood by the controller. It maps
logical names and versions to provider-specific identifiers.</p>
</td>
</tr>
<tr>
<td>
<code>machineDeploymentConfig</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineDeploymentConfig">
[]MachineDeploymentConfig
</a>
</em>
</td>
<td>
<p>MachineDeploymentConfigs is the list of machine type deployment config.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig
</h3>
<p>
<p>ControlPlaneConfig contains configuration settings for the control plane.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
kubevirt.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>ControlPlaneConfig</code></td>
</tr>
<tr>
<td>
<code>cloudControllerManager</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">
CloudControllerManagerConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudControllerManager contains configuration settings for the cloud-controller-manager.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig
</h3>
<p>
<p>InfrastructureConfig is the infrastructure configuration resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
kubevirt.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>InfrastructureConfig</code></td>
</tr>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksConfig">
NetworksConfig
</a>
</em>
</td>
<td>
<p>Networks is the configuration of the infrastructure networks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus
</h3>
<p>
<p>WorkerStatus contains information about created worker resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
kubevirt.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>WorkerStatus</code></td>
</tr>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImage">
[]MachineImage
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing <code>Worker</code>
resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
reconciliation is possible.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">CloudControllerManagerConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>)
</p>
<p>
<p>CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>featureGates</code></br>
<em>
map[string]bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>FeatureGates contains information about enabled feature gates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.InfrastructureStatus">InfrastructureStatus
</h3>
<p>
<p>InfrastructureStatus contains information about the status of the infrastructure resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksStatus">
NetworksStatus
</a>
</em>
</td>
<td>
<p>Networks is the status of the infrastructure networks.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineDeploymentConfig">MachineDeploymentConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
<p>MachineDeploymentConfig defines deployment config for VMs for specific machine type</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>machineTypeName</code></br>
<em>
string
</em>
</td>
<td>
<p>MachineTypeName is the name of the machine type, used as a reference to MachineType object</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImage">MachineImage
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus</a>)
</p>
<p>
<p>MachineImage is a mapping from logical names and versions to provider-specific machine image data.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the logical version of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>sourceUrl</code></br>
<em>
string
</em>
</td>
<td>
<p>SourceURL is the url of the machine image</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">MachineImageVersion
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages</a>)
</p>
<p>
<p>MachineImageVersion contains a version and a provider-specific identifier.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the version of the image.</p>
</td>
</tr>
<tr>
<td>
<code>sourceURL</code></br>
<em>
string
</em>
</td>
<td>
<p>SourceURL is the url of the image</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
<p>MachineImages is a mapping from logical names and versions to provider-specific identifiers.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>versions</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">
[]MachineImageVersion
</a>
</em>
</td>
<td>
<p>Versions contains versions and a provider-specific identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworkAttachmentDefinitionReference">NetworkAttachmentDefinitionReference
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksConfig">NetworksConfig</a>)
</p>
<p>
<p>NetworkAttachmentDefinitionReference represents a NetworkAttachmentDefinition reference.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the referenced NetworkAttachmentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace is the namespace of the referenced NetworkAttachmentDefinition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksConfig">NetworksConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig</a>)
</p>
<p>
<p>NetworksConfig contains information about the configuration of the infrastructure networks.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sharedNetworks</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworkAttachmentDefinitionReference">
[]NetworkAttachmentDefinitionReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SharedNetworks is a list of existing networks that can be shared between multiple clusters, e.g. storage networks.</p>
</td>
</tr>
<tr>
<td>
<code>tenantNetworks</code></br>
<em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.TenantNetwork">
[]TenantNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TenantNetworks is a list of &ldquo;tenant&rdquo; networks that are only used by this cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksStatus">NetworksStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.InfrastructureStatus">InfrastructureStatus</a>)
</p>
<p>
<p>NetworksStatus contains information about the status of the infrastructure networks.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkNames</code></br>
<em>
[]string
</em>
</td>
<td>
<p>NetworkNames are the names (in the format <name> or <namespace>/<name>) of all networks
(both shared and tenant) specified in the network configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="kubevirt.provider.extensions.gardener.cloud/v1alpha1.TenantNetwork">TenantNetwork
</h3>
<p>
(<em>Appears on:</em>
<a href="#kubevirt.provider.extensions.gardener.cloud/v1alpha1.NetworksConfig">NetworksConfig</a>)
</p>
<p>
<p>TenantNetwork represents a &ldquo;tenant&rdquo; network that is only used by a single cluster.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the tenant network.</p>
</td>
</tr>
<tr>
<td>
<code>config</code></br>
<em>
string
</em>
</td>
<td>
<p>Config is the configuration of the tenant network.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
