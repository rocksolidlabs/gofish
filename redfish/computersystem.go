//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/rocksolidlabs/gofish/common"
)

// BootOrderTypes is the choice of boot order property to use when controller
// the persistent boot order for this computer system.
type BootOrderTypes string

const (

	// BootOrderBootOrderTypes indicates an ordered array of persistent boot
	// order for this system.
	BootOrderBootOrderTypes BootOrderTypes = "BootOrder"
	// AliasBootOrderBootOrderTypes  indicates an ordered array of aliased
	// persistent boot order for this system.
	AliasBootOrderBootOrderTypes BootOrderTypes = "AliasBootOrder"
)

// BootSourceOverrideEnabled describes the state of the Boot Source Override feature.
type BootSourceOverrideEnabled string

const (

	// DisabledBootSourceOverrideEnabled the system will boot normally.
	DisabledBootSourceOverrideEnabled BootSourceOverrideEnabled = "Disabled"
	// OnceBootSourceOverrideEnabled on its next boot cycle, the system will
	// boot (one time) to the Boot Source Override Target. The value of
	// BootSourceOverrideEnabled is then reset back to Disabled.
	OnceBootSourceOverrideEnabled BootSourceOverrideEnabled = "Once"
	// ContinuousBootSourceOverrideEnabled the system will boot to the target
	// specified in the BootSourceOverrideTarget until this property is set
	// to Disabled.
	ContinuousBootSourceOverrideEnabled BootSourceOverrideEnabled = "Continuous"
)

// BootSourceOverrideMode is the BIOS mode (Legacy or UEFI) to be used.
type BootSourceOverrideMode string

const (

	// LegacyBootSourceOverrideMode the system will boot in non-UEFI boot
	// mode to the Boot Source Override Target.
	LegacyBootSourceOverrideMode BootSourceOverrideMode = "Legacy"
	// UEFIBootSourceOverrideMode the system will boot in UEFI boot mode to
	// the Boot Source Override Target.
	UEFIBootSourceOverrideMode BootSourceOverrideMode = "UEFI"
)

// HostingRole specifies different features that the hosting ComputerSystem supports.
type HostingRole string

const (

	// ApplicationServerHostingRole the system hosts functionality that
	// supports general purpose applications.
	ApplicationServerHostingRole HostingRole = "ApplicationServer"
	// StorageServerHostingRole the system hosts functionality that supports
	// the system acting as a storage server.
	StorageServerHostingRole HostingRole = "StorageServer"
	// SwitchHostingRole the system hosts functionality that supports the
	// system acting as a switch.
	SwitchHostingRole HostingRole = "Switch"
)

// InterfaceType is the Trusted Platform Module type.
type InterfaceType string

const (

	// TPM1_2InterfaceType Trusted Platform Module (TPM) 1.2.
	TPM1_2InterfaceType InterfaceType = "TPM1_2"
	// TPM2_0InterfaceType Trusted Platform Module (TPM) 2.0.
	TPM2_0InterfaceType InterfaceType = "TPM2_0"
	// TCM1_0InterfaceType Trusted Cryptography Module (TCM) 1.0.
	TCM1_0InterfaceType InterfaceType = "TCM1_0"
)

// InterfaceTypeSelection specify the method for switching the TrustedModule
// InterfaceType, for instance between TPM1_2 and TPM2_0, if supported.
type InterfaceTypeSelection string

const (

	// NoneInterfaceTypeSelection the TrustedModule does not support
	// switching the InterfaceType.
	NoneInterfaceTypeSelection InterfaceTypeSelection = "None"
	// FirmwareUpdateInterfaceTypeSelection the TrustedModule supports
	// switching InterfaceType via a firmware update.
	FirmwareUpdateInterfaceTypeSelection InterfaceTypeSelection = "FirmwareUpdate"
	// BiosSettingInterfaceTypeSelection the TrustedModule supports switching
	// InterfaceType via platform software, such as a BIOS configuration
	// Attribute.
	BiosSettingInterfaceTypeSelection InterfaceTypeSelection = "BiosSetting"
	// OemMethodInterfaceTypeSelection the TrustedModule supports switching
	// InterfaceType via an OEM proprietary mechanism.
	OemMethodInterfaceTypeSelection InterfaceTypeSelection = "OemMethod"
)

// MemoryMirroring indicates the memory mirroring setting
type MemoryMirroring string

