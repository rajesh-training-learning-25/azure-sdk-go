package resourcegraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ColumnDataType enumerates the values for column data type.
type ColumnDataType string

const (
	// Boolean ...
	Boolean ColumnDataType = "boolean"
	// Integer ...
	Integer ColumnDataType = "integer"
	// Number ...
	Number ColumnDataType = "number"
	// Object ...
	Object ColumnDataType = "object"
	// String ...
	String ColumnDataType = "string"
)

// PossibleColumnDataTypeValues returns an array of possible values for the ColumnDataType const type.
func PossibleColumnDataTypeValues() []ColumnDataType {
	return []ColumnDataType{Boolean, Integer, Number, Object, String}
}

// FacetSortOrder enumerates the values for facet sort order.
type FacetSortOrder string

const (
	// Asc ...
	Asc FacetSortOrder = "asc"
	// Desc ...
	Desc FacetSortOrder = "desc"
)

// PossibleFacetSortOrderValues returns an array of possible values for the FacetSortOrder const type.
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return []FacetSortOrder{Asc, Desc}
}

// ResultKind enumerates the values for result kind.
type ResultKind string

const (
	// Basic ...
	Basic ResultKind = "basic"
)

// PossibleResultKindValues returns an array of possible values for the ResultKind const type.
func PossibleResultKindValues() []ResultKind {
	return []ResultKind{Basic}
}

// ResultTruncated enumerates the values for result truncated.
type ResultTruncated string

const (
	// False ...
	False ResultTruncated = "false"
	// True ...
	True ResultTruncated = "true"
)

// PossibleResultTruncatedValues returns an array of possible values for the ResultTruncated const type.
func PossibleResultTruncatedValues() []ResultTruncated {
	return []ResultTruncated{False, True}
}

// ResultType enumerates the values for result type.
type ResultType string

const (
	// ResultTypeFacet ...
	ResultTypeFacet ResultType = "Facet"
	// ResultTypeFacetError ...
	ResultTypeFacetError ResultType = "FacetError"
	// ResultTypeFacetResult ...
	ResultTypeFacetResult ResultType = "FacetResult"
)

// PossibleResultTypeValues returns an array of possible values for the ResultType const type.
func PossibleResultTypeValues() []ResultType {
	return []ResultType{ResultTypeFacet, ResultTypeFacetError, ResultTypeFacetResult}
}
