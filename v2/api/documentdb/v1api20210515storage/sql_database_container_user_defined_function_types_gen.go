// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontaineruserdefinedfunctions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontaineruserdefinedfunctions/status,sqldatabasecontaineruserdefinedfunctions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210515.SqlDatabaseContainerUserDefinedFunction
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}
type SqlDatabaseContainerUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerUserDefinedFunction{}

// GetConditions returns the conditions of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetConditions() conditions.Conditions {
	return function.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (function *SqlDatabaseContainerUserDefinedFunction) SetConditions(conditions conditions.Conditions) {
	function.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerUserDefinedFunction{}

// AzureName returns the Azure name of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) AzureName() string {
	return function.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (function SqlDatabaseContainerUserDefinedFunction) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetSpec() genruntime.ConvertibleSpec {
	return &function.Spec
}

// GetStatus returns the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetStatus() genruntime.ConvertibleStatus {
	return &function.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (function *SqlDatabaseContainerUserDefinedFunction) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

// NewEmptyStatus returns a new empty (blank) status
func (function *SqlDatabaseContainerUserDefinedFunction) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (function *SqlDatabaseContainerUserDefinedFunction) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(function.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  function.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS); ok {
		function.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	function.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerUserDefinedFunction is the hub type for conversion
func (function *SqlDatabaseContainerUserDefinedFunction) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (function *SqlDatabaseContainerUserDefinedFunction) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: function.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerUserDefinedFunction",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.SqlDatabaseContainerUserDefinedFunction
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}
type SqlDatabaseContainerUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerUserDefinedFunction `json:"items"`
}

// Storage version of v1api20210515.DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec
type DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionResource    `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec from the provided source
func (function *DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == function {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(function)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec
func (function *DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == function {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(function)
}

// Storage version of v1api20210515.DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS
type DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS struct {
	Conditions  []conditions.Condition                               `json:"conditions,omitempty"`
	Id          *string                                              `json:"id,omitempty"`
	Location    *string                                              `json:"location,omitempty"`
	Name        *string                                              `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                    `json:"tags,omitempty"`
	Type        *string                                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS from the provided source
func (function *DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == function {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(function)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS
func (function *DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == function {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(function)
}

// Storage version of v1api20210515.SqlUserDefinedFunctionGetProperties_Resource_STATUS
type SqlUserDefinedFunctionGetProperties_Resource_STATUS struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// Storage version of v1api20210515.SqlUserDefinedFunctionResource
// Cosmos DB SQL userDefinedFunction resource object
type SqlUserDefinedFunctionResource struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerUserDefinedFunction{}, &SqlDatabaseContainerUserDefinedFunctionList{})
}
