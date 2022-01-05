//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
	"time"
)

// ConnectedCluster - Represents a connected cluster.
type ConnectedCluster struct {
	TrackedResource
	// REQUIRED; The identity of the connected cluster.
	Identity *ConnectedClusterIdentity `json:"identity,omitempty"`

	// REQUIRED; Describes the connected cluster resource properties.
	Properties *ConnectedClusterProperties `json:"properties,omitempty"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedCluster.
func (c ConnectedCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	return json.Marshal(objectMap)
}

// ConnectedClusterBeginCreateOptions contains the optional parameters for the ConnectedCluster.BeginCreate method.
type ConnectedClusterBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterBeginDeleteOptions contains the optional parameters for the ConnectedCluster.BeginDelete method.
type ConnectedClusterBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterGetOptions contains the optional parameters for the ConnectedCluster.Get method.
type ConnectedClusterGetOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterIdentity - Identity for the connected cluster.
type ConnectedClusterIdentity struct {
	// REQUIRED; The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no
	// identity is assigned to the connected cluster.
	Type *ResourceIdentityType `json:"type,omitempty"`

	// READ-ONLY; The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ConnectedClusterList - The paginated list of connected Clusters
type ConnectedClusterList struct {
	// The link to fetch the next page of connected cluster
	NextLink *string `json:"nextLink,omitempty"`

	// The list of connected clusters
	Value []*ConnectedCluster `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterList.
func (c ConnectedClusterList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ConnectedClusterListByResourceGroupOptions contains the optional parameters for the ConnectedCluster.ListByResourceGroup method.
type ConnectedClusterListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterListBySubscriptionOptions contains the optional parameters for the ConnectedCluster.ListBySubscription method.
type ConnectedClusterListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterListClusterUserCredentialOptions contains the optional parameters for the ConnectedCluster.ListClusterUserCredential method.
type ConnectedClusterListClusterUserCredentialOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterPatch - Object containing updates for patch operations.
type ConnectedClusterPatch struct {
	// Describes the connected cluster resource properties that can be updated during PATCH operation.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterPatch.
func (c ConnectedClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// ConnectedClusterProperties - Properties of the connected cluster.
type ConnectedClusterProperties struct {
	// REQUIRED; Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate *string `json:"agentPublicKeyCertificate,omitempty"`

	// The Kubernetes distribution running on this connected cluster.
	Distribution *string `json:"distribution,omitempty"`

	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string `json:"infrastructure,omitempty"`

	// Provisioning state of the connected cluster resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// READ-ONLY; Version of the agent running on the connected cluster resource
	AgentVersion *string `json:"agentVersion,omitempty" azure:"ro"`

	// READ-ONLY; Represents the connectivity status of the connected cluster.
	ConnectivityStatus *ConnectivityStatus `json:"connectivityStatus,omitempty" azure:"ro"`

	// READ-ONLY; The Kubernetes version of the connected cluster resource
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" azure:"ro"`

	// READ-ONLY; Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime *time.Time `json:"lastConnectivityTime,omitempty" azure:"ro"`

	// READ-ONLY; Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime *time.Time `json:"managedIdentityCertificateExpirationTime,omitempty" azure:"ro"`

	// READ-ONLY; Connected cluster offering
	Offering *string `json:"offering,omitempty" azure:"ro"`

	// READ-ONLY; Number of CPU cores present in the connected cluster resource
	TotalCoreCount *int32 `json:"totalCoreCount,omitempty" azure:"ro"`

	// READ-ONLY; Number of nodes present in the connected cluster resource
	TotalNodeCount *int32 `json:"totalNodeCount,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ConnectedClusterProperties.
func (c ConnectedClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agentPublicKeyCertificate", c.AgentPublicKeyCertificate)
	populate(objectMap, "agentVersion", c.AgentVersion)
	populate(objectMap, "connectivityStatus", c.ConnectivityStatus)
	populate(objectMap, "distribution", c.Distribution)
	populate(objectMap, "infrastructure", c.Infrastructure)
	populate(objectMap, "kubernetesVersion", c.KubernetesVersion)
	populateTimeRFC3339(objectMap, "lastConnectivityTime", c.LastConnectivityTime)
	populateTimeRFC3339(objectMap, "managedIdentityCertificateExpirationTime", c.ManagedIdentityCertificateExpirationTime)
	populate(objectMap, "offering", c.Offering)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "totalCoreCount", c.TotalCoreCount)
	populate(objectMap, "totalNodeCount", c.TotalNodeCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectedClusterProperties.
func (c *ConnectedClusterProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "agentPublicKeyCertificate":
			err = unpopulate(val, &c.AgentPublicKeyCertificate)
			delete(rawMsg, key)
		case "agentVersion":
			err = unpopulate(val, &c.AgentVersion)
			delete(rawMsg, key)
		case "connectivityStatus":
			err = unpopulate(val, &c.ConnectivityStatus)
			delete(rawMsg, key)
		case "distribution":
			err = unpopulate(val, &c.Distribution)
			delete(rawMsg, key)
		case "infrastructure":
			err = unpopulate(val, &c.Infrastructure)
			delete(rawMsg, key)
		case "kubernetesVersion":
			err = unpopulate(val, &c.KubernetesVersion)
			delete(rawMsg, key)
		case "lastConnectivityTime":
			err = unpopulateTimeRFC3339(val, &c.LastConnectivityTime)
			delete(rawMsg, key)
		case "managedIdentityCertificateExpirationTime":
			err = unpopulateTimeRFC3339(val, &c.ManagedIdentityCertificateExpirationTime)
			delete(rawMsg, key)
		case "offering":
			err = unpopulate(val, &c.Offering)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "totalCoreCount":
			err = unpopulate(val, &c.TotalCoreCount)
			delete(rawMsg, key)
		case "totalNodeCount":
			err = unpopulate(val, &c.TotalNodeCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConnectedClusterUpdateOptions contains the optional parameters for the ConnectedCluster.Update method.
type ConnectedClusterUpdateOptions struct {
	// placeholder for future optional parameters
}

// CredentialResult - The credential result response.
type CredentialResult struct {
	// READ-ONLY; The name of the credential.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Value []byte `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResult.
func (c CredentialResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", c.Name)
	populateByteArray(objectMap, "value", c.Value, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CredentialResult.
func (c *CredentialResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &c.Name)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &c.Value, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// CredentialResults - The list of credential result response.
type CredentialResults struct {
	// READ-ONLY; Contains the REP (rendezvous endpoint) and “Sender” access token.
	HybridConnectionConfig *HybridConnectionConfig `json:"hybridConnectionConfig,omitempty" azure:"ro"`

	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Kubeconfigs []*CredentialResult `json:"kubeconfigs,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResults.
func (c CredentialResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hybridConnectionConfig", c.HybridConnectionConfig)
	populate(objectMap, "kubeconfigs", c.Kubeconfigs)
	return json.Marshal(objectMap)
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData
// error response format.).
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The error object.
	InnerError *ErrorDetail `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// HybridConnectionConfig - Contains the REP (rendezvous endpoint) and “Sender” access token.
type HybridConnectionConfig struct {
	// READ-ONLY; Timestamp when this token will be expired.
	ExpirationTime *int64 `json:"expirationTime,omitempty" azure:"ro"`

	// READ-ONLY; Name of the connection
	HybridConnectionName *string `json:"hybridConnectionName,omitempty" azure:"ro"`

	// READ-ONLY; Name of the relay.
	Relay *string `json:"relay,omitempty" azure:"ro"`

	// READ-ONLY; Sender access token
	Token *string `json:"token,omitempty" azure:"ro"`
}

type ListClusterUserCredentialProperties struct {
	// REQUIRED; The mode of client authentication.
	AuthenticationMethod *AuthenticationMethod `json:"authenticationMethod,omitempty"`

	// REQUIRED; Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy *bool `json:"clientProxy,omitempty"`
}

// Operation - The Connected cluster API operation
type Operation struct {
	// READ-ONLY; The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty" azure:"ro"`

	// READ-ONLY; Operation name: {Microsoft.Kubernetes}/{resource}/{operation}
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string `json:"description,omitempty"`

	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.connectedClusters
	Provider *string `json:"provider,omitempty"`

	// Connected Cluster Resource on which the operation is performed
	Resource *string `json:"resource,omitempty"`
}

// OperationList - The paginated list of connected cluster API operations.
type OperationList struct {
	// The link to fetch the next page of connected cluster API operations.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; The list of connected cluster API operations.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationsGetOptions contains the optional parameters for the Operations.Get method.
type OperationsGetOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource modification (UTC).
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
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

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
