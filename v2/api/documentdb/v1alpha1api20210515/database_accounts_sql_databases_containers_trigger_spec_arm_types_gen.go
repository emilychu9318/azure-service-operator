// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec. Use v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec instead
type DatabaseAccounts_SqlDatabases_Containers_Trigger_SpecARM struct {
	Location   *string                              `json:"location,omitempty"`
	Name       string                               `json:"name,omitempty"`
	Properties *SqlTriggerCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_Containers_Trigger_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (trigger DatabaseAccounts_SqlDatabases_Containers_Trigger_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_SpecARM) GetName() string {
	return trigger.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// Deprecated version of SqlTriggerCreateUpdateProperties. Use v1beta20210515.SqlTriggerCreateUpdateProperties instead
type SqlTriggerCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM `json:"options,omitempty"`
	Resource *SqlTriggerResourceARM  `json:"resource,omitempty"`
}

// Deprecated version of SqlTriggerResource. Use v1beta20210515.SqlTriggerResource instead
type SqlTriggerResourceARM struct {
	Body             *string                              `json:"body,omitempty"`
	Id               *string                              `json:"id,omitempty"`
	TriggerOperation *SqlTriggerResource_TriggerOperation `json:"triggerOperation,omitempty"`
	TriggerType      *SqlTriggerResource_TriggerType      `json:"triggerType,omitempty"`
}

// Deprecated version of SqlTriggerResource_TriggerOperation. Use v1beta20210515.SqlTriggerResource_TriggerOperation instead
// +kubebuilder:validation:Enum={"All","Create","Delete","Replace","Update"}
type SqlTriggerResource_TriggerOperation string

const (
	SqlTriggerResource_TriggerOperation_All     = SqlTriggerResource_TriggerOperation("All")
	SqlTriggerResource_TriggerOperation_Create  = SqlTriggerResource_TriggerOperation("Create")
	SqlTriggerResource_TriggerOperation_Delete  = SqlTriggerResource_TriggerOperation("Delete")
	SqlTriggerResource_TriggerOperation_Replace = SqlTriggerResource_TriggerOperation("Replace")
	SqlTriggerResource_TriggerOperation_Update  = SqlTriggerResource_TriggerOperation("Update")
)

// Deprecated version of SqlTriggerResource_TriggerType. Use v1beta20210515.SqlTriggerResource_TriggerType instead
// +kubebuilder:validation:Enum={"Post","Pre"}
type SqlTriggerResource_TriggerType string

const (
	SqlTriggerResource_TriggerType_Post = SqlTriggerResource_TriggerType("Post")
	SqlTriggerResource_TriggerType_Pre  = SqlTriggerResource_TriggerType("Pre")
)