// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20210601 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601"
	v1api20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601storage"
	v1api20220120p "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview"
	v1api20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120previewstorage"
	v20210601 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
	v20220120p "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20220120preview"
	v20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20220120previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServerExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServerExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20210601.FlexibleServer{},
		&v1api20210601s.FlexibleServer{},
		&v1api20220120p.FlexibleServer{},
		&v1api20220120ps.FlexibleServer{},
		&v20210601.FlexibleServer{},
		&v20210601s.FlexibleServer{},
		&v20220120p.FlexibleServer{},
		&v20220120ps.FlexibleServer{}}
}
