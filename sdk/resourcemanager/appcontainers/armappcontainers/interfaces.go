//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

// JavaComponentPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetJavaComponentProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *JavaComponentProperties, *NacosComponent, *SpringBootAdminComponent, *SpringCloudConfigComponent, *SpringCloudEurekaComponent,
// - *SpringCloudGatewayComponent
type JavaComponentPropertiesClassification interface {
	// GetJavaComponentProperties returns the JavaComponentProperties content of the underlying type.
	GetJavaComponentProperties() *JavaComponentProperties
}