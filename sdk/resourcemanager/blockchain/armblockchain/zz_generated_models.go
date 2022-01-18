//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// APIKey - API key payload which is exposed in the request/response of the resource provider.
type APIKey struct {
	// Gets or sets the API key name.
	KeyName *string `json:"keyName,omitempty"`

	// Gets or sets the API key value.
	Value *string `json:"value,omitempty"`
}

// APIKeyCollection - Collection of the API key payload which is exposed in the response of the resource provider.
type APIKeyCollection struct {
	// Gets or sets the collection of API key.
	Keys []*APIKey `json:"keys,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type APIKeyCollection.
func (a APIKeyCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "keys", a.Keys)
	return json.Marshal(objectMap)
}

// Consortium payload
type Consortium struct {
	// Gets or sets the blockchain member name.
	Name *string `json:"name,omitempty"`

	// Gets or sets the protocol for the consortium.
	Protocol *BlockchainProtocol `json:"protocol,omitempty"`
}

// ConsortiumCollection - Collection of the consortium payload.
type ConsortiumCollection struct {
	// Gets or sets the collection of consortiums.
	Value []*Consortium `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConsortiumCollection.
func (c ConsortiumCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ConsortiumMember - Consortium approval
type ConsortiumMember struct {
	// Gets the consortium member modified date.
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Gets the consortium member display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Gets the consortium member join date.
	JoinDate *time.Time `json:"joinDate,omitempty"`

	// Gets the consortium member name.
	Name *string `json:"name,omitempty"`

	// Gets the consortium member role.
	Role *string `json:"role,omitempty"`

	// Gets the consortium member status.
	Status *string `json:"status,omitempty"`

	// Gets the consortium member subscription id.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConsortiumMember.
func (c ConsortiumMember) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "dateModified", c.DateModified)
	populate(objectMap, "displayName", c.DisplayName)
	populateTimeRFC3339(objectMap, "joinDate", c.JoinDate)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "role", c.Role)
	populate(objectMap, "status", c.Status)
	populate(objectMap, "subscriptionId", c.SubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConsortiumMember.
func (c *ConsortiumMember) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dateModified":
			err = unpopulateTimeRFC3339(val, &c.DateModified)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &c.DisplayName)
			delete(rawMsg, key)
		case "joinDate":
			err = unpopulateTimeRFC3339(val, &c.JoinDate)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &c.Name)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, &c.Role)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &c.Status)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, &c.SubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConsortiumMemberCollection - Collection of consortium payload.
type ConsortiumMemberCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets the collection of consortiums.
	Value []*ConsortiumMember `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConsortiumMemberCollection.
func (c ConsortiumMemberCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// FirewallRule - Ip range for firewall rules
type FirewallRule struct {
	// Gets or sets the end IP address of the firewall rule range.
	EndIPAddress *string `json:"endIpAddress,omitempty"`

	// Gets or sets the name of the firewall rules.
	RuleName *string `json:"ruleName,omitempty"`

	// Gets or sets the start IP address of the firewall rule range.
	StartIPAddress *string `json:"startIpAddress,omitempty"`
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// Name availability request payload.
	NameAvailabilityRequest *NameAvailabilityRequest
}

// LocationsClientListConsortiumsOptions contains the optional parameters for the LocationsClient.ListConsortiums method.
type LocationsClientListConsortiumsOptions struct {
	// placeholder for future optional parameters
}

// Member - Payload of the blockchain member which is exposed in the request/response of the resource provider.
type Member struct {
	// The GEO location of the blockchain service.
	Location *string `json:"location,omitempty"`

	// Gets or sets the blockchain member properties.
	Properties *MemberProperties `json:"properties,omitempty"`

	// Gets or sets the blockchain member Sku.
	SKU *SKU `json:"sku,omitempty"`

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Member.
func (m Member) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MemberCollection - Collection of the blockchain member payload which is exposed in the request/response of the resource
// provider.
type MemberCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets the collection of blockchain members.
	Value []*Member `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MemberCollection.
