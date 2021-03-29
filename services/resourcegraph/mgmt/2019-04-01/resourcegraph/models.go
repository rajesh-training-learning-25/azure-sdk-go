package resourcegraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"

// Column query result column descriptor.
type Column struct {
	// Name - Column name.
	Name *string `json:"name,omitempty"`
	// Type - Column data type. Possible values include: 'String', 'Integer', 'Number', 'Boolean', 'Object'
	Type ColumnDataType `json:"type,omitempty"`
}

// Error error details.
type Error struct {
	// Code - Error code identifying the specific error.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Details - Error details
	Details *[]ErrorDetails `json:"details,omitempty"`
}

// ErrorDetails ...
type ErrorDetails struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Code - Error code identifying the specific error.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDetails.
func (ed ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ed.Code != nil {
		objectMap["code"] = ed.Code
	}
	if ed.Message != nil {
		objectMap["message"] = ed.Message
	}
	for k, v := range ed.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ErrorDetails struct.
func (ed *ErrorDetails) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if ed.AdditionalProperties == nil {
					ed.AdditionalProperties = make(map[string]interface{})
				}
				ed.AdditionalProperties[k] = additionalProperties
			}
		case "code":
			if v != nil {
				var code string
				err = json.Unmarshal(*v, &code)
				if err != nil {
					return err
				}
				ed.Code = &code
			}
		case "message":
			if v != nil {
				var message string
				err = json.Unmarshal(*v, &message)
				if err != nil {
					return err
				}
				ed.Message = &message
			}
		}
	}

	return nil
}

// ErrorResponse an error response from the API.
type ErrorResponse struct {
	// Error - Error information.
	Error *Error `json:"error,omitempty"`
}

// BasicFacet a facet containing additional statistics on the response of a query. Can be either FacetResult or
// FacetError.
type BasicFacet interface {
	AsFacetResult() (*FacetResult, bool)
	AsFacetError() (*FacetError, bool)
	AsFacet() (*Facet, bool)
}