const (

	// SystemMemoryMirroring the system supports DIMM mirroring at the System
	// level. Individual DIMMs are not paired for mirroring in this mode.
	SystemMemoryMirroring MemoryMirroring = "System"
	// DIMMMemoryMirroring the system supports DIMM mirroring at the DIMM
	// level. Individual DIMMs can be mirrored.
	DIMMMemoryMirroring MemoryMirroring = "DIMM"
	// HybridMemoryMirroring the system supports a hybrid mirroring at the
	// system and DIMM levels. Individual DIMMs can be mirrored.
	HybridMemoryMirroring MemoryMirroring = "Hybrid"
	// NoneMemoryMirroring the system does not support DIMM mirroring.
	NoneMemoryMirroring MemoryMirroring = "None"
)

// PowerRestorePolicyTypes specifies the choice of power state for the system
// when power is applied.
type PowerRestorePolicyTypes string

const (

	// AlwaysOnPowerRestorePolicyTypes the system will always power on when
	// power is applied.
	AlwaysOnPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOn"
	// AlwaysOffPowerRestorePolicyTypes the system will always remain powered
	// off when power is applied.
	AlwaysOffPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOff"
	// LastStatePowerRestorePolicyTypes the system will return to its last
	// power state (on or off) when power is applied.
	LastStatePowerRestorePolicyTypes PowerRestorePolicyTypes = "LastState"
)

// PowerState is the power state of the system.
type PowerState string

const (

	// OnPowerState the system is powered on.
	OnPowerState PowerState = "On"
	// OffPowerState the system is powered off, although some components may
	// continue to have AUX power such as management controller.
	OffPowerState PowerState = "Off"
	// PoweringOnPowerState A temporary state between Off and On. This
	// temporary state can be very short.
	PoweringOnPowerState PowerState = "PoweringOn"
	// PoweringOffPowerState A temporary state between On and Off. The power
	// off action can take time while the OS is in the shutdown process.
	PoweringOffPowerState PowerState = "PoweringOff"
)

// SystemType is the type of system.
type SystemType string

const (

	// PhysicalSystemType is a System Type of Physical is typically used when
	// representing the hardware aspects of a system such as is done by a
	// management controller.
	PhysicalSystemType SystemType = "Physical"
	// VirtualSystemType is a System Type of Virtual is typically used when
	// representing a system that is actually a virtual machine instance.
	VirtualSystemType SystemType = "Virtual"
	// OSSystemType is a System Type of OS is typically used when representing
	// an OS or hypervisor view of the system.
	OSSystemType SystemType = "OS"
	// PhysicallyPartitionedSystemType is a System Type of PhysicallyPartition is
	// typically used when representing a single system constructed from
	// one or more physical systems via a firmware or hardware-based service.
	PhysicallyPartitionedSystemType SystemType = "PhysicallyPartitioned"
	// VirtuallyPartitionedSystemType is a System Type of VirtuallyPartition is
	// typically used when representing a single system constructed from
	// one or more virtual systems via a software-based service.
	VirtuallyPartitionedSystemType SystemType = "VirtuallyPartitioned"
	// ComposedSystemType is a System Type of Composed is typically used when
	// representing a single system constructed from disaggregated resource
	// via the Redfish Composition service.
	ComposedSystemType SystemType = "Composed"
)

// WatchdogTimeoutActions specifies the choice of action to take when the Host
// Watchdog Timer reaches its timeout value.
type WatchdogTimeoutActions string

const (

	// NoneWatchdogTimeoutActions means no action taken.
	NoneWatchdogTimeoutActions WatchdogTimeoutActions = "None"
	// ResetSystemWatchdogTimeoutActions means reset the system.
	ResetSystemWatchdogTimeoutActions WatchdogTimeoutActions = "ResetSystem"
	// PowerCycleWatchdogTimeoutActions means power cycle the system.
	PowerCycleWatchdogTimeoutActions WatchdogTimeoutActions = "PowerCycle"
	// PowerDownWatchdogTimeoutActions means power down the system.
	PowerDownWatchdogTimeoutActions WatchdogTimeoutActions = "PowerDown"
	// OEMWatchdogTimeoutActions means perform an OEM-defined action.
	OEMWatchdogTimeoutActions WatchdogTimeoutActions = "OEM"
)