func (m MemberCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MemberNodesSKU - Payload of the blockchain member nodes Sku for a blockchain member.
type MemberNodesSKU struct {
	// Gets or sets the nodes capacity.
	Capacity *int32 `json:"capacity,omitempty"`
}

// MemberOperationResultsClientGetOptions contains the optional parameters for the MemberOperationResultsClient.Get method.
type MemberOperationResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MemberProperties - Payload of the blockchain member properties for a blockchain member.
type MemberProperties struct {
	// Gets or sets the consortium for the blockchain member.
	Consortium *string `json:"consortium,omitempty"`

	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string `json:"consortiumManagementAccountPassword,omitempty"`

	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName *string `json:"consortiumMemberDisplayName,omitempty"`

	// Gets the role of the member in the consortium.
	ConsortiumRole *string `json:"consortiumRole,omitempty"`

	// Gets or sets firewall rules
	FirewallRules []*FirewallRule `json:"firewallRules,omitempty"`

	// Sets the basic auth password of the blockchain member.
	Password *string `json:"password,omitempty"`

	// Gets or sets the blockchain protocol.
	Protocol *BlockchainProtocol `json:"protocol,omitempty"`

	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSKU *MemberNodesSKU `json:"validatorNodesSku,omitempty"`

	// READ-ONLY; Gets the managed consortium management account address.
	ConsortiumManagementAccountAddress *string `json:"consortiumManagementAccountAddress,omitempty" azure:"ro"`

	// READ-ONLY; Gets the dns endpoint of the blockchain member.
	DNS *string `json:"dns,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the blockchain member provision state.
	ProvisioningState *BlockchainMemberProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Gets the public key of the blockchain member (default transaction node).
	PublicKey *string `json:"publicKey,omitempty" azure:"ro"`

	// READ-ONLY; Gets the Ethereum root contract address of the blockchain.
	RootContractAddress *string `json:"rootContractAddress,omitempty" azure:"ro"`

	// READ-ONLY; Gets the auth user name of the blockchain member.
	UserName *string `json:"userName,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type MemberProperties.
func (m MemberProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "consortium", m.Consortium)
	populate(objectMap, "consortiumManagementAccountAddress", m.ConsortiumManagementAccountAddress)
	populate(objectMap, "consortiumManagementAccountPassword", m.ConsortiumManagementAccountPassword)
	populate(objectMap, "consortiumMemberDisplayName", m.ConsortiumMemberDisplayName)
	populate(objectMap, "consortiumRole", m.ConsortiumRole)
	populate(objectMap, "dns", m.DNS)
	populate(objectMap, "firewallRules", m.FirewallRules)
	populate(objectMap, "password", m.Password)
	populate(objectMap, "protocol", m.Protocol)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "publicKey", m.PublicKey)
	populate(objectMap, "rootContractAddress", m.RootContractAddress)
	populate(objectMap, "userName", m.UserName)
	populate(objectMap, "validatorNodesSku", m.ValidatorNodesSKU)
	return json.Marshal(objectMap)
}

