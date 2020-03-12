package support

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/support/mgmt/2020-04-01/support"

// CommunicationDirection enumerates the values for communication direction.
type CommunicationDirection string

const (
	// Inbound ...
	Inbound CommunicationDirection = "inbound"
	// Outbound ...
	Outbound CommunicationDirection = "outbound"
)

// PossibleCommunicationDirectionValues returns an array of possible values for the CommunicationDirection const type.
func PossibleCommunicationDirectionValues() []CommunicationDirection {
	return []CommunicationDirection{Inbound, Outbound}
}

// CommunicationType enumerates the values for communication type.
type CommunicationType string

const (
	// Phone ...
	Phone CommunicationType = "phone"
	// Web ...
	Web CommunicationType = "web"
)

// PossibleCommunicationTypeValues returns an array of possible values for the CommunicationType const type.
func PossibleCommunicationTypeValues() []CommunicationType {
	return []CommunicationType{Phone, Web}
}

// PreferredContactMethod enumerates the values for preferred contact method.
type PreferredContactMethod string

const (
	// PreferredContactMethodEmail ...
	PreferredContactMethodEmail PreferredContactMethod = "email"
	// PreferredContactMethodPhone ...
	PreferredContactMethodPhone PreferredContactMethod = "phone"
)

// PossiblePreferredContactMethodValues returns an array of possible values for the PreferredContactMethod const type.
func PossiblePreferredContactMethodValues() []PreferredContactMethod {
	return []PreferredContactMethod{PreferredContactMethodEmail, PreferredContactMethodPhone}
}

// SeverityLevel enumerates the values for severity level.
type SeverityLevel string

const (
	// Critical ...
	Critical SeverityLevel = "critical"
	// Highestcriticalimpact ...
	Highestcriticalimpact SeverityLevel = "highestcriticalimpact"
	// Minimal ...
	Minimal SeverityLevel = "minimal"
	// Moderate ...
	Moderate SeverityLevel = "moderate"
)

// PossibleSeverityLevelValues returns an array of possible values for the SeverityLevel const type.
func PossibleSeverityLevelValues() []SeverityLevel {
	return []SeverityLevel{Critical, Highestcriticalimpact, Minimal, Moderate}
}

// Status enumerates the values for status.
type Status string

const (
	// Closed ...
	Closed Status = "closed"
	// Open ...
	Open Status = "open"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Closed, Open}
}

// Type enumerates the values for type.
type Type string

const (
	// MicrosoftSupportcommunications ...
	MicrosoftSupportcommunications Type = "Microsoft.Support/communications"
	// MicrosoftSupportsupportTickets ...
	MicrosoftSupportsupportTickets Type = "Microsoft.Support/supportTickets"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{MicrosoftSupportcommunications, MicrosoftSupportsupportTickets}
}

// CheckNameAvailabilityInput input of CheckNameAvailability API.
type CheckNameAvailabilityInput struct {
	// Name - The resource name to validate
	Name *string `json:"name,omitempty"`
	// Type - The type of resource. Possible values include: 'MicrosoftSupportsupportTickets', 'MicrosoftSupportcommunications'
	Type Type `json:"type,omitempty"`
}

// CheckNameAvailabilityOutput output of check name availability API.
type CheckNameAvailabilityOutput struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; Indicates whether the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason why the name is not available.
	Reason *string `json:"reason,omitempty"`
	// Message - READ-ONLY; The detailed error message describing why the name is not available.
	Message *string `json:"message,omitempty"`
}

// CommunicationDetails object that represents Communication resource
type CommunicationDetails struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Id of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource 'Microsoft.Support/communications'
	Type *string `json:"type,omitempty"`
	// CommunicationDetailsProperties - Properties of the resource
	*CommunicationDetailsProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for CommunicationDetails.
