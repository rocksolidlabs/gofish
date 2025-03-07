//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/rocksolidlabs/gofish/common"
)

// ProvisioningPolicy is used to specify space provisioning policy.
type ProvisioningPolicy string

const (
	// FixedProvisioningPolicy shall be fully allocated.
	FixedProvisioningPolicy ProvisioningPolicy = "Fixed"
	// ThinProvisioningPolicy specifies storage may be over allocated.
	ThinProvisioningPolicy ProvisioningPolicy = "Thin"
)

// StorageAccessCapability is used to describe abilities to read or write
// storage.
type StorageAccessCapability string

const (
	// ReadStorageAccessCapability shall indicate that the storage may be
	// read.
	ReadStorageAccessCapability StorageAccessCapability = "Read"
	// WriteStorageAccessCapability shall indicate that the storage may be
	// written multiple times.
	WriteStorageAccessCapability StorageAccessCapability = "Write"
	// WriteOnceStorageAccessCapability shall indicate that the storage may
	// be written only once.
	WriteOnceStorageAccessCapability StorageAccessCapability = "WriteOnce"
	// AppendStorageAccessCapability shall indicate that the storage may be
	// written only to append.
	AppendStorageAccessCapability StorageAccessCapability = "Append"
	// StreamingStorageAccessCapability shall indicate that the storage may
	// be read sequentially.
	StreamingStorageAccessCapability StorageAccessCapability = "Streaming"
	// ExecuteStorageAccessCapability shall indicate that Execute access is
	// allowed by the file share.
	ExecuteStorageAccessCapability StorageAccessCapability = "Execute"
)

// DataStorageLoSCapabilities describes capabilities of the system to
// support various data storage service options.
type DataStorageLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// MaximumRecoverableCapacitySourceCount is the maximum number of capacity
	// source resources that can be supported for the purpose of recovery when
	// in the event that an equivalent capacity source resource fails.
	MaximumRecoverableCapacitySourceCount int
	// SupportedAccessCapabilities specifies a storage access capabilities.
	SupportedAccessCapabilities []StorageAccessCapability
	// SupportedLinesOfService shall contain known and supported DataStorageLinesOfService.
	SupportedLinesOfService []DataStorageLineOfService
	// SupportedLinesOfServiceCount is
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedProvisioningPolicies specifies supported storage allocation policies.
	SupportedProvisioningPolicies []ProvisioningPolicy
	// SupportedRecoveryTimeObjectives specifies supported expectations for time
	// to access the primary store after recovery.
	SupportedRecoveryTimeObjectives []RecoveryAccessScope
	// SupportsSpaceEfficiency specifies whether storage compression or
	// deduplication is supported. The default value for this property is false.
	SupportsSpaceEfficiency bool
}

// GetDataStorageLoSCapabilities will get a DataStorageLoSCapabilities instance from the service.
func GetDataStorageLoSCapabilities(c common.Client, uri string) (*DataStorageLoSCapabilities, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var datastorageloscapabilities DataStorageLoSCapabilities
	err = json.NewDecoder(resp.Body).Decode(&datastorageloscapabilities)
	if err != nil {
		return nil, err
	}

	datastorageloscapabilities.SetClient(c)
	return &datastorageloscapabilities, nil
}

// ListReferencedDataStorageLoSCapabilities gets the collection of DataStorageLoSCapabilities from
// a provided reference.
func ListReferencedDataStorageLoSCapabilities(c common.Client, link string) ([]*DataStorageLoSCapabilities, error) {
	var result []*DataStorageLoSCapabilities
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, datastorageloscapabilitiesLink := range links.ItemLinks {
		datastorageloscapabilities, err := GetDataStorageLoSCapabilities(c, datastorageloscapabilitiesLink)
		if err != nil {
			return result, err
		}
		result = append(result, datastorageloscapabilities)
	}

	return result, nil
}
