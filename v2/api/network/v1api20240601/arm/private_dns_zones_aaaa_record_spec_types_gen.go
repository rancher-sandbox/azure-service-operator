// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PrivateDnsZonesAAAARecord_Spec struct {
	// Etag: The ETag of the record set.
	Etag *string `json:"etag,omitempty"`
	Name string  `json:"name,omitempty"`

	// Properties: The properties of the record set.
	Properties *RecordSetProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PrivateDnsZonesAAAARecord_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-06-01"
func (record PrivateDnsZonesAAAARecord_Spec) GetAPIVersion() string {
	return "2024-06-01"
}

// GetName returns the Name of the resource
func (record *PrivateDnsZonesAAAARecord_Spec) GetName() string {
	return record.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/AAAA"
func (record *PrivateDnsZonesAAAARecord_Spec) GetType() string {
	return "Microsoft.Network/privateDnsZones/AAAA"
}

// Represents the properties of the records in the record set.
type RecordSetProperties struct {
	// ARecords: The list of A records in the record set.
	ARecords []ARecord `json:"aRecords,omitempty"`

	// AaaaRecords: The list of AAAA records in the record set.
	AaaaRecords []AaaaRecord `json:"aaaaRecords,omitempty"`

	// CnameRecord: The CNAME record in the record set.
	CnameRecord *CnameRecord `json:"cnameRecord,omitempty"`

	// Metadata: The metadata attached to the record set.
	Metadata map[string]string `json:"metadata,omitempty"`

	// MxRecords: The list of MX records in the record set.
	MxRecords []MxRecord `json:"mxRecords,omitempty"`

	// PtrRecords: The list of PTR records in the record set.
	PtrRecords []PtrRecord `json:"ptrRecords,omitempty"`

	// SoaRecord: The SOA record in the record set.
	SoaRecord *SoaRecord `json:"soaRecord,omitempty"`

	// SrvRecords: The list of SRV records in the record set.
	SrvRecords []SrvRecord `json:"srvRecords,omitempty"`

	// Ttl: The TTL (time-to-live) of the records in the record set.
	Ttl *int `json:"ttl,omitempty"`

	// TxtRecords: The list of TXT records in the record set.
	TxtRecords []TxtRecord `json:"txtRecords,omitempty"`
}

// An AAAA record.
type AaaaRecord struct {
	// Ipv6Address: The IPv6 address of this AAAA record.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// An A record.
type ARecord struct {
	// Ipv4Address: The IPv4 address of this A record.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// A CNAME record.
type CnameRecord struct {
	// Cname: The canonical name for this CNAME record.
	Cname *string `json:"cname,omitempty"`
}

// An MX record.
type MxRecord struct {
	// Exchange: The domain name of the mail host for this MX record.
	Exchange *string `json:"exchange,omitempty"`

	// Preference: The preference value for this MX record.
	Preference *int `json:"preference,omitempty"`
}

// A PTR record.
type PtrRecord struct {
	// Ptrdname: The PTR target domain name for this PTR record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// An SOA record.
type SoaRecord struct {
	// Email: The email contact for this SOA record.
	Email *string `json:"email,omitempty"`

	// ExpireTime: The expire time for this SOA record.
	ExpireTime *int `json:"expireTime,omitempty"`

	// Host: The domain name of the authoritative name server for this SOA record.
	Host *string `json:"host,omitempty"`

	// MinimumTtl: The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTtl *int `json:"minimumTtl,omitempty"`

	// RefreshTime: The refresh value for this SOA record.
	RefreshTime *int `json:"refreshTime,omitempty"`

	// RetryTime: The retry time for this SOA record.
	RetryTime *int `json:"retryTime,omitempty"`

	// SerialNumber: The serial number for this SOA record.
	SerialNumber *int `json:"serialNumber,omitempty"`
}

// An SRV record.
type SrvRecord struct {
	// Port: The port value for this SRV record.
	Port *int `json:"port,omitempty"`

	// Priority: The priority value for this SRV record.
	Priority *int `json:"priority,omitempty"`

	// Target: The target domain name for this SRV record.
	Target *string `json:"target,omitempty"`

	// Weight: The weight value for this SRV record.
	Weight *int `json:"weight,omitempty"`
}

// A TXT record.
type TxtRecord struct {
	// Value: The text value of this TXT record.
	Value []string `json:"value,omitempty"`
}