// Facet a facet containing additional statistics on the response of a query. Can be either FacetResult or
// FacetError.
type Facet struct {
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

func unmarshalBasicFacet(body []byte) (BasicFacet, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["resultType"] {
	case string(ResultTypeFacetResult):
		var fr FacetResult
		err := json.Unmarshal(body, &fr)
		return fr, err
	case string(ResultTypeFacetError):
		var fe FacetError
		err := json.Unmarshal(body, &fe)
		return fe, err
	default:
		var f Facet
		err := json.Unmarshal(body, &f)
		return f, err
	}
}
func unmarshalBasicFacetArray(body []byte) ([]BasicFacet, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	fArray := make([]BasicFacet, len(rawMessages))

	for index, rawMessage := range rawMessages {
		f, err := unmarshalBasicFacet(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

// MarshalJSON is the custom marshaler for Facet.
func (f Facet) MarshalJSON() ([]byte, error) {
	f.ResultType = ResultTypeFacet
	objectMap := make(map[string]interface{})
	if f.Expression != nil {
		objectMap["expression"] = f.Expression
	}
	if f.ResultType != "" {
		objectMap["resultType"] = f.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for Facet.
func (f Facet) AsFacetResult() (*FacetResult, bool) {
	return nil, false
}

// AsFacetError is the BasicFacet implementation for Facet.
func (f Facet) AsFacetError() (*FacetError, bool) {
	return nil, false
}

// AsFacet is the BasicFacet implementation for Facet.
func (f Facet) AsFacet() (*Facet, bool) {
	return &f, true
}

// AsBasicFacet is the BasicFacet implementation for Facet.
func (f Facet) AsBasicFacet() (BasicFacet, bool) {
	return &f, true
}

// FacetError a facet whose execution resulted in an error.
type FacetError struct {
	// Errors - An array containing detected facet errors with details.
	Errors *[]ErrorDetails `json:"errors,omitempty"`
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

// MarshalJSON is the custom marshaler for FacetError.
func (fe FacetError) MarshalJSON() ([]byte, error) {
	fe.ResultType = ResultTypeFacetError
	objectMap := make(map[string]interface{})
	if fe.Errors != nil {
		objectMap["errors"] = fe.Errors
	}
	if fe.Expression != nil {
		objectMap["expression"] = fe.Expression
	}
	if fe.ResultType != "" {
		objectMap["resultType"] = fe.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacetResult() (*FacetResult, bool) {
	return nil, false
}

// AsFacetError is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacetError() (*FacetError, bool) {
	return &fe, true
}

// AsFacet is the BasicFacet implementation for FacetError.
func (fe FacetError) AsFacet() (*Facet, bool) {
	return nil, false
}

// AsBasicFacet is the BasicFacet implementation for FacetError.
func (fe FacetError) AsBasicFacet() (BasicFacet, bool) {
	return &fe, true
}

// FacetRequest a request to compute additional statistics (facets) over the query results.
type FacetRequest struct {
	// Expression - The column or list of columns to summarize by
	Expression *string `json:"expression,omitempty"`
	// Options - The options for facet evaluation
	Options *FacetRequestOptions `json:"options,omitempty"`
}

// FacetRequestOptions the options for facet evaluation
type FacetRequestOptions struct {
	// SortBy - The column name or query expression to sort on. Defaults to count if not present.
	SortBy *string `json:"sortBy,omitempty"`
	// SortOrder - The sorting order by the selected column (count by default). Possible values include: 'Asc', 'Desc'
	SortOrder FacetSortOrder `json:"sortOrder,omitempty"`
	// Filter - Specifies the filter condition for the 'where' clause which will be run on main query's result, just before the actual faceting.
	Filter *string `json:"filter,omitempty"`
	// Top - The maximum number of facet rows that should be returned.
	Top *int32 `json:"$top,omitempty"`
}

// FacetResult successfully executed facet containing additional statistics on the response of a query.
type FacetResult struct {
	// TotalRecords - Number of total records in the facet results.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Count - Number of records returned in the facet response.
	Count *int32 `json:"count,omitempty"`
	// Data - A table containing the desired facets. Only present if the facet is valid.
	Data interface{} `json:"data,omitempty"`
	// Expression - Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`
	// ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
	ResultType ResultType `json:"resultType,omitempty"`
}

// MarshalJSON is the custom marshaler for FacetResult.
func (fr FacetResult) MarshalJSON() ([]byte, error) {
	fr.ResultType = ResultTypeFacetResult
	objectMap := make(map[string]interface{})
	if fr.TotalRecords != nil {
		objectMap["totalRecords"] = fr.TotalRecords
	}
	if fr.Count != nil {
		objectMap["count"] = fr.Count
	}
	if fr.Data != nil {
		objectMap["data"] = fr.Data
	}
	if fr.Expression != nil {
		objectMap["expression"] = fr.Expression
	}
	if fr.ResultType != "" {
		objectMap["resultType"] = fr.ResultType
	}
	return json.Marshal(objectMap)
}

// AsFacetResult is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacetResult() (*FacetResult, bool) {
	return &fr, true
}

// AsFacetError is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacetError() (*FacetError, bool) {
	return nil, false
}

// AsFacet is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsFacet() (*Facet, bool) {
	return nil, false
}

// AsBasicFacet is the BasicFacet implementation for FacetResult.
func (fr FacetResult) AsBasicFacet() (BasicFacet, bool) {
	return &fr, true
}

// Operation resource Graph REST API operation definition.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The origin of operations.
	Origin *string `json:"origin,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft Resource Graph.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description for the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Resource Graph operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Resource Graph operations supported by the Resource Graph resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// QueryRequest describes a query to be executed.
type QueryRequest struct {
	// Subscriptions - Azure subscriptions against which to execute the query.
	Subscriptions *[]string `json:"subscriptions,omitempty"`
	// Query - The resources query.
	Query *string `json:"query,omitempty"`
	// Options - The query evaluation options
	Options *QueryRequestOptions `json:"options,omitempty"`
	// Facets - An array of facet requests to be computed against the query result.
	Facets *[]FacetRequest `json:"facets,omitempty"`
}

// QueryRequestOptions the options for query evaluation
type QueryRequestOptions struct {
	// SkipToken - Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string `json:"$skipToken,omitempty"`
	// Top - The maximum number of rows that the query should return. Overrides the page size when ```$skipToken``` property is present.
	Top *int32 `json:"$top,omitempty"`
	// Skip - The number of rows to skip from the beginning of the results. Overrides the next page offset when ```$skipToken``` property is present.
	Skip *int32 `json:"$skip,omitempty"`
	// ResultFormat - Defines in which format query result returned. Possible values include: 'ResultFormatTable', 'ResultFormatObjectArray'
	ResultFormat ResultFormat `json:"resultFormat,omitempty"`
}

// QueryResponse query result.
type QueryResponse struct {
	autorest.Response `json:"-"`
	// TotalRecords - Number of total records matching the query.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
	// Count - Number of records returned in the current response. In the case of paging, this is the number of records in the current page.
	Count *int64 `json:"count,omitempty"`
	// ResultTruncated - Indicates whether the query results are truncated. Possible values include: 'True', 'False'
	ResultTruncated ResultTruncated `json:"resultTruncated,omitempty"`
	// SkipToken - When present, the value can be passed to a subsequent query call (together with the same query and subscriptions used in the current request) to retrieve the next page of data.
	SkipToken *string `json:"$skipToken,omitempty"`
	// Data - Query output in tabular format.
	Data interface{} `json:"data,omitempty"`
	// Facets - Query facets.
	Facets *[]BasicFacet `json:"facets,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for QueryResponse struct.
func (qr *QueryResponse) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "totalRecords":
			if v != nil {
				var totalRecords int64
				err = json.Unmarshal(*v, &totalRecords)
				if err != nil {
					return err
				}
				qr.TotalRecords = &totalRecords
			}
		case "count":
			if v != nil {
				var count int64
				err = json.Unmarshal(*v, &count)
				if err != nil {
					return err
				}
				qr.Count = &count
			}
		case "resultTruncated":
			if v != nil {
				var resultTruncated ResultTruncated
				err = json.Unmarshal(*v, &resultTruncated)
				if err != nil {
					return err
				}
				qr.ResultTruncated = resultTruncated
			}
		case "$skipToken":
			if v != nil {
				var skipToken string
				err = json.Unmarshal(*v, &skipToken)
				if err != nil {
					return err
				}
				qr.SkipToken = &skipToken
			}
		case "data":
			if v != nil {
				var data interface{}
				err = json.Unmarshal(*v, &data)
				if err != nil {
					return err
				}
				qr.Data = data
			}
		case "facets":
			if v != nil {
				facets, err := unmarshalBasicFacetArray(*v)
				if err != nil {
					return err
				}
				qr.Facets = &facets
			}
		}
	}

	return nil
}

// Table query output in tabular format.
type Table struct {
	// Columns - Query result column descriptors.
	Columns *[]Column `json:"columns,omitempty"`
	// Rows - Query result rows.
	Rows *[][]interface{} `json:"rows,omitempty"`
}
