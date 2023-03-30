// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisEnterprise_Spec. Use v1api20210301.RedisEnterprise_Spec instead
type RedisEnterprise_Spec_ARM struct {
	Location   *string                `json:"location,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Properties *ClusterProperties_ARM `json:"properties,omitempty"`
	Sku        *Sku_ARM               `json:"sku,omitempty"`
	Tags       map[string]string      `json:"tags,omitempty"`
	Zones      []string               `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterprise_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (enterprise RedisEnterprise_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (enterprise *RedisEnterprise_Spec_ARM) GetName() string {
	return enterprise.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// Deprecated version of ClusterProperties. Use v1api20210301.ClusterProperties instead
type ClusterProperties_ARM struct {
	MinimumTlsVersion *ClusterProperties_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`
}

// Deprecated version of Sku. Use v1api20210301.Sku instead
type Sku_ARM struct {
	Capacity *int      `json:"capacity,omitempty"`
	Name     *Sku_Name `json:"name,omitempty"`
}

// Deprecated version of Sku_Name. Use v1api20210301.Sku_Name instead
// +kubebuilder:validation:Enum={"EnterpriseFlash_F1500","EnterpriseFlash_F300","EnterpriseFlash_F700","Enterprise_E10","Enterprise_E100","Enterprise_E20","Enterprise_E50"}
type Sku_Name string

const (
	Sku_Name_EnterpriseFlash_F1500 = Sku_Name("EnterpriseFlash_F1500")
	Sku_Name_EnterpriseFlash_F300  = Sku_Name("EnterpriseFlash_F300")
	Sku_Name_EnterpriseFlash_F700  = Sku_Name("EnterpriseFlash_F700")
	Sku_Name_Enterprise_E10        = Sku_Name("Enterprise_E10")
	Sku_Name_Enterprise_E100       = Sku_Name("Enterprise_E100")
	Sku_Name_Enterprise_E20        = Sku_Name("Enterprise_E20")
	Sku_Name_Enterprise_E50        = Sku_Name("Enterprise_E50")
)
