//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	moduleVersion = "v0.9.0"
)

// AuthorizationScopeFilter - Defines what level of authorization resources should be returned based on the which subscriptions
// and management groups are passed as scopes.
type AuthorizationScopeFilter string

const (
	AuthorizationScopeFilterAtScopeAboveAndBelow AuthorizationScopeFilter = "AtScopeAboveAndBelow"
	AuthorizationScopeFilterAtScopeAndAbove      AuthorizationScopeFilter = "AtScopeAndAbove"
	AuthorizationScopeFilterAtScopeAndBelow      AuthorizationScopeFilter = "AtScopeAndBelow"
	AuthorizationScopeFilterAtScopeExact         AuthorizationScopeFilter = "AtScopeExact"
)

// PossibleAuthorizationScopeFilterValues returns the possible values for the AuthorizationScopeFilter const type.
func PossibleAuthorizationScopeFilterValues() []AuthorizationScopeFilter {
	return []AuthorizationScopeFilter{
		AuthorizationScopeFilterAtScopeAboveAndBelow,
		AuthorizationScopeFilterAtScopeAndAbove,
		AuthorizationScopeFilterAtScopeAndBelow,
		AuthorizationScopeFilterAtScopeExact,
	}
}

// ColumnDataType - Data type of a column in a table.
type ColumnDataType string

const (
	ColumnDataTypeBoolean  ColumnDataType = "boolean"
	ColumnDataTypeDatetime ColumnDataType = "datetime"
	ColumnDataTypeInteger  ColumnDataType = "integer"
	ColumnDataTypeNumber   ColumnDataType = "number"
	ColumnDataTypeObject   ColumnDataType = "object"
	ColumnDataTypeString   ColumnDataType = "string"
)

// PossibleColumnDataTypeValues returns the possible values for the ColumnDataType const type.
func PossibleColumnDataTypeValues() []ColumnDataType {
	return []ColumnDataType{
		ColumnDataTypeBoolean,
		ColumnDataTypeDatetime,
		ColumnDataTypeInteger,
		ColumnDataTypeNumber,
		ColumnDataTypeObject,
		ColumnDataTypeString,
	}
}

// FacetSortOrder - The sorting order by the selected column (count by default).
type FacetSortOrder string

const (
	FacetSortOrderAsc  FacetSortOrder = "asc"
	FacetSortOrderDesc FacetSortOrder = "desc"
)

// PossibleFacetSortOrderValues returns the possible values for the FacetSortOrder const type.
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return []FacetSortOrder{
		FacetSortOrderAsc,
		FacetSortOrderDesc,
	}
}

// ResultFormat - Defines in which format query result returned.
type ResultFormat string

const (
	ResultFormatObjectArray ResultFormat = "objectArray"
	ResultFormatTable       ResultFormat = "table"
)

// PossibleResultFormatValues returns the possible values for the ResultFormat const type.
func PossibleResultFormatValues() []ResultFormat {
	return []ResultFormat{
		ResultFormatObjectArray,
		ResultFormatTable,
	}
}

// ResultTruncated - Indicates whether the query results are truncated.
type ResultTruncated string

const (
	ResultTruncatedFalse ResultTruncated = "false"
	ResultTruncatedTrue  ResultTruncated = "true"
)

// PossibleResultTruncatedValues returns the possible values for the ResultTruncated const type.
func PossibleResultTruncatedValues() []ResultTruncated {
	return []ResultTruncated{
		ResultTruncatedFalse,
		ResultTruncatedTrue,
	}
}
