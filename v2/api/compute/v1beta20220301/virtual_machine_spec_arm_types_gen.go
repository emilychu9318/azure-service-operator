// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualMachine_Spec. Use v1api20220301.VirtualMachine_Spec instead
type VirtualMachine_Spec_ARM struct {
	ExtendedLocation *ExtendedLocation_ARM         `json:"extendedLocation,omitempty"`
	Identity         *VirtualMachineIdentity_ARM   `json:"identity,omitempty"`
	Location         *string                       `json:"location,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Plan             *Plan_ARM                     `json:"plan,omitempty"`
	Properties       *VirtualMachineProperties_ARM `json:"properties,omitempty"`
	Tags             map[string]string             `json:"tags,omitempty"`
	Zones            []string                      `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualMachine_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (machine VirtualMachine_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (machine *VirtualMachine_Spec_ARM) GetName() string {
	return machine.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachines"
func (machine *VirtualMachine_Spec_ARM) GetType() string {
	return "Microsoft.Compute/virtualMachines"
}

// Deprecated version of VirtualMachineIdentity. Use v1api20220301.VirtualMachineIdentity instead
type VirtualMachineIdentity_ARM struct {
	Type *VirtualMachineIdentity_Type `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineProperties. Use v1api20220301.VirtualMachineProperties instead
type VirtualMachineProperties_ARM struct {
	AdditionalCapabilities  *AdditionalCapabilities_ARM     `json:"additionalCapabilities,omitempty"`
	ApplicationProfile      *ApplicationProfile_ARM         `json:"applicationProfile,omitempty"`
	AvailabilitySet         *SubResource_ARM                `json:"availabilitySet,omitempty"`
	BillingProfile          *BillingProfile_ARM             `json:"billingProfile,omitempty"`
	CapacityReservation     *CapacityReservationProfile_ARM `json:"capacityReservation,omitempty"`
	DiagnosticsProfile      *DiagnosticsProfile_ARM         `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy          *EvictionPolicy                 `json:"evictionPolicy,omitempty"`
	ExtensionsTimeBudget    *string                         `json:"extensionsTimeBudget,omitempty"`
	HardwareProfile         *HardwareProfile_ARM            `json:"hardwareProfile,omitempty"`
	Host                    *SubResource_ARM                `json:"host,omitempty"`
	HostGroup               *SubResource_ARM                `json:"hostGroup,omitempty"`
	LicenseType             *string                         `json:"licenseType,omitempty"`
	NetworkProfile          *NetworkProfile_ARM             `json:"networkProfile,omitempty"`
	OsProfile               *OSProfile_ARM                  `json:"osProfile,omitempty"`
	PlatformFaultDomain     *int                            `json:"platformFaultDomain,omitempty"`
	Priority                *Priority                       `json:"priority,omitempty"`
	ProximityPlacementGroup *SubResource_ARM                `json:"proximityPlacementGroup,omitempty"`
	ScheduledEventsProfile  *ScheduledEventsProfile_ARM     `json:"scheduledEventsProfile,omitempty"`
	SecurityProfile         *SecurityProfile_ARM            `json:"securityProfile,omitempty"`
	StorageProfile          *StorageProfile_ARM             `json:"storageProfile,omitempty"`
	UserData                *string                         `json:"userData,omitempty"`
	VirtualMachineScaleSet  *SubResource_ARM                `json:"virtualMachineScaleSet,omitempty"`
}

// Deprecated version of ApplicationProfile. Use v1api20220301.ApplicationProfile instead
type ApplicationProfile_ARM struct {
	GalleryApplications []VMGalleryApplication_ARM `json:"galleryApplications,omitempty"`
}

// Deprecated version of BillingProfile. Use v1api20220301.BillingProfile instead
type BillingProfile_ARM struct {
	MaxPrice *float64 `json:"maxPrice,omitempty"`
}

// Deprecated version of CapacityReservationProfile. Use v1api20220301.CapacityReservationProfile instead
type CapacityReservationProfile_ARM struct {
	CapacityReservationGroup *SubResource_ARM `json:"capacityReservationGroup,omitempty"`
}

// Deprecated version of DiagnosticsProfile. Use v1api20220301.DiagnosticsProfile instead
type DiagnosticsProfile_ARM struct {
	BootDiagnostics *BootDiagnostics_ARM `json:"bootDiagnostics,omitempty"`
}

// Deprecated version of HardwareProfile. Use v1api20220301.HardwareProfile instead
type HardwareProfile_ARM struct {
	VmSize           *HardwareProfile_VmSize `json:"vmSize,omitempty"`
	VmSizeProperties *VMSizeProperties_ARM   `json:"vmSizeProperties,omitempty"`
}

// Deprecated version of NetworkProfile. Use v1api20220301.NetworkProfile instead
type NetworkProfile_ARM struct {
	NetworkApiVersion              *NetworkProfile_NetworkApiVersion                 `json:"networkApiVersion,omitempty"`
	NetworkInterfaceConfigurations []VirtualMachineNetworkInterfaceConfiguration_ARM `json:"networkInterfaceConfigurations,omitempty"`
	NetworkInterfaces              []NetworkInterfaceReference_ARM                   `json:"networkInterfaces,omitempty"`
}

// Deprecated version of OSProfile. Use v1api20220301.OSProfile instead
type OSProfile_ARM struct {
	AdminPassword               *string                   `json:"adminPassword,omitempty"`
	AdminUsername               *string                   `json:"adminUsername,omitempty"`
	AllowExtensionOperations    *bool                     `json:"allowExtensionOperations,omitempty"`
	ComputerName                *string                   `json:"computerName,omitempty"`
	CustomData                  *string                   `json:"customData,omitempty"`
	LinuxConfiguration          *LinuxConfiguration_ARM   `json:"linuxConfiguration,omitempty"`
	RequireGuestProvisionSignal *bool                     `json:"requireGuestProvisionSignal,omitempty"`
	Secrets                     []VaultSecretGroup_ARM    `json:"secrets,omitempty"`
	WindowsConfiguration        *WindowsConfiguration_ARM `json:"windowsConfiguration,omitempty"`
}

// Deprecated version of ScheduledEventsProfile. Use v1api20220301.ScheduledEventsProfile instead
type ScheduledEventsProfile_ARM struct {
	TerminateNotificationProfile *TerminateNotificationProfile_ARM `json:"terminateNotificationProfile,omitempty"`
}

// Deprecated version of SecurityProfile. Use v1api20220301.SecurityProfile instead
type SecurityProfile_ARM struct {
	EncryptionAtHost *bool                         `json:"encryptionAtHost,omitempty"`
	SecurityType     *SecurityProfile_SecurityType `json:"securityType,omitempty"`
	UefiSettings     *UefiSettings_ARM             `json:"uefiSettings,omitempty"`
}

// Deprecated version of StorageProfile. Use v1api20220301.StorageProfile instead
type StorageProfile_ARM struct {
	DataDisks      []DataDisk_ARM      `json:"dataDisks,omitempty"`
	ImageReference *ImageReference_ARM `json:"imageReference,omitempty"`
	OsDisk         *OSDisk_ARM         `json:"osDisk,omitempty"`
}

// Deprecated version of VirtualMachineIdentity_Type. Use v1api20220301.VirtualMachineIdentity_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type VirtualMachineIdentity_Type string

const (
	VirtualMachineIdentity_Type_None                       = VirtualMachineIdentity_Type("None")
	VirtualMachineIdentity_Type_SystemAssigned             = VirtualMachineIdentity_Type("SystemAssigned")
	VirtualMachineIdentity_Type_SystemAssignedUserAssigned = VirtualMachineIdentity_Type("SystemAssigned, UserAssigned")
	VirtualMachineIdentity_Type_UserAssigned               = VirtualMachineIdentity_Type("UserAssigned")
)

// Deprecated version of BootDiagnostics. Use v1api20220301.BootDiagnostics instead
type BootDiagnostics_ARM struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	StorageUri *string `json:"storageUri,omitempty"`
}

// Deprecated version of DataDisk. Use v1api20220301.DataDisk instead
type DataDisk_ARM struct {
	Caching                 *Caching                   `json:"caching,omitempty"`
	CreateOption            *CreateOption              `json:"createOption,omitempty"`
	DeleteOption            *DeleteOption              `json:"deleteOption,omitempty"`
	DetachOption            *DetachOption              `json:"detachOption,omitempty"`
	DiskSizeGB              *int                       `json:"diskSizeGB,omitempty"`
	Image                   *VirtualHardDisk_ARM       `json:"image,omitempty"`
	Lun                     *int                       `json:"lun,omitempty"`
	ManagedDisk             *ManagedDiskParameters_ARM `json:"managedDisk,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	ToBeDetached            *bool                      `json:"toBeDetached,omitempty"`
	Vhd                     *VirtualHardDisk_ARM       `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                      `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of ImageReference. Use v1api20220301.ImageReference instead
type ImageReference_ARM struct {
	CommunityGalleryImageId *string `json:"communityGalleryImageId,omitempty"`
	Id                      *string `json:"id,omitempty"`
	Offer                   *string `json:"offer,omitempty"`
	Publisher               *string `json:"publisher,omitempty"`
	SharedGalleryImageId    *string `json:"sharedGalleryImageId,omitempty"`
	Sku                     *string `json:"sku,omitempty"`
	Version                 *string `json:"version,omitempty"`
}

// Deprecated version of LinuxConfiguration. Use v1api20220301.LinuxConfiguration instead
type LinuxConfiguration_ARM struct {
	DisablePasswordAuthentication *bool                   `json:"disablePasswordAuthentication,omitempty"`
	PatchSettings                 *LinuxPatchSettings_ARM `json:"patchSettings,omitempty"`
	ProvisionVMAgent              *bool                   `json:"provisionVMAgent,omitempty"`
	Ssh                           *SshConfiguration_ARM   `json:"ssh,omitempty"`
}

// Deprecated version of NetworkInterfaceReference. Use v1api20220301.NetworkInterfaceReference instead
type NetworkInterfaceReference_ARM struct {
	Id         *string                                  `json:"id,omitempty"`
	Properties *NetworkInterfaceReferenceProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of OSDisk. Use v1api20220301.OSDisk instead
type OSDisk_ARM struct {
	Caching                 *Caching                    `json:"caching,omitempty"`
	CreateOption            *CreateOption               `json:"createOption,omitempty"`
	DeleteOption            *DeleteOption               `json:"deleteOption,omitempty"`
	DiffDiskSettings        *DiffDiskSettings_ARM       `json:"diffDiskSettings,omitempty"`
	DiskSizeGB              *int                        `json:"diskSizeGB,omitempty"`
	EncryptionSettings      *DiskEncryptionSettings_ARM `json:"encryptionSettings,omitempty"`
	Image                   *VirtualHardDisk_ARM        `json:"image,omitempty"`
	ManagedDisk             *ManagedDiskParameters_ARM  `json:"managedDisk,omitempty"`
	Name                    *string                     `json:"name,omitempty"`
	OsType                  *OSDisk_OsType              `json:"osType,omitempty"`
	Vhd                     *VirtualHardDisk_ARM        `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                       `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of TerminateNotificationProfile. Use v1api20220301.TerminateNotificationProfile instead
type TerminateNotificationProfile_ARM struct {
	Enable           *bool   `json:"enable,omitempty"`
	NotBeforeTimeout *string `json:"notBeforeTimeout,omitempty"`
}

// Deprecated version of UefiSettings. Use v1api20220301.UefiSettings instead
type UefiSettings_ARM struct {
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
	VTpmEnabled       *bool `json:"vTpmEnabled,omitempty"`
}

// Deprecated version of VaultSecretGroup. Use v1api20220301.VaultSecretGroup instead
type VaultSecretGroup_ARM struct {
	SourceVault       *SubResource_ARM       `json:"sourceVault,omitempty"`
	VaultCertificates []VaultCertificate_ARM `json:"vaultCertificates,omitempty"`
}

// Deprecated version of VirtualMachineNetworkInterfaceConfiguration. Use v1api20220301.VirtualMachineNetworkInterfaceConfiguration instead
type VirtualMachineNetworkInterfaceConfiguration_ARM struct {
	Name       *string                                                    `json:"name,omitempty"`
	Properties *VirtualMachineNetworkInterfaceConfigurationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VMGalleryApplication. Use v1api20220301.VMGalleryApplication instead
type VMGalleryApplication_ARM struct {
	ConfigurationReference          *string `json:"configurationReference,omitempty"`
	EnableAutomaticUpgrade          *bool   `json:"enableAutomaticUpgrade,omitempty"`
	Order                           *int    `json:"order,omitempty"`
	PackageReferenceId              *string `json:"packageReferenceId,omitempty"`
	Tags                            *string `json:"tags,omitempty"`
	TreatFailureAsDeploymentFailure *bool   `json:"treatFailureAsDeploymentFailure,omitempty"`
}

// Deprecated version of VMSizeProperties. Use v1api20220301.VMSizeProperties instead
type VMSizeProperties_ARM struct {
	VCPUsAvailable *int `json:"vCPUsAvailable,omitempty"`
	VCPUsPerCore   *int `json:"vCPUsPerCore,omitempty"`
}

// Deprecated version of WindowsConfiguration. Use v1api20220301.WindowsConfiguration instead
type WindowsConfiguration_ARM struct {
	AdditionalUnattendContent []AdditionalUnattendContent_ARM `json:"additionalUnattendContent,omitempty"`
	EnableAutomaticUpdates    *bool                           `json:"enableAutomaticUpdates,omitempty"`
	PatchSettings             *PatchSettings_ARM              `json:"patchSettings,omitempty"`
	ProvisionVMAgent          *bool                           `json:"provisionVMAgent,omitempty"`
	TimeZone                  *string                         `json:"timeZone,omitempty"`
	WinRM                     *WinRMConfiguration_ARM         `json:"winRM,omitempty"`
}

// Deprecated version of AdditionalUnattendContent. Use v1api20220301.AdditionalUnattendContent instead
type AdditionalUnattendContent_ARM struct {
	ComponentName *AdditionalUnattendContent_ComponentName `json:"componentName,omitempty"`
	Content       *string                                  `json:"content,omitempty"`
	PassName      *AdditionalUnattendContent_PassName      `json:"passName,omitempty"`
	SettingName   *AdditionalUnattendContent_SettingName   `json:"settingName,omitempty"`
}

// Deprecated version of DiffDiskSettings. Use v1api20220301.DiffDiskSettings instead
type DiffDiskSettings_ARM struct {
	Option    *DiffDiskOption    `json:"option,omitempty"`
	Placement *DiffDiskPlacement `json:"placement,omitempty"`
}

// Deprecated version of DiskEncryptionSettings. Use v1api20220301.DiskEncryptionSettings instead
type DiskEncryptionSettings_ARM struct {
	DiskEncryptionKey *KeyVaultSecretReference_ARM `json:"diskEncryptionKey,omitempty"`
	Enabled           *bool                        `json:"enabled,omitempty"`
	KeyEncryptionKey  *KeyVaultKeyReference_ARM    `json:"keyEncryptionKey,omitempty"`
}

// Deprecated version of LinuxPatchSettings. Use v1api20220301.LinuxPatchSettings instead
type LinuxPatchSettings_ARM struct {
	AssessmentMode              *LinuxPatchSettings_AssessmentMode                `json:"assessmentMode,omitempty"`
	AutomaticByPlatformSettings *LinuxVMGuestPatchAutomaticByPlatformSettings_ARM `json:"automaticByPlatformSettings,omitempty"`
	PatchMode                   *LinuxPatchSettings_PatchMode                     `json:"patchMode,omitempty"`
}

// Deprecated version of ManagedDiskParameters. Use v1api20220301.ManagedDiskParameters instead
type ManagedDiskParameters_ARM struct {
	DiskEncryptionSet  *SubResource_ARM           `json:"diskEncryptionSet,omitempty"`
	Id                 *string                    `json:"id,omitempty"`
	SecurityProfile    *VMDiskSecurityProfile_ARM `json:"securityProfile,omitempty"`
	StorageAccountType *StorageAccountType        `json:"storageAccountType,omitempty"`
}

// Deprecated version of NetworkInterfaceReferenceProperties. Use v1api20220301.NetworkInterfaceReferenceProperties instead
type NetworkInterfaceReferenceProperties_ARM struct {
	DeleteOption *NetworkInterfaceReferenceProperties_DeleteOption `json:"deleteOption,omitempty"`
	Primary      *bool                                             `json:"primary,omitempty"`
}

// Deprecated version of PatchSettings. Use v1api20220301.PatchSettings instead
type PatchSettings_ARM struct {
	AssessmentMode              *PatchSettings_AssessmentMode                       `json:"assessmentMode,omitempty"`
	AutomaticByPlatformSettings *WindowsVMGuestPatchAutomaticByPlatformSettings_ARM `json:"automaticByPlatformSettings,omitempty"`
	EnableHotpatching           *bool                                               `json:"enableHotpatching,omitempty"`
	PatchMode                   *PatchSettings_PatchMode                            `json:"patchMode,omitempty"`
}

// Deprecated version of SshConfiguration. Use v1api20220301.SshConfiguration instead
type SshConfiguration_ARM struct {
	PublicKeys []SshPublicKeySpec_ARM `json:"publicKeys,omitempty"`
}

// Deprecated version of VaultCertificate. Use v1api20220301.VaultCertificate instead
type VaultCertificate_ARM struct {
	CertificateStore *string `json:"certificateStore,omitempty"`
	CertificateUrl   *string `json:"certificateUrl,omitempty"`
}

// Deprecated version of VirtualHardDisk. Use v1api20220301.VirtualHardDisk instead
type VirtualHardDisk_ARM struct {
	Uri *string `json:"uri,omitempty"`
}

// Deprecated version of VirtualMachineNetworkInterfaceConfigurationProperties. Use v1api20220301.VirtualMachineNetworkInterfaceConfigurationProperties instead
type VirtualMachineNetworkInterfaceConfigurationProperties_ARM struct {
	DeleteOption                *VirtualMachineNetworkInterfaceConfigurationProperties_DeleteOption `json:"deleteOption,omitempty"`
	DnsSettings                 *VirtualMachineNetworkInterfaceDnsSettingsConfiguration_ARM         `json:"dnsSettings,omitempty"`
	DscpConfiguration           *SubResource_ARM                                                    `json:"dscpConfiguration,omitempty"`
	EnableAcceleratedNetworking *bool                                                               `json:"enableAcceleratedNetworking,omitempty"`
	EnableFpga                  *bool                                                               `json:"enableFpga,omitempty"`
	EnableIPForwarding          *bool                                                               `json:"enableIPForwarding,omitempty"`
	IpConfigurations            []VirtualMachineNetworkInterfaceIPConfiguration_ARM                 `json:"ipConfigurations,omitempty"`
	NetworkSecurityGroup        *SubResource_ARM                                                    `json:"networkSecurityGroup,omitempty"`
	Primary                     *bool                                                               `json:"primary,omitempty"`
}

// Deprecated version of WinRMConfiguration. Use v1api20220301.WinRMConfiguration instead
type WinRMConfiguration_ARM struct {
	Listeners []WinRMListener_ARM `json:"listeners,omitempty"`
}

// Deprecated version of KeyVaultKeyReference. Use v1api20220301.KeyVaultKeyReference instead
type KeyVaultKeyReference_ARM struct {
	KeyUrl      *string          `json:"keyUrl,omitempty"`
	SourceVault *SubResource_ARM `json:"sourceVault,omitempty"`
}

// Deprecated version of KeyVaultSecretReference. Use v1api20220301.KeyVaultSecretReference instead
type KeyVaultSecretReference_ARM struct {
	SecretUrl   *string          `json:"secretUrl,omitempty"`
	SourceVault *SubResource_ARM `json:"sourceVault,omitempty"`
}

// Deprecated version of LinuxVMGuestPatchAutomaticByPlatformSettings. Use v1api20220301.LinuxVMGuestPatchAutomaticByPlatformSettings instead
type LinuxVMGuestPatchAutomaticByPlatformSettings_ARM struct {
	RebootSetting *LinuxVMGuestPatchAutomaticByPlatformSettings_RebootSetting `json:"rebootSetting,omitempty"`
}

// Deprecated version of SshPublicKeySpec. Use v1api20220301.SshPublicKeySpec instead
type SshPublicKeySpec_ARM struct {
	KeyData *string `json:"keyData,omitempty"`
	Path    *string `json:"path,omitempty"`
}

// Deprecated version of VirtualMachineNetworkInterfaceDnsSettingsConfiguration. Use v1api20220301.VirtualMachineNetworkInterfaceDnsSettingsConfiguration instead
type VirtualMachineNetworkInterfaceDnsSettingsConfiguration_ARM struct {
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Deprecated version of VirtualMachineNetworkInterfaceIPConfiguration. Use v1api20220301.VirtualMachineNetworkInterfaceIPConfiguration instead
type VirtualMachineNetworkInterfaceIPConfiguration_ARM struct {
	Name       *string                                                      `json:"name,omitempty"`
	Properties *VirtualMachineNetworkInterfaceIPConfigurationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of VMDiskSecurityProfile. Use v1api20220301.VMDiskSecurityProfile instead
type VMDiskSecurityProfile_ARM struct {
	DiskEncryptionSet      *SubResource_ARM                              `json:"diskEncryptionSet,omitempty"`
	SecurityEncryptionType *VMDiskSecurityProfile_SecurityEncryptionType `json:"securityEncryptionType,omitempty"`
}

// Deprecated version of WindowsVMGuestPatchAutomaticByPlatformSettings. Use v1api20220301.WindowsVMGuestPatchAutomaticByPlatformSettings instead
type WindowsVMGuestPatchAutomaticByPlatformSettings_ARM struct {
	RebootSetting *WindowsVMGuestPatchAutomaticByPlatformSettings_RebootSetting `json:"rebootSetting,omitempty"`
}

// Deprecated version of WinRMListener. Use v1api20220301.WinRMListener instead
type WinRMListener_ARM struct {
	CertificateUrl *string                 `json:"certificateUrl,omitempty"`
	Protocol       *WinRMListener_Protocol `json:"protocol,omitempty"`
}

// Deprecated version of VirtualMachineNetworkInterfaceIPConfigurationProperties. Use v1api20220301.VirtualMachineNetworkInterfaceIPConfigurationProperties instead
type VirtualMachineNetworkInterfaceIPConfigurationProperties_ARM struct {
	ApplicationGatewayBackendAddressPools []SubResource_ARM                                                                `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []SubResource_ARM                                                                `json:"applicationSecurityGroups,omitempty"`
	LoadBalancerBackendAddressPools       []SubResource_ARM                                                                `json:"loadBalancerBackendAddressPools,omitempty"`
	Primary                               *bool                                                                            `json:"primary,omitempty"`
	PrivateIPAddressVersion               *VirtualMachineNetworkInterfaceIPConfigurationProperties_PrivateIPAddressVersion `json:"privateIPAddressVersion,omitempty"`
	PublicIPAddressConfiguration          *VirtualMachinePublicIPAddressConfiguration_ARM                                  `json:"publicIPAddressConfiguration,omitempty"`
	Subnet                                *SubResource_ARM                                                                 `json:"subnet,omitempty"`
}

// Deprecated version of VirtualMachinePublicIPAddressConfiguration. Use v1api20220301.VirtualMachinePublicIPAddressConfiguration instead
type VirtualMachinePublicIPAddressConfiguration_ARM struct {
	Name       *string                                                   `json:"name,omitempty"`
	Properties *VirtualMachinePublicIPAddressConfigurationProperties_ARM `json:"properties,omitempty"`
	Sku        *PublicIPAddressSku_ARM                                   `json:"sku,omitempty"`
}

// Deprecated version of PublicIPAddressSku. Use v1api20220301.PublicIPAddressSku instead
type PublicIPAddressSku_ARM struct {
	Name *PublicIPAddressSku_Name `json:"name,omitempty"`
	Tier *PublicIPAddressSku_Tier `json:"tier,omitempty"`
}

// Deprecated version of VirtualMachinePublicIPAddressConfigurationProperties. Use v1api20220301.VirtualMachinePublicIPAddressConfigurationProperties instead
type VirtualMachinePublicIPAddressConfigurationProperties_ARM struct {
	DeleteOption             *VirtualMachinePublicIPAddressConfigurationProperties_DeleteOption             `json:"deleteOption,omitempty"`
	DnsSettings              *VirtualMachinePublicIPAddressDnsSettingsConfiguration_ARM                     `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int                                                                           `json:"idleTimeoutInMinutes,omitempty"`
	IpTags                   []VirtualMachineIpTag_ARM                                                      `json:"ipTags,omitempty"`
	PublicIPAddressVersion   *VirtualMachinePublicIPAddressConfigurationProperties_PublicIPAddressVersion   `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *VirtualMachinePublicIPAddressConfigurationProperties_PublicIPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_ARM                                                               `json:"publicIPPrefix,omitempty"`
}

// Deprecated version of VirtualMachineIpTag. Use v1api20220301.VirtualMachineIpTag instead
type VirtualMachineIpTag_ARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of VirtualMachinePublicIPAddressDnsSettingsConfiguration. Use v1api20220301.VirtualMachinePublicIPAddressDnsSettingsConfiguration instead
type VirtualMachinePublicIPAddressDnsSettingsConfiguration_ARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
}
