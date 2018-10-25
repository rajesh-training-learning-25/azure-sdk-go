package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go//services/preview/storage/mgmt/2015-05-01-preview/storage"

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available ...
	Available AccountStatus = "Available"
	// Unavailable ...
	Unavailable AccountStatus = "Unavailable"
)

// PossibleAccountStatusValues returns an array of possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{Available, Unavailable}
}

// AccountType enumerates the values for account type.
type AccountType string

const (
	// PremiumLRS ...
	PremiumLRS AccountType = "Premium_LRS"
	// StandardGRS ...
	StandardGRS AccountType = "Standard_GRS"
	// StandardLRS ...
	StandardLRS AccountType = "Standard_LRS"
	// StandardRAGRS ...
	StandardRAGRS AccountType = "Standard_RAGRS"
	// StandardZRS ...
	StandardZRS AccountType = "Standard_ZRS"
)

// PossibleAccountTypeValues returns an array of possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{PremiumLRS, StandardGRS, StandardLRS, StandardRAGRS, StandardZRS}
}

// KeyName enumerates the values for key name.
type KeyName string

const (
	// Key1 ...
	Key1 KeyName = "key1"
	// Key2 ...
	Key2 KeyName = "key2"
)

// PossibleKeyNameValues returns an array of possible values for the KeyName const type.
func PossibleKeyNameValues() []KeyName {
	return []KeyName{Key1, Key2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, ResolvingDNS, Succeeded}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Bytes, BytesPerSecond, Count, CountsPerSecond, Percent, Seconds}
}

// Account the storage account.
type Account struct {
	autorest.Response  `json:"-"`
	*AccountProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		}
	}

	return nil
}

