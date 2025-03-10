//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/rocksolidlabs/gofish/common"
)

var ioConnectivityLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#IOConnectivityLoSCapabilities.IOConnectivityLoSCapabilities",
		"@odata.type": "#IOConnectivityLoSCapabilities.v1_1_1.IOConnectivityLoSCapabilities",
		"@odata.id": "/redfish/v1/IOConnectivityLoSCapabilities",
		"Id": "IOConnectivityLoSCapabilities-1",
		"Name": "IOConnectivityLoSCapabilitiesOne",
		"Description": "IOConnectivityLoSCapabilities One",
		"MaxSupportedBytesPerSecond": 5000000000,
		"MaxSupportedIOPS": 1000000000,
		"SupportedAccessProtocols": [
			"FC",
			"FCP",
			"FCoE",
			"iSCSI"
		],
		"SupportedLinesOfService": [{
				"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
				"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
				"Id": "IOConnectivityLineOfService-1",
				"Name": "IOConnectivityLineOfServiceOne",
				"Description": "IOConnectivityLineOfService One",
				"AccessProtocols": [
					"FC",
					"FCP",
					"FCoE",
					"iSCSI"
				],
				"MaxBytesPerSecond": 5000000000,
				"MaxIOPS": 1000000000
			},
			{
				"@odata.context": "/redfish/v1/$metadata#IOConnectivityLineOfService.IOConnectivityLineOfService",
				"@odata.type": "#IOConnectivityLineOfService.v1_1_1.IOConnectivityLineOfService",
				"@odata.id": "/redfish/v1/IOConnectivityLineOfService",
				"Id": "IOConnectivityLineOfService-2",
				"Name": "IOConnectivityLineOfServiceTwo",
				"Description": "IOConnectivityLineOfService Two",
				"AccessProtocols": [
					"FC",
					"FCP",
					"FCoE"
				],
				"MaxBytesPerSecond": 5000000000,
				"MaxIOPS": 1000000000
			}
		]
	}`)

// TestIOConnectivityLoSCapabilities tests the parsing of IOConnectivityLoSCapabilities objects.
func TestIOConnectivityLoSCapabilities(t *testing.T) {
	var result IOConnectivityLoSCapabilities
	err := json.NewDecoder(ioConnectivityLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "IOConnectivityLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "IOConnectivityLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.MaxSupportedBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.MaxSupportedBytesPerSecond)
	}

	if result.MaxSupportedIOPS != 1000000000 {
		t.Errorf("MaxSupportedIOPS: %d", result.MaxSupportedIOPS)
	}

	if result.SupportedAccessProtocols[1] != common.FCPProtocol {
		t.Errorf("Invalid AccessProtocol: %s", result.SupportedAccessProtocols[1])
	}

	if result.SupportedLinesOfService[0].MaxBytesPerSecond != 5000000000 {
		t.Errorf("Invalid MaxSupportedBytesPerSecond: %d", result.SupportedLinesOfService[0].MaxBytesPerSecond)
	}
}
