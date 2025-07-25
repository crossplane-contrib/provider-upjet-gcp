// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionInitParameters struct {

	// A list of URLs of the IP resources used for this NAT rule.
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +listType=set
	SourceNATActiveIps []*string `json:"sourceNatActiveIps,omitempty" tf:"source_nat_active_ips,omitempty"`

	// References to Address in compute to populate sourceNatActiveIps.
	// +kubebuilder:validation:Optional
	SourceNATActiveIpsRefs []v1.Reference `json:"sourceNatActiveIpsRefs,omitempty" tf:"-"`

	// Selector for a list of Address in compute to populate sourceNatActiveIps.
	// +kubebuilder:validation:Optional
	SourceNATActiveIpsSelector *v1.Selector `json:"sourceNatActiveIpsSelector,omitempty" tf:"-"`

	// A list of URLs of the subnetworks used as source ranges for this NAT Rule.
	// These subnetworks must have purpose set to PRIVATE_NAT.
	// This field is used for private NAT.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +listType=set
	SourceNATActiveRanges []*string `json:"sourceNatActiveRanges,omitempty" tf:"source_nat_active_ranges,omitempty"`

	// References to Subnetwork in compute to populate sourceNatActiveRanges.
	// +kubebuilder:validation:Optional
	SourceNATActiveRangesRefs []v1.Reference `json:"sourceNatActiveRangesRefs,omitempty" tf:"-"`

	// Selector for a list of Subnetwork in compute to populate sourceNatActiveRanges.
	// +kubebuilder:validation:Optional
	SourceNATActiveRangesSelector *v1.Selector `json:"sourceNatActiveRangesSelector,omitempty" tf:"-"`

	// A list of URLs of the IP resources to be drained.
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	// +listType=set
	SourceNATDrainIps []*string `json:"sourceNatDrainIps,omitempty" tf:"source_nat_drain_ips,omitempty"`

	// A list of URLs of subnetworks representing source ranges to be drained.
	// This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
	// This field is used for private NAT.
	// +listType=set
	SourceNATDrainRanges []*string `json:"sourceNatDrainRanges,omitempty" tf:"source_nat_drain_ranges,omitempty"`
}

type ActionObservation struct {

	// A list of URLs of the IP resources used for this NAT rule.
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	// +listType=set
	SourceNATActiveIps []*string `json:"sourceNatActiveIps,omitempty" tf:"source_nat_active_ips,omitempty"`

	// A list of URLs of the subnetworks used as source ranges for this NAT Rule.
	// These subnetworks must have purpose set to PRIVATE_NAT.
	// This field is used for private NAT.
	// +listType=set
	SourceNATActiveRanges []*string `json:"sourceNatActiveRanges,omitempty" tf:"source_nat_active_ranges,omitempty"`

	// A list of URLs of the IP resources to be drained.
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	// +listType=set
	SourceNATDrainIps []*string `json:"sourceNatDrainIps,omitempty" tf:"source_nat_drain_ips,omitempty"`

	// A list of URLs of subnetworks representing source ranges to be drained.
	// This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
	// This field is used for private NAT.
	// +listType=set
	SourceNATDrainRanges []*string `json:"sourceNatDrainRanges,omitempty" tf:"source_nat_drain_ranges,omitempty"`
}

type ActionParameters struct {

	// A list of URLs of the IP resources used for this NAT rule.
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceNATActiveIps []*string `json:"sourceNatActiveIps,omitempty" tf:"source_nat_active_ips,omitempty"`

	// References to Address in compute to populate sourceNatActiveIps.
	// +kubebuilder:validation:Optional
	SourceNATActiveIpsRefs []v1.Reference `json:"sourceNatActiveIpsRefs,omitempty" tf:"-"`

	// Selector for a list of Address in compute to populate sourceNatActiveIps.
	// +kubebuilder:validation:Optional
	SourceNATActiveIpsSelector *v1.Selector `json:"sourceNatActiveIpsSelector,omitempty" tf:"-"`

	// A list of URLs of the subnetworks used as source ranges for this NAT Rule.
	// These subnetworks must have purpose set to PRIVATE_NAT.
	// This field is used for private NAT.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta2.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceNATActiveRanges []*string `json:"sourceNatActiveRanges,omitempty" tf:"source_nat_active_ranges,omitempty"`

	// References to Subnetwork in compute to populate sourceNatActiveRanges.
	// +kubebuilder:validation:Optional
	SourceNATActiveRangesRefs []v1.Reference `json:"sourceNatActiveRangesRefs,omitempty" tf:"-"`

	// Selector for a list of Subnetwork in compute to populate sourceNatActiveRanges.
	// +kubebuilder:validation:Optional
	SourceNATActiveRangesSelector *v1.Selector `json:"sourceNatActiveRangesSelector,omitempty" tf:"-"`

	// A list of URLs of the IP resources to be drained.
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceNATDrainIps []*string `json:"sourceNatDrainIps,omitempty" tf:"source_nat_drain_ips,omitempty"`

	// A list of URLs of subnetworks representing source ranges to be drained.
	// This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
	// This field is used for private NAT.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceNATDrainRanges []*string `json:"sourceNatDrainRanges,omitempty" tf:"source_nat_drain_ranges,omitempty"`
}

