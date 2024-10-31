// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type NetworkSecurityGroupsSecurityRule_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the security rule.
	Properties *SecurityRulePropertiesFormat_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

// Security rule resource.
type SecurityRulePropertiesFormat_STATUS struct {
	// Access: The network traffic is allowed or denied.
	Access *SecurityRuleAccess_STATUS `json:"access,omitempty"`

	// Description: A description for this rule. Restricted to 140 chars.
	Description *string `json:"description,omitempty"`

	// DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to
	// match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty"`

	// DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty"`

	// DestinationApplicationSecurityGroups: The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded `json:"destinationApplicationSecurityGroups,omitempty"`

	// DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used
	// to match all ports.
	DestinationPortRange *string `json:"destinationPortRange,omitempty"`

	// DestinationPortRanges: The destination port ranges.
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty"`

	// Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction *SecurityRuleDirection_STATUS `json:"direction,omitempty"`

	// Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
	// rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `json:"priority,omitempty"`

	// Protocol: Network protocol this rule applies to.
	Protocol *SecurityRulePropertiesFormat_Protocol_STATUS `json:"protocol,omitempty"`

	// ProvisioningState: The provisioning state of the security rule resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// SourceAddressPrefix: The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags
	// such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies
	// where network traffic originates from.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty"`

	// SourceAddressPrefixes: The CIDR or source IP ranges.
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty"`

	// SourceApplicationSecurityGroups: The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded `json:"sourceApplicationSecurityGroups,omitempty"`

	// SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match
	// all ports.
	SourcePortRange *string `json:"sourcePortRange,omitempty"`

	// SourcePortRanges: The source port ranges.
	SourcePortRanges []string `json:"sourcePortRanges,omitempty"`
}

// An application security group in a resource group.
type ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Whether network traffic is allowed or denied.
type SecurityRuleAccess_STATUS string

const (
	SecurityRuleAccess_STATUS_Allow = SecurityRuleAccess_STATUS("Allow")
	SecurityRuleAccess_STATUS_Deny  = SecurityRuleAccess_STATUS("Deny")
)

// Mapping from string to SecurityRuleAccess_STATUS
var securityRuleAccess_STATUS_Values = map[string]SecurityRuleAccess_STATUS{
	"allow": SecurityRuleAccess_STATUS_Allow,
	"deny":  SecurityRuleAccess_STATUS_Deny,
}

// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
type SecurityRuleDirection_STATUS string

const (
	SecurityRuleDirection_STATUS_Inbound  = SecurityRuleDirection_STATUS("Inbound")
	SecurityRuleDirection_STATUS_Outbound = SecurityRuleDirection_STATUS("Outbound")
)

// Mapping from string to SecurityRuleDirection_STATUS
var securityRuleDirection_STATUS_Values = map[string]SecurityRuleDirection_STATUS{
	"inbound":  SecurityRuleDirection_STATUS_Inbound,
	"outbound": SecurityRuleDirection_STATUS_Outbound,
}

type SecurityRulePropertiesFormat_Protocol_STATUS string

const (
	SecurityRulePropertiesFormat_Protocol_STATUS_Ah   = SecurityRulePropertiesFormat_Protocol_STATUS("Ah")
	SecurityRulePropertiesFormat_Protocol_STATUS_Esp  = SecurityRulePropertiesFormat_Protocol_STATUS("Esp")
	SecurityRulePropertiesFormat_Protocol_STATUS_Icmp = SecurityRulePropertiesFormat_Protocol_STATUS("Icmp")
	SecurityRulePropertiesFormat_Protocol_STATUS_Star = SecurityRulePropertiesFormat_Protocol_STATUS("*")
	SecurityRulePropertiesFormat_Protocol_STATUS_Tcp  = SecurityRulePropertiesFormat_Protocol_STATUS("Tcp")
	SecurityRulePropertiesFormat_Protocol_STATUS_Udp  = SecurityRulePropertiesFormat_Protocol_STATUS("Udp")
)

// Mapping from string to SecurityRulePropertiesFormat_Protocol_STATUS
var securityRulePropertiesFormat_Protocol_STATUS_Values = map[string]SecurityRulePropertiesFormat_Protocol_STATUS{
	"ah":   SecurityRulePropertiesFormat_Protocol_STATUS_Ah,
	"esp":  SecurityRulePropertiesFormat_Protocol_STATUS_Esp,
	"icmp": SecurityRulePropertiesFormat_Protocol_STATUS_Icmp,
	"*":    SecurityRulePropertiesFormat_Protocol_STATUS_Star,
	"tcp":  SecurityRulePropertiesFormat_Protocol_STATUS_Tcp,
	"udp":  SecurityRulePropertiesFormat_Protocol_STATUS_Udp,
}