// WatchdogWarningActions specifies the choice of action to take when the Host
// Watchdog Timer is close (typically 3-10 seconds) to reaching its timeout value.
type WatchdogWarningActions string

const (

	// NoneWatchdogWarningActions means no action taken.
	NoneWatchdogWarningActions WatchdogWarningActions = "None"
	// DiagnosticInterruptWatchdogWarningActions means raise a (typically non-
	// maskable) Diagnostic Interrupt.
	DiagnosticInterruptWatchdogWarningActions WatchdogWarningActions = "DiagnosticInterrupt"
	// SMIWatchdogWarningActions means raise a Systems Management Interrupt (SMI).
	SMIWatchdogWarningActions WatchdogWarningActions = "SMI"
	// MessagingInterruptWatchdogWarningActions means raise a legacy IPMI messaging
	// interrupt.
	MessagingInterruptWatchdogWarningActions WatchdogWarningActions = "MessagingInterrupt"
	// SCIWatchdogWarningActions means raise an interrupt using the ACPI System
	// Control Interrupt (SCI).
	SCIWatchdogWarningActions WatchdogWarningActions = "SCI"
	// OEMWatchdogWarningActions means perform an OEM-defined action.
	OEMWatchdogWarningActions WatchdogWarningActions = "OEM"
)

// Boot contains properties which describe boot information for a system.
type Boot struct {
	// AliasBootOrder shall be an ordered array
	// of boot source aliases (of type BootSource) representing the
	// persistent Boot Order of this computer system.
	AliasBootOrder []string
	// BootNext shall be the
	// BootOptionReference of the UEFI Boot Option for one time boot, as
	// defined by the UEFI Specification. The valid values for this property
	// are specified in the values of the BootOrder array.
	// BootSourceOverrideEnabled = Continuous is not supported for UEFI
	// BootNext as this setting is defined in UEFI as a one-time boot only.
	BootNext string
	// BootOptions shall be a link to a
	// collection of type BootOptionCollection.
	bootOptions string
	// BootOrder shall be an ordered array of
	// BootOptionReference strings representing the persistent Boot Order of
	// this computer system. For UEFI systems, this is the UEFI BootOrder as
	// defined by the UEFI Specification.
	BootOrder []string
	// BootOrderPropertySelection shall
	// indicate which boot order property the system uses when specifying the
	// persistent boot order.
	BootOrderPropertySelection string
	// BootSourceOverrideEnabled shall be Once
	// if this is a one time boot override and Continuous if this selection
	// should remain active until cancelled. If the property value is set to
	// Once, the value will be reset back to Disabled after the
	// BootSourceOverrideTarget actions have been completed. Changes to this
	// property do not alter the BIOS persistent boot order configuration.
	BootSourceOverrideEnabled string
	// BootSourceOverrideMode shall be Legacy
	// for non-UEFI BIOS boot or UEFI for UEFI boot from boot source
	// specified in BootSourceOverrideTarget property.
	BootSourceOverrideMode string
	// BootSourceOverrideTarget shall contain
	// the source to boot the system from, overriding the normal boot order.
	// The valid values for this property are specified through the
	// Redfish.AllowableValues annotation. Pxe indicates to PXE boot from the
	// primary NIC; Floppy, Cd, Usb, Hdd indicates to boot from their devices
	// respectively. BiosSetup indicates to boot into the native BIOS screen
	// setup. Utilities and Diags indicate to boot from the local utilities
	// or diags partitions. UefiTarget indicates to boot from the UEFI device
	// path found in UefiTargetBootSourceOverride. UefiBootNext indicates to
	// boot from the UEFI BootOptionReference found in BootNext. Changes to
	// this property do not alter the BIOS persistent boot order
	// configuration.
	BootSourceOverrideTarget string
	// UefiTargetBootSourceOverride shall be
	// the UEFI device path of the override boot target. The valid values for
	// this property are specified through the Redfish.AllowableValues
	// annotation. BootSourceOverrideEnabled = Continuous is not supported
	// for UEFI Boot Source Override as this setting is defined in UEFI as a
	// one time boot only. Changes to this property do not alter the BIOS
	// persistent boot order configuration.
	UefiTargetBootSourceOverride string
}