// AccountCheckNameAvailabilityParameters ...
type AccountCheckNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AccountCreateParameters the parameters to provide for the account.
type AccountCreateParameters struct {
	*AccountPropertiesCreateParameters `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.AccountPropertiesCreateParameters != nil {
		objectMap["properties"] = acp.AccountPropertiesCreateParameters
	}
	if acp.ID != nil {
		objectMap["id"] = acp.ID
	}
	if acp.Name != nil {
		objectMap["name"] = acp.Name
	}
	if acp.Type != nil {
		objectMap["type"] = acp.Type
	}
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountCreateParameters struct.
func (acp *AccountCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountPropertiesCreateParameters AccountPropertiesCreateParameters
				err = json.Unmarshal(*v, &accountPropertiesCreateParameters)
				if err != nil {
					return err
				}
				acp.AccountPropertiesCreateParameters = &accountPropertiesCreateParameters
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				acp.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				acp.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				acp.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				acp.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				acp.Tags = tags
			}
		}
	}

	return nil
}

// AccountKeys the access keys for the storage account.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Gets the value of key 1.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Gets the value of key 2.
	Key2 *string `json:"key2,omitempty"`
}

// AccountListResult the list storage accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets the list of storage accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - Gets the link to the next set of results. Currently this will always be empty as the API does not support pagination.
	NextLink *string `json:"nextLink,omitempty"`
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AccountListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListResultIterator) Response() AccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListResultIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer(ctx context.Context) (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(context.Context, AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AccountListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListResultPage) Values() []Account {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// AccountProperties ...
type AccountProperties struct {
	// ProvisioningState - Gets the status of the storage account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AccountType - Gets the type of the storage account. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
	// PrimaryEndpoints - Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
	PrimaryEndpoints *Endpoints `json:"primaryEndpoints,omitempty"`
	// PrimaryLocation - Gets the location of the primary for the storage account.
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// StatusOfPrimary - Gets the status indicating whether the primary location of the storage account is available or unavailable. Possible values include: 'Available', 'Unavailable'
	StatusOfPrimary AccountStatus `json:"statusOfPrimary,omitempty"`
	// LastGeoFailoverTime - Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
	LastGeoFailoverTime *date.Time `json:"lastGeoFailoverTime,omitempty"`
	// SecondaryLocation - Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
	// StatusOfSecondary - Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS. Possible values include: 'Available', 'Unavailable'
	StatusOfSecondary AccountStatus `json:"statusOfSecondary,omitempty"`
	// CreationTime - Gets the creation date and time of the storage account in UTC.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// CustomDomain - Gets the user assigned custom domain assigned to this storage account.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	// SecondaryEndpoints - Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
	SecondaryEndpoints *Endpoints `json:"secondaryEndpoints,omitempty"`
}

// AccountPropertiesCreateParameters ...
type AccountPropertiesCreateParameters struct {
	// AccountType - Gets or sets the account type. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
}

// AccountPropertiesUpdateParameters ...
type AccountPropertiesUpdateParameters struct {
	// AccountType - Gets or sets the account type. Note that StandardZRS and PremiumLRS accounts cannot be changed to other account types, and other account types cannot be changed to StandardZRS or PremiumLRS. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
	// CustomDomain - User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
}

// AccountRegenerateKeyParameters ...
type AccountRegenerateKeyParameters struct {
	// KeyName - Possible values include: 'Key1', 'Key2'
	KeyName KeyName `json:"keyName,omitempty"`
}

// AccountsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type AccountsCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AccountsCreateFuture) Result(client AccountsClient) (a Account, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.AccountsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("storage.AccountsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.CreateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storage.AccountsCreateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AccountUpdateParameters the parameters to update on the account.
type AccountUpdateParameters struct {
	*AccountPropertiesUpdateParameters `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.AccountPropertiesUpdateParameters != nil {
		objectMap["properties"] = aup.AccountPropertiesUpdateParameters
	}
	if aup.ID != nil {
		objectMap["id"] = aup.ID
	}
	if aup.Name != nil {
		objectMap["name"] = aup.Name
	}
	if aup.Type != nil {
		objectMap["type"] = aup.Type
	}
	if aup.Location != nil {
		objectMap["location"] = aup.Location
	}
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountPropertiesUpdateParameters AccountPropertiesUpdateParameters
				err = json.Unmarshal(*v, &accountPropertiesUpdateParameters)
				if err != nil {
					return err
				}
				aup.AccountPropertiesUpdateParameters = &accountPropertiesUpdateParameters
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				aup.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				aup.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				aup.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				aup.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				aup.Tags = tags
			}
		}
	}

	return nil
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - Gets a boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken or invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - Gets the reason that a storage account name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'AccountNameInvalid', 'AlreadyExists'
	Reason Reason `json:"reason,omitempty"`
	// Message - Gets an error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// CustomDomain the custom domain assigned to this storage account. This can be set via Update.
type CustomDomain struct {
	// Name - Gets or sets the custom domain name. Name is the CNAME source.
	Name *string `json:"name,omitempty"`
	// UseSubDomain - Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates
	UseSubDomain *bool `json:"useSubDomain,omitempty"`
}

// Endpoints the URIs that are used to perform a retrieval of a public blob, queue or table object.
type Endpoints struct {
	// Blob - Gets the blob endpoint.
	Blob *string `json:"blob,omitempty"`
	// Queue - Gets the queue endpoint.
	Queue *string `json:"queue,omitempty"`
	// Table - Gets the table endpoint.
	Table *string `json:"table,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// Usage describes Storage Resource Usage.
type Usage struct {
	// Unit - Gets the unit of measurement. Possible values include: 'Count', 'Bytes', 'Seconds', 'Percent', 'CountsPerSecond', 'BytesPerSecond'
	Unit UsageUnit `json:"unit,omitempty"`
	// CurrentValue - Gets the current count of the allocated resources in the subscription.
	CurrentValue *int32 `json:"currentValue,omitempty"`
	// Limit - Gets the maximum count of the resources that can be allocated in the subscription.
	Limit *int32 `json:"limit,omitempty"`
	// Name - Gets the name of the type of usage.
	Name *UsageName `json:"name,omitempty"`
}

// UsageListResult the List Usages operation response.
type UsageListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets the list Storage Resource Usages.
	Value *[]Usage `json:"value,omitempty"`
}

// UsageName the Usage Names.
type UsageName struct {
	// Value - Gets a string describing the resource name.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - Gets a localized string describing the resource name.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}