func (cd CommunicationDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cd.CommunicationDetailsProperties != nil {
		objectMap["properties"] = cd.CommunicationDetailsProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CommunicationDetails struct.
func (cd *CommunicationDetails) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cd.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cd.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cd.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var communicationDetailsProperties CommunicationDetailsProperties
				err = json.Unmarshal(*v, &communicationDetailsProperties)
				if err != nil {
					return err
				}
				cd.CommunicationDetailsProperties = &communicationDetailsProperties
			}
		}
	}

	return nil
}

// CommunicationDetailsProperties describes the properties of a communication resource.
type CommunicationDetailsProperties struct {
	// CommunicationType - READ-ONLY; Communication type. Possible values include: 'Web', 'Phone'
	CommunicationType CommunicationType `json:"communicationType,omitempty"`
	// CommunicationDirection - READ-ONLY; Direction of communication. Possible values include: 'Inbound', 'Outbound'
	CommunicationDirection CommunicationDirection `json:"communicationDirection,omitempty"`
	// Sender - Email address of the sender. This property is required if called by a service principal
	Sender *string `json:"sender,omitempty"`
	// Subject - Subject of the communication
	Subject *string `json:"subject,omitempty"`
	// Body - Body of the communication
	Body *string `json:"body,omitempty"`
	// CreatedDate - READ-ONLY; Time in UTC (ISO 8601 format) when the communication was created.
	CreatedDate *date.Time `json:"createdDate,omitempty"`
}

// CommunicationsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CommunicationsCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CommunicationsCreateFuture) Result(client CommunicationsClient) (cd CommunicationDetails, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.CommunicationsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("support.CommunicationsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cd.Response.Response, err = future.GetResult(sender); err == nil && cd.Response.Response.StatusCode != http.StatusNoContent {
		cd, err = client.CreateResponder(cd.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "support.CommunicationsCreateFuture", "Result", cd.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CommunicationsListResult collection of Communication resources.
type CommunicationsListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Communication resources.
	Value *[]CommunicationDetails `json:"value,omitempty"`
	// NextLink - The URI to fetch the next page of Communication resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// CommunicationsListResultIterator provides access to a complete listing of CommunicationDetails values.
type CommunicationsListResultIterator struct {
	i    int
	page CommunicationsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *CommunicationsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsListResultIterator.NextWithContext")
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
func (iter *CommunicationsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter CommunicationsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter CommunicationsListResultIterator) Response() CommunicationsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter CommunicationsListResultIterator) Value() CommunicationDetails {
	if !iter.page.NotDone() {
		return CommunicationDetails{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the CommunicationsListResultIterator type.
func NewCommunicationsListResultIterator(page CommunicationsListResultPage) CommunicationsListResultIterator {
	return CommunicationsListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (clr CommunicationsListResult) IsEmpty() bool {
	return clr.Value == nil || len(*clr.Value) == 0
}

// communicationsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (clr CommunicationsListResult) communicationsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if clr.NextLink == nil || len(to.String(clr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(clr.NextLink)))
}

// CommunicationsListResultPage contains a page of CommunicationDetails values.
type CommunicationsListResultPage struct {
	fn  func(context.Context, CommunicationsListResult) (CommunicationsListResult, error)
	clr CommunicationsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *CommunicationsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommunicationsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.clr)
	if err != nil {
		return err
	}
	page.clr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *CommunicationsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page CommunicationsListResultPage) NotDone() bool {
	return !page.clr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page CommunicationsListResultPage) Response() CommunicationsListResult {
	return page.clr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page CommunicationsListResultPage) Values() []CommunicationDetails {
	if page.clr.IsEmpty() {
		return nil
	}
	return *page.clr.Value
}

// Creates a new instance of the CommunicationsListResultPage type.
func NewCommunicationsListResultPage(getNextPage func(context.Context, CommunicationsListResult) (CommunicationsListResult, error)) CommunicationsListResultPage {
	return CommunicationsListResultPage{fn: getNextPage}
}

// ContactProfile contact information associated with support ticket.
type ContactProfile struct {
	// FirstName - First name.
	FirstName *string `json:"firstName,omitempty"`
	// LastName - Last name.
	LastName *string `json:"lastName,omitempty"`
	// PreferredContactMethod - Preferred contact method. Possible values include: 'PreferredContactMethodEmail', 'PreferredContactMethodPhone'
	PreferredContactMethod PreferredContactMethod `json:"preferredContactMethod,omitempty"`
	// PrimaryEmailAddress - Primary email address.
	PrimaryEmailAddress *string `json:"primaryEmailAddress,omitempty"`
	// AdditionalEmailAddresses - Additional email addresses listed will be copied on any correspondence about the support ticket.
	AdditionalEmailAddresses *[]string `json:"additionalEmailAddresses,omitempty"`
	// PhoneNumber - Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// PreferredTimeZone - Time zone of the user. This is the name of the time zone from [Microsoft Time Zone Index Values](https://support.microsoft.com/help/973627/microsoft-time-zone-index-values).
	PreferredTimeZone *string `json:"preferredTimeZone,omitempty"`
	// Country - Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string `json:"country,omitempty"`
	// PreferredSupportLanguage - Preferred language of support from Azure. Support languages vary based on the severity you choose for your support ticket. Learn more at [Azure Severity and responsiveness](https://azure.microsoft.com/support/plans/response). Use the standard language-country code. Valid values are 'en-us' for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French, 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for Chinese and 'de-de' for German.
	PreferredSupportLanguage *string `json:"preferredSupportLanguage,omitempty"`
}

// Engineer support engineer information.
type Engineer struct {
	// EmailAddress - READ-ONLY; Email address of the Azure Support engineer assigned to the support ticket.
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// ExceptionResponse the api error.
type ExceptionResponse struct {
	// Error - The api error details.
	Error *ServiceError `json:"error,omitempty"`
}

// Operation the operation supported by Microsoft Support RP.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that describes the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that describes the operation.
type OperationDisplay struct {
	// Description - READ-ONLY; The description of the operation
	Description *string `json:"description,omitempty"`
	// Operation - READ-ONLY; The action that users can perform, based on their permission level
	Operation *string `json:"operation,omitempty"`
	// Provider - READ-ONLY; Service provider: Microsoft Support
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed
	Resource *string `json:"resource,omitempty"`
}

// OperationsListResult the list of operations supported by Microsoft Support resource provider.
type OperationsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of operations supported by Microsoft Support resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// ProblemClassification problemClassification resource object
type ProblemClassification struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Id of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource 'Microsoft.Support/problemClassification'
	Type *string `json:"type,omitempty"`
	// ProblemClassificationProperties - Properties of the resource
	*ProblemClassificationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ProblemClassification.
func (pc ProblemClassification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pc.ProblemClassificationProperties != nil {
		objectMap["properties"] = pc.ProblemClassificationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ProblemClassification struct.
func (pc *ProblemClassification) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				pc.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var problemClassificationProperties ProblemClassificationProperties
				err = json.Unmarshal(*v, &problemClassificationProperties)
				if err != nil {
					return err
				}
				pc.ProblemClassificationProperties = &problemClassificationProperties
			}
		}
	}

	return nil
}

// ProblemClassificationProperties details about a problem classification available for an Azure service
type ProblemClassificationProperties struct {
	// DisplayName - Localized name of problem classification.
	DisplayName *string `json:"displayName,omitempty"`
}

// ProblemClassificationsListResult collection of ProblemClassification resources
type ProblemClassificationsListResult struct {
	autorest.Response `json:"-"`
	// Value - List of ProblemClassification resources
	Value *[]ProblemClassification `json:"value,omitempty"`
}

// QuotaChangeRequest this property is required for providing the region and new quota limits
type QuotaChangeRequest struct {
	// Region - Region for which the quota increase request is being made.
	Region *string `json:"region,omitempty"`
	// Payload - Payload of the quota increase request.
	Payload *string `json:"payload,omitempty"`
}

// QuotaTicketDetails additional set of information required for quota increase support ticket for certain
// quota types, e.g.: Virtual machine cores. Get complete details about Quota payload support request along
// with examples at [Support quota request](https://aka.ms/supportrpquotarequestpayload).
type QuotaTicketDetails struct {
	// QuotaChangeRequestSubType - Required for certain quota types when there is a sub type that you are requesting quota increase for. Example: Batch
	QuotaChangeRequestSubType *string `json:"quotaChangeRequestSubType,omitempty"`
	// QuotaChangeRequestVersion - Quota change request version
	QuotaChangeRequestVersion *string `json:"quotaChangeRequestVersion,omitempty"`
	// QuotaChangeRequests - This property is required for providing the region and new quota limits.
	QuotaChangeRequests *[]QuotaChangeRequest `json:"quotaChangeRequests,omitempty"`
}

// Service object that represents a Service resource.
type Service struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Id of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource 'Microsoft.Support/services'
	Type *string `json:"type,omitempty"`
	// ServiceProperties - Properties of the resource
	*ServiceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Service.
func (s Service) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.ServiceProperties != nil {
		objectMap["properties"] = s.ServiceProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Service struct.
func (s *Service) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				s.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				s.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				s.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var serviceProperties ServiceProperties
				err = json.Unmarshal(*v, &serviceProperties)
				if err != nil {
					return err
				}
				s.ServiceProperties = &serviceProperties
			}
		}
	}

	return nil
}