// UnmarshalJSON unmarshals a Boot object from the raw JSON.
func (boot *Boot) UnmarshalJSON(b []byte) error {
	type temp Boot
	var t struct {
		temp
		BootOptions common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boot = Boot(t.temp)

	// Extract the links to other entities for later
	boot.bootOptions = string(t.BootOptions)

	return nil
}

// ResetType describe the type off reset to be issue by the resource
type ResetType string

const (
	// OnResetType shall be used to power on the machine
	OnResetType ResetType = "On"
	// ForceOffResetType shall be used to power off the machine without wait the OS to shutdown
	ForceOffResetType ResetType = "ForceOff"
	// ForceRestartResetType shall be used to restart the machine without wait the OS to shutdown
	ForceRestartResetType ResetType = "ForceRestart"
	// GracefulShutdownResetType shall be used to restart the machine waiting the OS shutdown gracefully
	GracefulShutdownResetType ResetType = "GracefulShutdown"
	// PushPowerButtonResetType shall be used to emulate pushing the power button
	PushPowerButtonResetType ResetType = "PushPowerButton"
	// NmiResetType shall be used to trigger a crash/core dump file
	NmiResetType ResetType = "Nmi"
)

// CSActions shall contain the available actions for this resource
type CSActions struct {
	// ComputerSystemReset shall perform a reset of the ComputerSystem. For
	// systems which implement ACPI Power Button functionality, the
	// PushPowerButton value shall perform or emulate an ACPI Power Button push.
	// The ForceOff value shall remove power from the system or perform an ACPI
	// Power Button Override (commonly known as a 4-second hold of the Power
	// Button). The ForceRestart value shall perform a ForceOff action followed
	// by a On action.
	ComputerSystemReset struct {
		ResetType []ResetType `json:"ResetType@Redfish.AllowableValues"`
		Target    string
	} `json:"#ComputerSystem.Reset"`
}

// ComputerSystem is used to represent resources that represent a
// computing system in the Redfish specification.
type ComputerSystem struct {
	common.Entity

	// ODataContext is the @odata.context
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the @odata.etag
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the @odata.id
	ODataID string `json:"@odata.id"`
	// ODataType is the @odata.type
	ODataType string `json:"@odata.type"`

	// Actions type shall contain the available actions for this resource.
	Actions CSActions
	// AssetTag shall contain the value of the asset tag of the system.
	AssetTag string
	// bios shall be a link to a resource of
	// type Bios that lists the Bios settings for this system.
	bios string
	// BIOSVersion shall be the version string
	// of the currently installed and running BIOS (for x86 systems). For
	// other systems, the value may contain a version string representing the
	// primary system firmware.
	BIOSVersion string `json:"BiosVersion"`
	// Boot describes boot information for the current resource.
	Boot Boot
	// Description is the resource description.
	Description string
	// EthernetInterfaces shall be a link to a
	// collection of type EthernetInterfaceCollection.
	ethernetInterfaces string
	// HostName shall be the host name for this
	// system, as reported by the operating system or hypervisor. This value
	// is typically provided to the Manager by a service running in the host
	// operating system.
	HostName string
	// HostWatchdogTimer shall contain properties which
	// describe the host watchdog timer functionality for this
	// ComputerSystem.
	HostWatchdogTimer WatchdogTimer
	// HostedServices shall describe services supported by this computer system.
	HostedServices string
	// HostingRoles shall be the hosting roles supported by this computer system.
	HostingRoles []string

	// IndicatorLED shall contain the indicator
	// light state for the indicator light associated with this system.
	IndicatorLED common.IndicatorLED
	// logServices shall be a link to a collection of type LogServiceCollection.
	logServices string
	// Manufacturer shall contain a value that represents the manufacturer of the system.
	Manufacturer string
	// Memory shall be a link to a collection of type MemoryCollection.
	memory string
	// memoryDomains shall be a link to a collection of type MemoryDomainCollection.
	memoryDomains string
	// MemorySummary is This object shall contain properties which describe
	// the central memory for the current resource.
	MemorySummary MemorySummary
	// Model shall contain the information
	// about how the manufacturer references this system. This is typically
	// the product name, without the manufacturer name.
	Model string
	// Name is the resource name.
	Name string
	// networkInterfaces shall be a link to a collection of type
	// NetworkInterfaceCollection.
	networkInterfaces string
	// PCIeDevices shall be an array of references of type PCIeDevice.
	pcieDevices []string
	// PCIeDevicesCount is the number of PCIeDevices.
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// PCIeFunctions shall be an array of references of type PCIeFunction.
	pcieFunctions []string
	// PCIeFunctionsCount is the number of PCIeFunctions.
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// PartNumber shall contain the part number
	// for the system as defined by the manufacturer.
	PartNumber string
	// PowerRestorePolicy is the desired
	// PowerState of the system when power is applied to the system. A value
	// of 'LastState' shall return the system to the PowerState it was in
	// when power was lost.
	PowerRestorePolicy PowerState
	// PowerState shall contain the power state of the system.
	PowerState PowerState
	// ProcessorSummary shall contain properties which
	// describe the central processors for the current resource.
	ProcessorSummary ProcessorSummary
	// Processors shall be a link to a collection of type ProcessorCollection.
	processors string
	// Redundancy references a redundancy
	// entity that specifies a kind and level of redundancy and a collection
	// (RedundancySet) of other ComputerSystems that provide the specified
	// redundancy to this ComputerSystem.
	Redundancy string
	// RedundancyCount is the number of Reduncy objects.
	RedundancyCount string `json:"Redundancy@odata.count"`
	// SKU shall contain the Stock Keeping Unit (SKU) for the system.
	SKU string
	// secureBoot shall be a link to a resource of type SecureBoot.
	secureBoot string
	// SerialNumber shall contain the serial number for the system.
	SerialNumber string
	// SimpleStorage shall be a link to a collection of type SimpleStorageCollection.
	simpleStorage string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// storage shall be a link to a collection
	// of type StorageCollection.
	storage string
	// SubModel shall contain the information
	// about the sub-model (or config) of the system. This shall not include
	// the model/product name or the manufacturer name.
	SubModel string
	// SystemType indicates the kind of system that this resource represents.
	SystemType SystemType
	// TrustedModules shall contain an array of objects with
	// properties which describe the trusted modules for the current resource.
	TrustedModules []TrustedModules
	// UUID is used to contain a universal unique identifier number for the
	// system. RFC4122 describes methods that can be used to create the
	// value. The value should be considered to be opaque. Client software
	// should only treat the overall value as a universally unique identifier
	// and should not interpret any sub-fields within the UUID. If the system
	// supports SMBIOS, the value of the property should be formed by
	// following the SMBIOS 2.6+ recommendation for converting the SMBIOS
	// 16-byte UUID structure into the redfish canonical xxxxxxxx-xxxx-xxxx-
	// xxxx-xxxxxxxxxxxx string format so that the property value matches the
	// byte order presented by current OS APIs such as WMI and dmidecode.
	UUID string
	// Chassis is an array of references to the chassis in which this system is contained.
	chassis []string
}

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
func (computersystem *ComputerSystem) UnmarshalJSON(b []byte) error {
	type temp ComputerSystem
	var t struct {
		temp
		Bios               common.Link
		Processors         common.Link
		Memory             common.Link
		EthernetInterfaces common.Link
		SimpleStorage      common.Link
		SecureBoot         common.Link
		Storage            common.Link
		NetworkInterfaces  common.Link
		LogServices        common.Link
		MemoryDomains      common.Link
		PCIeDevices        common.Links
		PCIeFunctions      common.Links
		Links              CSLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*computersystem = ComputerSystem(t.temp)

	// Extract the links to other entities for later
	computersystem.bios = string(t.Bios)
	computersystem.processors = string(t.Processors)
	computersystem.memory = string(t.Memory)
	computersystem.ethernetInterfaces = string(t.EthernetInterfaces)
	computersystem.simpleStorage = string(t.SimpleStorage)
	computersystem.networkInterfaces = string(t.NetworkInterfaces)
	computersystem.secureBoot = string(t.SecureBoot)
	computersystem.storage = string(t.Storage)
	computersystem.logServices = string(t.LogServices)
	computersystem.memoryDomains = string(t.MemoryDomains)
	computersystem.pcieDevices = t.PCIeDevices.ToStrings()
	computersystem.pcieFunctions = t.PCIeFunctions.ToStrings()
	computersystem.chassis = t.Links.Chassis.ToStrings()

	return nil
}

// GetComputerSystem will get a ComputerSystem instance from the service.
func GetComputerSystem(c common.Client, uri string) (*ComputerSystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var computersystem ComputerSystem
	err = json.NewDecoder(resp.Body).Decode(&computersystem)
	if err != nil {
		return nil, err
	}

	computersystem.SetClient(c)
	return &computersystem, nil
}

// ListReferencedComputerSystems gets the collection of ComputerSystem from
// a provided reference.
func ListReferencedComputerSystems(c common.Client, link string) ([]*ComputerSystem, error) {
	var result []*ComputerSystem
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, computersystemLink := range links.ItemLinks {
		computersystem, err := GetComputerSystem(c, computersystemLink)
		if err != nil {
			return result, err
		}
		result = append(result, computersystem)
	}

	return result, nil
}

// Bios gets the Bios information for this ComputerSystem.
func (computersystem *ComputerSystem) Bios() (*Bios, error) {
	if computersystem.bios == "" {
		return nil, nil
	}

	return GetBios(computersystem.Client, computersystem.bios)
}

// EthernetInterfaces get this system's ethernet interfaces.
func (computersystem *ComputerSystem) EthernetInterfaces() ([]*EthernetInterface, error) {
	return ListReferencedEthernetInterfaces(computersystem.Client, computersystem.ethernetInterfaces)
}

// LogServices get this system's log services.
func (computersystem *ComputerSystem) LogServices() ([]*LogService, error) {
	return ListReferencedLogServices(computersystem.Client, computersystem.logServices)
}

// Memory gets this system's memory.
func (computersystem *ComputerSystem) Memory() ([]*Memory, error) {
	return ListReferencedMemorys(computersystem.Client, computersystem.memory)
}

// MemoryDomains gets this system's memory domains.
func (computersystem *ComputerSystem) MemoryDomains() ([]*MemoryDomain, error) {
	return ListReferencedMemoryDomains(computersystem.Client, computersystem.memoryDomains)
}

// NetworkInterfaces returns a collection of network interfaces in this system.
func (computersystem *ComputerSystem) NetworkInterfaces() ([]*NetworkInterface, error) {
	return ListReferencedNetworkInterfaces(computersystem.Client, computersystem.networkInterfaces)
}

// PCIeDevices gets all PCIeDevices for this system.
func (computersystem *ComputerSystem) PCIeDevices() ([]*PCIeDevice, error) {
	var result []*PCIeDevice
	for _, pciedeviceLink := range computersystem.pcieDevices {
		pciedevice, err := GetPCIeDevice(computersystem.Client, pciedeviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, pciedevice)
	}
	return result, nil
}

// PCIeFunctions gets all PCIeFunctions for this system.
func (computersystem *ComputerSystem) PCIeFunctions() ([]*PCIeFunction, error) {
	var result []*PCIeFunction
	for _, pciefunctionLink := range computersystem.pcieFunctions {
		pciefunction, err := GetPCIeFunction(computersystem.Client, pciefunctionLink)
		if err != nil {
			return result, err
		}
		result = append(result, pciefunction)
	}
	return result, nil
}

// Processors returns a collection of processors from this system
func (computersystem *ComputerSystem) Processors() ([]*Processor, error) {
	return ListReferencedProcessors(computersystem.Client, computersystem.processors)
}

// SecureBoot gets the secure boot information for the system.
func (computersystem *ComputerSystem) SecureBoot() (*SecureBoot, error) {
	if computersystem.secureBoot == "" {
		return nil, nil
	}

	return GetSecureBoot(computersystem.Client, computersystem.secureBoot)
}

// SimpleStorages gets all simple storage services of this system.
func (computersystem *ComputerSystem) SimpleStorages() ([]*SimpleStorage, error) {
	return ListReferencedSimpleStorages(computersystem.Client, computersystem.simpleStorage)
}

// Storage gets the storage associated with this system.
func (computersystem *ComputerSystem) Storage() ([]*Storage, error) {
	return ListReferencedStorages(computersystem.Client, computersystem.storage)
}

// CSLinks are references to resources that are related to, but not contained
// by (subordinate to), this resource.
type CSLinks struct {
	// Chassis shall reference a resource of
	// type Chassis that represents the physical container associated with
	// this resource.
	Chassis common.Links
	// ChassisCount is
	ChassisCount int `json:"Chassis@odata.count"`
	// ConsumingComputerSystems shall be an array of references
	// to ComputerSystems that are realized, in whole or in part, from this
	// ComputerSystem.
	ConsumingComputerSystems common.Links
	// ConsumingComputerSystemsCount is
	ConsumingComputerSystemsCount int `json:"ConsumingComputerSystems@odata.count"`
	// CooledBy shall be an array of IDs
	// containing pointers consistent with JSON pointer syntax to the
	// resource that powers this computer system.
	CooledBy common.Links
	// CooledByCount is
	CooledByCount int `json:"CooledBy@odata.count"`
	// Endpoints shall be a reference to the
	// resources that this system is associated with and shall reference a
	// resource of type Endpoint.
	Endpoints common.Links
	// EndpointsCount is
	EndpointsCount int `json:"Endpoints@odata.count"`
	// ManagedBy shall reference a resource of
	// type manager that represents the resource with management
	// responsibility for this resource.
	ManagedBy common.Links
	// ManagedByCount is
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// PoweredBy shall be an array of IDs
	// containing pointers consistent with JSON pointer syntax to the
	// resource that powers this computer system.
	PoweredBy common.Links
	// PoweredByCount is the number of PoweredBy objects.
	PoweredByCount int `json:"PoweredBy@odata.count"`
	// ResourceBlocks is used in this Computer System.
	ResourceBlocks common.Links
	// ResourceBlocksCount is the nunmber of ResourceBlocks.
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
	// SupplyingComputerSystems shall be an array of references
	// to ComputerSystems that contribute, in whole or in part, to the
	// implementation of this ComputerSystem.
	SupplyingComputerSystems common.Links
	// SupplyingComputerSystemsCount is the number of SupplyingComputerSystems.
	SupplyingComputerSystemsCount int `json:"SupplyingComputerSystems@odata.count"`
}

// MemorySummary contains properties which describe the central memory for a system.
type MemorySummary struct {
	// MemoryMirroring is the ability and type of memory mirring supported by this system.
	MemoryMirroring MemoryMirroring
	// Status is the status or health properties of the resource.
	Status common.Status
	// TotalSystemMemoryGiB is the amount of configured system general purpose
	// volatile (RAM) memory as measured in gibibytes.
	TotalSystemMemoryGiB float32
	// TotalSystemPersistentMemoryGiB is the total amount of configured
	// persistent memory available to the system as measured in gibibytes.
	TotalSystemPersistentMemoryGiB float32
}

// ProcessorSummary is This type shall contain properties which describe
// the central processors for a system.
type ProcessorSummary struct {
	// Count is the number of physical central processors in the system.
	Count int
	// LogicalProcessorCount is the number of logical central processors in the system.
	LogicalProcessorCount int
	// Model is the processor model for the central processors in the system,
	// per the description in the Processor Information - Processor Family
	// section of the SMBIOS Specification DSP0134 2.8 or later.
	Model string
	// Status is any status or health properties of the resource.
	Status common.Status
}

// TrustedModules is This type shall describe a truted module for a system.
type TrustedModules struct {
	// FirmwareVersion is the firmware version as
	// defined by the manufacturer for the Trusted Module.
	FirmwareVersion string
	// FirmwareVersion2 is the 2nd firmware
	// version, if applicable, as defined by the manufacturer for the Trusted
	// Module.
	FirmwareVersion2 string
	// InterfaceType is the interface type of the installed Trusted Module.
	InterfaceType InterfaceType
	// InterfaceTypeSelection is the Interface
	// Type Selection method (for example to switch between TPM1_2 and
	// TPM2_0) that is supported by this TrustedModule.
	InterfaceTypeSelection InterfaceTypeSelection
	// Status is any status or health properties
	// of the resource.
	Status common.Status
}

// WatchdogTimer contains properties which describe the
// host watchdog timer functionality for this ComputerSystem.
type WatchdogTimer struct {
	// FunctionEnabled shall indicate whether
	// the host watchdog timer functionality has been enabled or not. This
	// property indicates only that the functionality is enabled or disabled
	// by the user, and updates to this property shall not initiate a
	// watchdog timer countdown.
	FunctionEnabled bool
	// Status is any status or health properties of the resource.
	Status struct {
		State common.State
	}
	// TimeoutAction is the action to perform
	// upon the  expiration of the Watchdog Timer.
	TimeoutAction string
	// WarningAction is the action to perform
	// prior to the expiration of the Watchdog Timer. This action typically
	// occurs 3-10 seconds prior to the timeout value, but the exact timing
	// is dependent on the implementation.
	WarningAction string
}