type Nat64SubnetworkInitParameters struct {

	// Self-link of the subnetwork resource that will use NAT64
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type Nat64SubnetworkObservation struct {

	// Self-link of the subnetwork resource that will use NAT64
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type Nat64SubnetworkParameters struct {

	// Self-link of the subnetwork resource that will use NAT64
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type RouterNATInitParameters struct {

	// The network tier to use when automatically reserving NAT IP addresses.
	// Must be one of: PREMIUM, STANDARD. If not specified, then the current
	// project-level default tier is used.
	// Possible values are: PREMIUM, STANDARD.
	AutoNetworkTier *string `json:"autoNetworkTier,omitempty" tf:"auto_network_tier,omitempty"`

	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	// +listType=set
	DrainNATIps []*string `json:"drainNatIps,omitempty" tf:"drain_nat_ips,omitempty"`

	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty" tf:"enable_dynamic_port_allocation,omitempty"`

	// Enable endpoint independent mapping.
	// For more information see the official documentation.
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty" tf:"enable_endpoint_independent_mapping,omitempty"`

	// Specifies the endpoint Types supported by the NAT Gateway.
	// Supported values include:
	// ENDPOINT_TYPE_VM, ENDPOINT_TYPE_SWG,
	// ENDPOINT_TYPE_MANAGED_PROXY_LB.
	EndpointTypes []*string `json:"endpointTypes,omitempty" tf:"endpoint_types,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec *float64 `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`

	// Configuration for logging on NAT
	// Structure is documented below.
	LogConfig []RouterNATLogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	MaxPortsPerVM *float64 `json:"maxPortsPerVm,omitempty" tf:"max_ports_per_vm,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT. Defaults to 64 for static port allocation and 32 dynamic port allocation if not set.
	MinPortsPerVM *float64 `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`

	// How external IPs should be allocated for this NAT. Valid values are
	// AUTO_ONLY for only allowing NAT IPs allocated by Google Cloud
	// Platform, or MANUAL_ONLY for only user-allocated NAT IP addresses.
	// Possible values are: MANUAL_ONLY, AUTO_ONLY.
	NATIPAllocateOption *string `json:"natIpAllocateOption,omitempty" tf:"nat_ip_allocate_option,omitempty"`

	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +listType=set
	NATIps []*string `json:"natIps,omitempty" tf:"nat_ips,omitempty"`

	// References to Address in compute to populate natIps.
	// +kubebuilder:validation:Optional
	NATIpsRefs []v1.Reference `json:"natIpsRefs,omitempty" tf:"-"`

	// Selector for a list of Address in compute to populate natIps.
	// +kubebuilder:validation:Optional
	NATIpsSelector *v1.Selector `json:"natIpsSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A list of rules associated with this NAT.
	// Structure is documented below.
	Rules []RulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// How NAT should be configured per Subnetwork.
	// If ALL_SUBNETWORKS_ALL_IP_RANGES, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	// Possible values are: ALL_SUBNETWORKS_ALL_IP_RANGES, ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, LIST_OF_SUBNETWORKS.
	SourceSubnetworkIPRangesToNAT *string `json:"sourceSubnetworkIpRangesToNat,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`

	// One or more subnetwork NAT configurations. Only used if
	// source_subnetwork_ip_ranges_to_nat is set to LIST_OF_SUBNETWORKS
	// Structure is documented below.
	Subnetwork []SubnetworkInitParameters `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TCPEstablishedIdleTimeoutSec *float64 `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	TCPTimeWaitTimeoutSec *float64 `json:"tcpTimeWaitTimeoutSec,omitempty" tf:"tcp_time_wait_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TCPTransitoryIdleTimeoutSec *float64 `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UDPIdleTimeoutSec *float64 `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type RouterNATLogConfigInitParameters struct {

	// Indicates whether or not to export logs.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the desired filtering of logs on this NAT.
	// Possible values are: ERRORS_ONLY, TRANSLATIONS_ONLY, ALL.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type RouterNATLogConfigObservation struct {

	// Indicates whether or not to export logs.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the desired filtering of logs on this NAT.
	// Possible values are: ERRORS_ONLY, TRANSLATIONS_ONLY, ALL.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type RouterNATLogConfigParameters struct {

	// Indicates whether or not to export logs.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// Specifies the desired filtering of logs on this NAT.
	// Possible values are: ERRORS_ONLY, TRANSLATIONS_ONLY, ALL.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`
}

type RouterNATObservation struct {

	// The network tier to use when automatically reserving NAT IP addresses.
	// Must be one of: PREMIUM, STANDARD. If not specified, then the current
	// project-level default tier is used.
	// Possible values are: PREMIUM, STANDARD.
	AutoNetworkTier *string `json:"autoNetworkTier,omitempty" tf:"auto_network_tier,omitempty"`

	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	// +listType=set
	DrainNATIps []*string `json:"drainNatIps,omitempty" tf:"drain_nat_ips,omitempty"`

	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty" tf:"enable_dynamic_port_allocation,omitempty"`

	// Enable endpoint independent mapping.
	// For more information see the official documentation.
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty" tf:"enable_endpoint_independent_mapping,omitempty"`

	// Specifies the endpoint Types supported by the NAT Gateway.
	// Supported values include:
	// ENDPOINT_TYPE_VM, ENDPOINT_TYPE_SWG,
	// ENDPOINT_TYPE_MANAGED_PROXY_LB.
	EndpointTypes []*string `json:"endpointTypes,omitempty" tf:"endpoint_types,omitempty"`

	// an identifier for the resource with format {{project}}/{{region}}/{{router}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec *float64 `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`

	// Configuration for logging on NAT
	// Structure is documented below.
	LogConfig []RouterNATLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	MaxPortsPerVM *float64 `json:"maxPortsPerVm,omitempty" tf:"max_ports_per_vm,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT. Defaults to 64 for static port allocation and 32 dynamic port allocation if not set.
	MinPortsPerVM *float64 `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`

	// How external IPs should be allocated for this NAT. Valid values are
	// AUTO_ONLY for only allowing NAT IPs allocated by Google Cloud
	// Platform, or MANUAL_ONLY for only user-allocated NAT IP addresses.
	// Possible values are: MANUAL_ONLY, AUTO_ONLY.
	NATIPAllocateOption *string `json:"natIpAllocateOption,omitempty" tf:"nat_ip_allocate_option,omitempty"`

	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	// +listType=set
	NATIps []*string `json:"natIps,omitempty" tf:"nat_ips,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and NAT reside.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the Cloud Router in which this NAT will be configured.
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// A list of rules associated with this NAT.
	// Structure is documented below.
	Rules []RulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// How NAT should be configured per Subnetwork.
	// If ALL_SUBNETWORKS_ALL_IP_RANGES, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	// Possible values are: ALL_SUBNETWORKS_ALL_IP_RANGES, ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, LIST_OF_SUBNETWORKS.
	SourceSubnetworkIPRangesToNAT *string `json:"sourceSubnetworkIpRangesToNat,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`

	// One or more subnetwork NAT configurations. Only used if
	// source_subnetwork_ip_ranges_to_nat is set to LIST_OF_SUBNETWORKS
	// Structure is documented below.
	Subnetwork []SubnetworkObservation `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TCPEstablishedIdleTimeoutSec *float64 `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	TCPTimeWaitTimeoutSec *float64 `json:"tcpTimeWaitTimeoutSec,omitempty" tf:"tcp_time_wait_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TCPTransitoryIdleTimeoutSec *float64 `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UDPIdleTimeoutSec *float64 `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type RouterNATParameters struct {

	// The network tier to use when automatically reserving NAT IP addresses.
	// Must be one of: PREMIUM, STANDARD. If not specified, then the current
	// project-level default tier is used.
	// Possible values are: PREMIUM, STANDARD.
	// +kubebuilder:validation:Optional
	AutoNetworkTier *string `json:"autoNetworkTier,omitempty" tf:"auto_network_tier,omitempty"`

	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	// +kubebuilder:validation:Optional
	// +listType=set
	DrainNATIps []*string `json:"drainNatIps,omitempty" tf:"drain_nat_ips,omitempty"`

	// Enable Dynamic Port Allocation.
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	// Mutually exclusive with enableEndpointIndependentMapping.
	// +kubebuilder:validation:Optional
	EnableDynamicPortAllocation *bool `json:"enableDynamicPortAllocation,omitempty" tf:"enable_dynamic_port_allocation,omitempty"`

	// Enable endpoint independent mapping.
	// For more information see the official documentation.
	// +kubebuilder:validation:Optional
	EnableEndpointIndependentMapping *bool `json:"enableEndpointIndependentMapping,omitempty" tf:"enable_endpoint_independent_mapping,omitempty"`

	// Specifies the endpoint Types supported by the NAT Gateway.
	// Supported values include:
	// ENDPOINT_TYPE_VM, ENDPOINT_TYPE_SWG,
	// ENDPOINT_TYPE_MANAGED_PROXY_LB.
	// +kubebuilder:validation:Optional
	EndpointTypes []*string `json:"endpointTypes,omitempty" tf:"endpoint_types,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	IcmpIdleTimeoutSec *float64 `json:"icmpIdleTimeoutSec,omitempty" tf:"icmp_idle_timeout_sec,omitempty"`

	// Configuration for logging on NAT
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []RouterNATLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// Maximum number of ports allocated to a VM from this NAT.
	// This field can only be set when enableDynamicPortAllocation is enabled.
	// +kubebuilder:validation:Optional
	MaxPortsPerVM *float64 `json:"maxPortsPerVm,omitempty" tf:"max_ports_per_vm,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT. Defaults to 64 for static port allocation and 32 dynamic port allocation if not set.
	// +kubebuilder:validation:Optional
	MinPortsPerVM *float64 `json:"minPortsPerVm,omitempty" tf:"min_ports_per_vm,omitempty"`

	// How external IPs should be allocated for this NAT. Valid values are
	// AUTO_ONLY for only allowing NAT IPs allocated by Google Cloud
	// Platform, or MANUAL_ONLY for only user-allocated NAT IP addresses.
	// Possible values are: MANUAL_ONLY, AUTO_ONLY.
	// +kubebuilder:validation:Optional
	NATIPAllocateOption *string `json:"natIpAllocateOption,omitempty" tf:"nat_ip_allocate_option,omitempty"`

	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	NATIps []*string `json:"natIps,omitempty" tf:"nat_ips,omitempty"`

	// References to Address in compute to populate natIps.
	// +kubebuilder:validation:Optional
	NATIpsRefs []v1.Reference `json:"natIpsRefs,omitempty" tf:"-"`

	// Selector for a list of Address in compute to populate natIps.
	// +kubebuilder:validation:Optional
	NATIpsSelector *v1.Selector `json:"natIpsSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router and NAT reside.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The name of the Cloud Router in which this NAT will be configured.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Router
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router in compute to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router in compute to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// A list of rules associated with this NAT.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// How NAT should be configured per Subnetwork.
	// If ALL_SUBNETWORKS_ALL_IP_RANGES, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	// Possible values are: ALL_SUBNETWORKS_ALL_IP_RANGES, ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, LIST_OF_SUBNETWORKS.
	// +kubebuilder:validation:Optional
	SourceSubnetworkIPRangesToNAT *string `json:"sourceSubnetworkIpRangesToNat,omitempty" tf:"source_subnetwork_ip_ranges_to_nat,omitempty"`

	// One or more subnetwork NAT configurations. Only used if
	// source_subnetwork_ip_ranges_to_nat is set to LIST_OF_SUBNETWORKS
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Subnetwork []SubnetworkParameters `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	// +kubebuilder:validation:Optional
	TCPEstablishedIdleTimeoutSec *float64 `json:"tcpEstablishedIdleTimeoutSec,omitempty" tf:"tcp_established_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	// Defaults to 120s if not set.
	// +kubebuilder:validation:Optional
	TCPTimeWaitTimeoutSec *float64 `json:"tcpTimeWaitTimeoutSec,omitempty" tf:"tcp_time_wait_timeout_sec,omitempty"`

	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	TCPTransitoryIdleTimeoutSec *float64 `json:"tcpTransitoryIdleTimeoutSec,omitempty" tf:"tcp_transitory_idle_timeout_sec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	// +kubebuilder:validation:Optional
	UDPIdleTimeoutSec *float64 `json:"udpIdleTimeoutSec,omitempty" tf:"udp_idle_timeout_sec,omitempty"`
}

type RulesInitParameters struct {

	// The action to be enforced for traffic that matches this rule.
	// Structure is documented below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	// If it evaluates to true, the corresponding action is enforced.
	// The following examples are valid match expressions for public NAT:
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	// The following example is a valid match expression for private NAT:
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// An integer uniquely identifying a rule in the list.
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`
}

type RulesObservation struct {

	// The action to be enforced for traffic that matches this rule.
	// Structure is documented below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	// If it evaluates to true, the corresponding action is enforced.
	// The following examples are valid match expressions for public NAT:
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	// The following example is a valid match expression for private NAT:
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// An integer uniquely identifying a rule in the list.
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	RuleNumber *float64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`
}

type RulesParameters struct {

	// The action to be enforced for traffic that matches this rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CEL expression that specifies the match condition that egress traffic from a VM is evaluated against.
	// If it evaluates to true, the corresponding action is enforced.
	// The following examples are valid match expressions for public NAT:
	// "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')"
	// "destination.ip == '1.1.0.1' || destination.ip == '8.8.8.8'"
	// The following example is a valid match expression for private NAT:
	// "nexthop.hub == 'https://networkconnectivity.googleapis.com/v1alpha1/projects/my-project/global/hub/hub-1'"
	// +kubebuilder:validation:Optional
	Match *string `json:"match" tf:"match,omitempty"`

	// An integer uniquely identifying a rule in the list.
	// The rule number must be a positive value between 0 and 65000, and must be unique among rules within a NAT.
	// +kubebuilder:validation:Optional
	RuleNumber *float64 `json:"ruleNumber" tf:"rule_number,omitempty"`
}

type SubnetworkInitParameters struct {

	// Self-link of subnetwork to NAT
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// List of the secondary ranges of the subnetwork that are allowed
	// to use NAT. This can be populated only if
	// LIST_OF_SECONDARY_IP_RANGES is one of the values in
	// sourceIpRangesToNat
	// +listType=set
	SecondaryIPRangeNames []*string `json:"secondaryIpRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`

	// List of options for which source IPs in the subnetwork
	// should have NAT enabled. Supported values include:
	// ALL_IP_RANGES, LIST_OF_SECONDARY_IP_RANGES,
	// PRIMARY_IP_RANGE.
	// +listType=set
	SourceIPRangesToNAT []*string `json:"sourceIpRangesToNat,omitempty" tf:"source_ip_ranges_to_nat,omitempty"`
}

type SubnetworkObservation struct {

	// Self-link of subnetwork to NAT
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of the secondary ranges of the subnetwork that are allowed
	// to use NAT. This can be populated only if
	// LIST_OF_SECONDARY_IP_RANGES is one of the values in
	// sourceIpRangesToNat
	// +listType=set
	SecondaryIPRangeNames []*string `json:"secondaryIpRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`

	// List of options for which source IPs in the subnetwork
	// should have NAT enabled. Supported values include:
	// ALL_IP_RANGES, LIST_OF_SECONDARY_IP_RANGES,
	// PRIMARY_IP_RANGE.
	// +listType=set
	SourceIPRangesToNAT []*string `json:"sourceIpRangesToNat,omitempty" tf:"source_ip_ranges_to_nat,omitempty"`
}

type SubnetworkParameters struct {

	// Self-link of subnetwork to NAT
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// List of the secondary ranges of the subnetwork that are allowed
	// to use NAT. This can be populated only if
	// LIST_OF_SECONDARY_IP_RANGES is one of the values in
	// sourceIpRangesToNat
	// +kubebuilder:validation:Optional
	// +listType=set
	SecondaryIPRangeNames []*string `json:"secondaryIpRangeNames,omitempty" tf:"secondary_ip_range_names,omitempty"`

	// List of options for which source IPs in the subnetwork
	// should have NAT enabled. Supported values include:
	// ALL_IP_RANGES, LIST_OF_SECONDARY_IP_RANGES,
	// PRIMARY_IP_RANGE.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceIPRangesToNAT []*string `json:"sourceIpRangesToNat" tf:"source_ip_ranges_to_nat,omitempty"`
}

// RouterNATSpec defines the desired state of RouterNAT
type RouterNATSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterNATParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RouterNATInitParameters `json:"initProvider,omitempty"`
}

// RouterNATStatus defines the observed state of RouterNAT.
type RouterNATStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterNATObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouterNAT is the Schema for the RouterNATs API. A NAT service created in a router.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceSubnetworkIpRangesToNat) || (has(self.initProvider) && has(self.initProvider.sourceSubnetworkIpRangesToNat))",message="spec.forProvider.sourceSubnetworkIpRangesToNat is a required parameter"
	Spec   RouterNATSpec   `json:"spec"`
	Status RouterNATStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterNATList contains a list of RouterNATs
type RouterNATList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterNAT `json:"items"`
}

// Repository type metadata.
var (
	RouterNAT_Kind             = "RouterNAT"
	RouterNAT_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterNAT_Kind}.String()
	RouterNAT_KindAPIVersion   = RouterNAT_Kind + "." + CRDGroupVersion.String()
	RouterNAT_GroupVersionKind = CRDGroupVersion.WithKind(RouterNAT_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterNAT{}, &RouterNATList{})
}
