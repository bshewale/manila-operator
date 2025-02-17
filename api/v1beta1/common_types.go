/*
Copyright 2020 Red Hat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/endpoint"
	corev1 "k8s.io/api/core/v1"
)

// ManilaTemplate defines common input parameters used by all Manila services
type ManilaTemplate struct {

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=manila
	// ServiceUser - optional username used for this service to register in manila
	ServiceUser string `json:"serviceUser"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=manila
	// DatabaseUser - optional username used for manila DB, defaults to manila
	// TODO: -> implement needs work in mariadb-operator, right now only manila
	DatabaseUser string `json:"databaseUser,omitempty"`

	// +kubebuilder:validation:Optional
	// Secret containing OpenStack password information for ManilaDatabasePassword, AdminPassword
	Secret string `json:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={database: ManilaDatabasePassword, service: ManilaPassword}
	// PasswordSelectors - Selectors to identify the DB and ServiceUser password from the Secret
	PasswordSelectors PasswordSelector `json:"passwordSelectors,omitempty"`
}

// ManilaServiceTemplate defines the input parameters that can be defined for a given
// Manila service
type ManilaServiceTemplate struct {

	// +kubebuilder:validation:Required
	// ContainerImage - Manila API Container Image URL
	ContainerImage string `json:"containerImage"`

	// +kubebuilder:validation:Optional
	// NodeSelector to target subset of worker nodes running this service. Setting here overrides
	// any global NodeSelector settings within the Manila CR.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	// Debug - enable debug for different deploy stages. If an init container is used, it runs and the
	// actual action pod gets started with sleep infinity
	Debug ManilaServiceDebug `json:"debug,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="# add your customization here"
	// CustomServiceConfig - customize the service config using this parameter to change service defaults,
	// or overwrite rendered information using raw OpenStack config format. The content gets added to
	// to /etc/<service>/<service>.conf.d directory as custom.conf file.
	CustomServiceConfig string `json:"customServiceConfig,omitempty"`

	// +kubebuilder:validation:Optional
	// ConfigOverwrite - interface to overwrite default config files like e.g. policy.json.
	// But can also be used to add additional files. Those get added to the service config dir in /etc/<service> .
	DefaultConfigOverwrite map[string]string `json:"defaultConfigOverwrite,omitempty"`

	// +kubebuilder:validation:Optional
	// Resources - Compute Resources required by this service (Limits/Requests).
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// NetworkAttachments is a list of NetworkAttachment resource names to expose the services to the given network
	NetworkAttachments []string `json:"networkAttachments,omitempty"`
}

// PasswordSelector to identify the DB and AdminUser password from the Secret
type PasswordSelector struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="ManilaDatabasePassword"
	// Database - Selector to get the manila database user password from the Secret
	// TODO: not used, need change in mariadb-operator
	Database string `json:"database,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="ManilaPassword"
	// Service - Selector to get the manila service password from the Secret
	Service string `json:"service,omitempty"`
}

// ManilaDebug indicates whether certain stages of Manila deployment should
// pause in debug mode
type ManilaDebug struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// dbInitContainer enable debug (waits until /tmp/stop-init-container disappears)
	DBInitContainer bool `json:"dbInitContainer,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// dbSync enable debug
	DBSync bool `json:"dbSync,omitempty"`
}

// ManilaServiceDebug indicates whether certain stages of Manila service
// deployment should pause in debug mode
type ManilaServiceDebug struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// initContainer enable debug (waits until /tmp/stop-init-container disappears)
	InitContainer bool `json:"initContainer,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// service enable debug
	Service bool `json:"service,omitempty"`
}

// MetalLBConfig to configure the MetalLB loadbalancer service
type MetalLBConfig struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=internal;public
	// Endpoint, OpenStack endpoint this service maps to
	Endpoint endpoint.Endpoint `json:"endpoint"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// IPAddressPool expose VIP via MetalLB on the IPAddressPool
	IPAddressPool string `json:"ipAddressPool"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// SharedIP if true, VIP/VIPs get shared with multiple services
	SharedIP bool `json:"sharedIP"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	// SharedIPKey specifies the sharing key which gets set as the annotation on the LoadBalancer service.
	// Services which share the same VIP must have the same SharedIPKey. Defaults to the IPAddressPool if
	// SharedIP is true, but no SharedIPKey specified.
	SharedIPKey string `json:"sharedIPKey"`

	// +kubebuilder:validation:Optional
	// LoadBalancerIPs, request given IPs from the pool if available. Using a list to allow dual stack (IPv4/IPv6) support
	LoadBalancerIPs []string `json:"loadBalancerIPs,omitempty"`
}