// MemberPropertiesUpdate - Update the payload of the blockchain member properties for a blockchain member.
type MemberPropertiesUpdate struct {
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string `json:"consortiumManagementAccountPassword,omitempty"`

	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule `json:"firewallRules,omitempty"`

	// Sets the transaction node dns endpoint basic auth password.
	Password *string `json:"password,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MemberPropertiesUpdate.
func (m MemberPropertiesUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "consortiumManagementAccountPassword", m.ConsortiumManagementAccountPassword)
	populate(objectMap, "firewallRules", m.FirewallRules)
	populate(objectMap, "password", m.Password)
	return json.Marshal(objectMap)
}

// MemberUpdate - Update the payload of the blockchain member which is exposed in the request/response of the resource provider.
type MemberUpdate struct {
	// Gets or sets the blockchain member update properties.
	Properties *MemberPropertiesUpdate `json:"properties,omitempty"`

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MemberUpdate.
func (m MemberUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	return json.Marshal(objectMap)
}

// MembersClientBeginCreateOptions contains the optional parameters for the MembersClient.BeginCreate method.
type MembersClientBeginCreateOptions struct {
	// Payload to create a blockchain member.
	BlockchainMember *Member
}

// MembersClientBeginDeleteOptions contains the optional parameters for the MembersClient.BeginDelete method.
type MembersClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// MembersClientGetOptions contains the optional parameters for the MembersClient.Get method.
type MembersClientGetOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListAPIKeysOptions contains the optional parameters for the MembersClient.ListAPIKeys method.
type MembersClientListAPIKeysOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListAllOptions contains the optional parameters for the MembersClient.ListAll method.
type MembersClientListAllOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListConsortiumMembersOptions contains the optional parameters for the MembersClient.ListConsortiumMembers
// method.
type MembersClientListConsortiumMembersOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListOptions contains the optional parameters for the MembersClient.List method.
type MembersClientListOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListRegenerateAPIKeysOptions contains the optional parameters for the MembersClient.ListRegenerateAPIKeys
// method.
type MembersClientListRegenerateAPIKeysOptions struct {
	// api key to be regenerate
	APIKey *APIKey
}

// MembersClientUpdateOptions contains the optional parameters for the MembersClient.Update method.
type MembersClientUpdateOptions struct {
	// Payload to update the blockchain member.
	BlockchainMember *MemberUpdate
}

// NameAvailability - Name availability payload which is exposed in the response of the resource provider.
type NameAvailability struct {
	// Gets or sets the message.
	Message *string `json:"message,omitempty"`

	// Gets or sets the value indicating whether the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// Gets or sets the name availability reason.
	Reason *NameAvailabilityReason `json:"reason,omitempty"`
}

// NameAvailabilityRequest - Name availability request payload which is exposed in the request of the resource provider.
type NameAvailabilityRequest struct {
	// Gets or sets the name to check.
	Name *string `json:"name,omitempty"`

	// Gets or sets the type of the resource to check.
	Type *string `json:"type,omitempty"`
}

// OperationResult - Operation result payload which is exposed in the response of the resource provider.
type OperationResult struct {
	// Gets or sets the operation end time.
	EndTime *time.Time `json:"endTime,omitempty"`

	// Gets or sets the operation name.
	Name *string `json:"name,omitempty"`

	// Gets or sets the operation start time.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationResult.
func (o OperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", o.EndTime)
	populate(objectMap, "name", o.Name)
	populateTimeRFC3339(objectMap, "startTime", o.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResult.
func (o *OperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, &o.EndTime)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &o.Name)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &o.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - The core properties of the resources.
type Resource struct {
	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceProviderOperation - Operation payload which is exposed in the response of the resource provider.
type ResourceProviderOperation struct {
	// Gets or sets operation display
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`

	// Gets or sets a value indicating whether the operation is a data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Gets or sets the operation name.
	Name *string `json:"name,omitempty"`

	// Gets or sets the origin.
	Origin *string `json:"origin,omitempty"`
}

// ResourceProviderOperationCollection - Collection of operation payload which is exposed in the response of the resource
// provider.
type ResourceProviderOperationCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets the collection of operations.
	Value []*ResourceProviderOperation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderOperationCollection.