// ServiceError the api error details.
type ServiceError struct {
	// Code - The error code.
	Code *string `json:"code,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
	// Target - The target of the error.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The list of error details.
	Details *[]ServiceErrorDetail `json:"details,omitempty"`
}

// ServiceErrorDetail the error details.
type ServiceErrorDetail struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - The target of the error.
	Target *string `json:"target,omitempty"`
}

// ServiceLevelAgreement service Level Agreement details for a support ticket.
type ServiceLevelAgreement struct {
	// StartTime - READ-ONLY; Time in UTC (ISO 8601 format) when service level agreement starts.
	StartTime *date.Time `json:"startTime,omitempty"`
	// ExpirationTime - READ-ONLY; Time in UTC (ISO 8601 format) when service level agreement expires.
	ExpirationTime *date.Time `json:"expirationTime,omitempty"`
	// SLAMinutes - READ-ONLY; Service Level Agreement in minutes
	SLAMinutes *int32 `json:"slaMinutes,omitempty"`
}

// ServiceProperties details about Azure service available for support ticket creation
type ServiceProperties struct {
	// DisplayName - Localized name of Azure service
	DisplayName *string `json:"displayName,omitempty"`
	// ResourceTypes - ARM Resource types
	ResourceTypes *[]string `json:"resourceTypes,omitempty"`
}

// ServicesListResult collection of Service resources.
type ServicesListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Service resources
	Value *[]Service `json:"value,omitempty"`
}

// TechnicalTicketDetails additional information for technical support ticket.
type TechnicalTicketDetails struct {
	// ResourceID - This is the resource id of the Azure service resource (For example: A virtual machine resource or an HDInsight resource) for which the support ticket is created.
	ResourceID *string `json:"resourceId,omitempty"`
}

// TicketDetails object that represents SupportTicketDetails resource
type TicketDetails struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Id of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource 'Microsoft.Support/supportTickets'
	Type *string `json:"type,omitempty"`
	// TicketDetailsProperties - Properties of the resource
	*TicketDetailsProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for TicketDetails.
func (td TicketDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if td.TicketDetailsProperties != nil {
		objectMap["properties"] = td.TicketDetailsProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for TicketDetails struct.
func (td *TicketDetails) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				td.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				td.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				td.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var ticketDetailsProperties TicketDetailsProperties
				err = json.Unmarshal(*v, &ticketDetailsProperties)
				if err != nil {
					return err
				}
				td.TicketDetailsProperties = &ticketDetailsProperties
			}
		}
	}

	return nil
}

// TicketDetailsProperties describes the properties of a support ticket.
type TicketDetailsProperties struct {
	// SupportTicketID - System generated support ticket id that is unique.
	SupportTicketID *string `json:"supportTicketId,omitempty"`
	// Description - Detailed description of the question or issue.
	Description *string `json:"description,omitempty"`
	// ProblemClassificationID - Each Azure service has its own set of issue category called problem classification that corresponds to the type of problem you're experiencing. This parameter is the resource id of ProblemClassification resource.
	ProblemClassificationID *string `json:"problemClassificationId,omitempty"`
	// ProblemClassificationDisplayName - READ-ONLY; Localized name of problem classification.
	ProblemClassificationDisplayName *string `json:"problemClassificationDisplayName,omitempty"`
	// Severity - A value that indicates the urgency of the case, which in turn determines the response time according to the service level agreement of the technical support plan you have with Azure. Note: 'Highest critical impact' severity is reserved only to our Premium customers. Possible values include: 'Minimal', 'Moderate', 'Critical', 'Highestcriticalimpact'
	Severity SeverityLevel `json:"severity,omitempty"`
	// EnrollmentID - READ-ONLY; Enrollment ID associated with the support ticket.
	EnrollmentID *string `json:"enrollmentId,omitempty"`
	// Require24X7Response - Indicates if this requires a 24x7 response from Azure.
	Require24X7Response *bool `json:"require24X7Response,omitempty"`
	// ContactDetails - Contact information of the user requesting to create a support ticket.
	ContactDetails *ContactProfile `json:"contactDetails,omitempty"`
	// ServiceLevelAgreement - Service Level Agreement information for this support ticket.
	ServiceLevelAgreement *ServiceLevelAgreement `json:"serviceLevelAgreement,omitempty"`
	// SupportEngineer - Information about support engineer working on this support ticket.
	SupportEngineer *Engineer `json:"supportEngineer,omitempty"`
	// SupportPlanType - READ-ONLY; Support plan type associated with the support ticket.
	SupportPlanType *string `json:"supportPlanType,omitempty"`
	// Title - Title of the support ticket.
	Title *string `json:"title,omitempty"`
	// ProblemStartTime - Time in UTC (ISO 8601 format) when the problem started.
	ProblemStartTime *date.Time `json:"problemStartTime,omitempty"`
	// ServiceID - This is the resource id of the Azure service resource associated with the support ticket.
	ServiceID *string `json:"serviceId,omitempty"`
	// ServiceDisplayName - READ-ONLY; Localized name of Azure service.
	ServiceDisplayName *string `json:"serviceDisplayName,omitempty"`
	// Status - READ-ONLY; Status of the support ticket.
	Status *string `json:"status,omitempty"`
	// CreatedDate - READ-ONLY; Time in UTC (ISO 8601 format) when support ticket was created.
	CreatedDate *date.Time `json:"createdDate,omitempty"`
	// ModifiedDate - READ-ONLY; Time in UTC (ISO 8601 format) when support ticket was last modified.
	ModifiedDate *date.Time `json:"modifiedDate,omitempty"`
	// TechnicalTicketDetails - Additional ticket details associated with a technical support ticket request.
	TechnicalTicketDetails *TechnicalTicketDetails `json:"technicalTicketDetails,omitempty"`
	// QuotaTicketDetails - Additional ticket details associated with a quota support ticket request.
	QuotaTicketDetails *QuotaTicketDetails `json:"quotaTicketDetails,omitempty"`
}

// TicketsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type TicketsCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *TicketsCreateFuture) Result(client TicketsClient) (td TicketDetails, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "support.TicketsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("support.TicketsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if td.Response.Response, err = future.GetResult(sender); err == nil && td.Response.Response.StatusCode != http.StatusNoContent {
		td, err = client.CreateResponder(td.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "support.TicketsCreateFuture", "Result", td.Response.Response, "Failure responding to request")
		}
	}
	return
}

// TicketsListResult object that represents a collection of SupportTicket resources.
type TicketsListResult struct {
	autorest.Response `json:"-"`
	// Value - List of SupportTicket resources.
	Value *[]TicketDetails `json:"value,omitempty"`
	// NextLink - The URI to fetch the next page of SupportTicket resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// TicketsListResultIterator provides access to a complete listing of TicketDetails values.
type TicketsListResultIterator struct {
	i    int
	page TicketsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *TicketsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsListResultIterator.NextWithContext")
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
func (iter *TicketsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter TicketsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter TicketsListResultIterator) Response() TicketsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter TicketsListResultIterator) Value() TicketDetails {
	if !iter.page.NotDone() {
		return TicketDetails{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the TicketsListResultIterator type.
func NewTicketsListResultIterator(page TicketsListResultPage) TicketsListResultIterator {
	return TicketsListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (tlr TicketsListResult) IsEmpty() bool {
	return tlr.Value == nil || len(*tlr.Value) == 0
}

// ticketsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (tlr TicketsListResult) ticketsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if tlr.NextLink == nil || len(to.String(tlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(tlr.NextLink)))
}

// TicketsListResultPage contains a page of TicketDetails values.
type TicketsListResultPage struct {
	fn  func(context.Context, TicketsListResult) (TicketsListResult, error)
	tlr TicketsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *TicketsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TicketsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.tlr)
	if err != nil {
		return err
	}
	page.tlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *TicketsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page TicketsListResultPage) NotDone() bool {
	return !page.tlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page TicketsListResultPage) Response() TicketsListResult {
	return page.tlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page TicketsListResultPage) Values() []TicketDetails {
	if page.tlr.IsEmpty() {
		return nil
	}
	return *page.tlr.Value
}

// Creates a new instance of the TicketsListResultPage type.
func NewTicketsListResultPage(getNextPage func(context.Context, TicketsListResult) (TicketsListResult, error)) TicketsListResultPage {
	return TicketsListResultPage{fn: getNextPage}
}

// UpdateContactProfile contact information associated with the support ticket.
type UpdateContactProfile struct {
	// FirstName - First name
	FirstName *string `json:"firstName,omitempty"`
	// LastName - Last name
	LastName *string `json:"lastName,omitempty"`
	// PreferredContactMethod - Preferred contact method. Possible values include: 'PreferredContactMethodEmail', 'PreferredContactMethodPhone'
	PreferredContactMethod PreferredContactMethod `json:"preferredContactMethod,omitempty"`
	// PrimaryEmailAddress - Primary email address
	PrimaryEmailAddress *string `json:"primaryEmailAddress,omitempty"`
	// AdditionalEmailAddresses - Email addresses listed will be copied on any correspondence about the support ticket
	AdditionalEmailAddresses *[]string `json:"additionalEmailAddresses,omitempty"`
	// PhoneNumber - Phone number. This is required if preferred contact method is phone.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// PreferredTimeZone - Time zone of the user. This is the name of the time zone from [Microsoft Time Zone Index Values](https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values).
	PreferredTimeZone *string `json:"preferredTimeZone,omitempty"`
	// Country - Country of the user. This is the ISO 3166-1 alpha-3 code.
	Country *string `json:"country,omitempty"`
	// PreferredSupportLanguage - Preferred language of support from Azure. Support languages vary based on the severity you choose for your support ticket. Learn more at [Azure Severity and responsiveness](https://azure.microsoft.com/support/plans/response/). Use the standard language-country code. Valid values are 'en-us' for English, 'zh-hans' for Chinese, 'es-es' for Spanish, 'fr-fr' for French, 'ja-jp' for Japanese, 'ko-kr' for Korean, 'ru-ru' for Russian, 'pt-br' for Portuguese, 'it-it' for Italian, 'zh-tw' for Chinese and 'de-de' for German.
	PreferredSupportLanguage *string `json:"preferredSupportLanguage,omitempty"`
}

// UpdateSupportTicket updates severity, ticket status and contact details in the support ticket.
type UpdateSupportTicket struct {
	// Severity - Severity level. Possible values include: 'Minimal', 'Moderate', 'Critical', 'Highestcriticalimpact'
	Severity SeverityLevel `json:"severity,omitempty"`
	// Status - Status to be updated on the ticket. Possible values include: 'Open', 'Closed'
	Status Status `json:"status,omitempty"`
	// ContactDetails - Contact details to be updated on the support ticket.
	ContactDetails *UpdateContactProfile `json:"contactDetails,omitempty"`
}