func (r ResourceProviderOperationCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// ResourceProviderOperationDisplay - Operation display payload which is exposed in the response of the resource provider.
type ResourceProviderOperationDisplay struct {
	// Gets or sets the description of the provider for display purposes.
	Description *string `json:"description,omitempty"`

	// Gets or sets the name of the operation for display purposes.
	Operation *string `json:"operation,omitempty"`

	// Gets or sets the name of the provider for display purposes.
	Provider *string `json:"provider,omitempty"`

	// Gets or sets the name of the resource type for display purposes.
	Resource *string `json:"resource,omitempty"`
}

// ResourceTypeSKU - Resource type Sku.
type ResourceTypeSKU struct {
	// Gets or sets the resource type
	ResourceType *string `json:"resourceType,omitempty"`

	// Gets or sets the Skus
	SKUs []*SKUSetting `json:"skus,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeSKU.
func (r ResourceTypeSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "skus", r.SKUs)
	return json.Marshal(objectMap)
}

// ResourceTypeSKUCollection - Collection of the resource type Sku.
type ResourceTypeSKUCollection struct {
	// Gets or sets the collection of resource type Sku.
	Value []*ResourceTypeSKU `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeSKUCollection.
func (r ResourceTypeSKUCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// SKU - Blockchain member Sku in payload
type SKU struct {
	// Gets or sets Sku name
	Name *string `json:"name,omitempty"`

	// Gets or sets Sku tier
	Tier *string `json:"tier,omitempty"`
}

// SKUSetting - Sku Setting.
type SKUSetting struct {
	// Gets or sets the locations.
	Locations []*string `json:"locations,omitempty"`

	// Gets or sets the Sku name.
	Name *string `json:"name,omitempty"`

	// Gets or sets the required features.
	RequiredFeatures []*string `json:"requiredFeatures,omitempty"`

	// Gets or sets the Sku tier.
	Tier *string `json:"tier,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SKUSetting.
func (s SKUSetting) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "locations", s.Locations)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "requiredFeatures", s.RequiredFeatures)
	populate(objectMap, "tier", s.Tier)
	return json.Marshal(objectMap)
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.List method.
type SKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// TrackedResource - The resource model definition for a top level resource.
type TrackedResource struct {
	// The GEO location of the blockchain service.
	Location *string `json:"location,omitempty"`

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// TransactionNode - Payload of the transaction node which is the request/response of the resource provider.
type TransactionNode struct {
	// Gets or sets the transaction node location.
	Location *string `json:"location,omitempty"`

	// Gets or sets the blockchain member properties.
	Properties *TransactionNodeProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TransactionNodeCollection - Collection of transaction node payload which is exposed in the request/response of the resource
// provider.
type TransactionNodeCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets the collection of transaction nodes.
	Value []*TransactionNode `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TransactionNodeCollection.
func (t TransactionNodeCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", t.NextLink)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// TransactionNodeProperties - Payload of transaction node properties payload in the transaction node payload.
type TransactionNodeProperties struct {
	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule `json:"firewallRules,omitempty"`

	// Sets the transaction node dns endpoint basic auth password.
	Password *string `json:"password,omitempty"`

	// READ-ONLY; Gets or sets the transaction node dns endpoint.
	DNS *string `json:"dns,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the blockchain member provision state.
	ProvisioningState *NodeProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the transaction node public key.
	PublicKey *string `json:"publicKey,omitempty" azure:"ro"`

	// READ-ONLY; Gets or sets the transaction node dns endpoint basic auth user name.
	UserName *string `json:"userName,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TransactionNodeProperties.
func (t TransactionNodeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dns", t.DNS)
	populate(objectMap, "firewallRules", t.FirewallRules)
	populate(objectMap, "password", t.Password)
	populate(objectMap, "provisioningState", t.ProvisioningState)
	populate(objectMap, "publicKey", t.PublicKey)
	populate(objectMap, "userName", t.UserName)
	return json.Marshal(objectMap)
}

// TransactionNodePropertiesUpdate - Update the payload of the transaction node properties in the transaction node payload.
type TransactionNodePropertiesUpdate struct {
	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule `json:"firewallRules,omitempty"`

	// Sets the transaction node dns endpoint basic auth password.
	Password *string `json:"password,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TransactionNodePropertiesUpdate.
func (t TransactionNodePropertiesUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "firewallRules", t.FirewallRules)
	populate(objectMap, "password", t.Password)
	return json.Marshal(objectMap)
}

// TransactionNodeUpdate - Update the transaction node payload which is exposed in the request/response of the resource provider.
type TransactionNodeUpdate struct {
	// Gets or sets the transaction node update properties.
	Properties *TransactionNodePropertiesUpdate `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TransactionNodeUpdate.
func (t TransactionNodeUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", t.Properties)
	return json.Marshal(objectMap)
}

// TransactionNodesClientBeginCreateOptions contains the optional parameters for the TransactionNodesClient.BeginCreate method.
type TransactionNodesClientBeginCreateOptions struct {
	// Payload to create the transaction node.
	TransactionNode *TransactionNode
}

// TransactionNodesClientBeginDeleteOptions contains the optional parameters for the TransactionNodesClient.BeginDelete method.
type TransactionNodesClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientGetOptions contains the optional parameters for the TransactionNodesClient.Get method.
type TransactionNodesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListAPIKeys method.
type TransactionNodesClientListAPIKeysOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListOptions contains the optional parameters for the TransactionNodesClient.List method.
type TransactionNodesClientListOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListRegenerateAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListRegenerateAPIKeys
// method.
type TransactionNodesClientListRegenerateAPIKeysOptions struct {
	// api key to be regenerated
	APIKey *APIKey
}

// TransactionNodesClientUpdateOptions contains the optional parameters for the TransactionNodesClient.Update method.
type TransactionNodesClientUpdateOptions struct {
	// Payload to create the transaction node.
	TransactionNode *TransactionNodeUpdate
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
