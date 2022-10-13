//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecuritydevops

import "time"

type ActionableRemediation struct {
	// Branch onboarding info.
	BranchConfiguration *TargetBranchConfiguration  `json:"branchConfiguration,omitempty"`
	Categories          []*RuleCategory             `json:"categories,omitempty"`
	SeverityLevels      []*string                   `json:"severityLevels,omitempty"`
	State               *ActionableRemediationState `json:"state,omitempty"`
}

type AuthorizationInfo struct {
	// Gets or sets one-time OAuth code to exchange for refresh and access tokens.
	Code *string `json:"code,omitempty"`
}

type AzureDevOpsConnector struct {
	// REQUIRED; The geo-location where the resource lives
	Location   *string                         `json:"location,omitempty"`
	Properties *AzureDevOpsConnectorProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureDevOpsConnectorClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginCreateOrUpdate
// method.
type AzureDevOpsConnectorClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsConnectorClientBeginDeleteOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginDelete
// method.
type AzureDevOpsConnectorClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsConnectorClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginUpdate
// method.
type AzureDevOpsConnectorClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsConnectorClientGetOptions contains the optional parameters for the AzureDevOpsConnectorClient.Get method.
type AzureDevOpsConnectorClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsConnectorClientListByResourceGroupOptions contains the optional parameters for the AzureDevOpsConnectorClient.ListByResourceGroup
// method.
type AzureDevOpsConnectorClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsConnectorClientListBySubscriptionOptions contains the optional parameters for the AzureDevOpsConnectorClient.ListBySubscription
// method.
type AzureDevOpsConnectorClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

type AzureDevOpsConnectorListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*AzureDevOpsConnector `json:"value,omitempty"`
}

type AzureDevOpsConnectorProperties struct {
	Authorization *AuthorizationInfo `json:"authorization,omitempty"`

	// Gets or sets org onboarding information.
	Orgs              []*AzureDevOpsOrgMetadata `json:"orgs,omitempty"`
	ProvisioningState *ProvisioningState        `json:"provisioningState,omitempty"`
}

type AzureDevOpsConnectorStats struct {
	Properties *AzureDevOpsConnectorStatsProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureDevOpsConnectorStatsClientGetOptions contains the optional parameters for the AzureDevOpsConnectorStatsClient.Get
// method.
type AzureDevOpsConnectorStatsClientGetOptions struct {
	// placeholder for future optional parameters
}

type AzureDevOpsConnectorStatsListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*AzureDevOpsConnectorStats `json:"value,omitempty"`
}

type AzureDevOpsConnectorStatsProperties struct {
	// Gets or sets orgs count.
	OrgsCount *int64 `json:"orgsCount,omitempty"`

	// Gets or sets projects count.
	ProjectsCount     *int64             `json:"projectsCount,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// Gets or sets repos count.
	ReposCount *int64 `json:"reposCount,omitempty"`
}

// AzureDevOpsOrg - Azure DevOps Org Proxy Resource.
type AzureDevOpsOrg struct {
	// AzureDevOps Org properties.
	Properties *AzureDevOpsOrgProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureDevOpsOrgClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsOrgClient.BeginCreateOrUpdate
// method.
type AzureDevOpsOrgClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsOrgClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsOrgClient.BeginUpdate method.
type AzureDevOpsOrgClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsOrgClientGetOptions contains the optional parameters for the AzureDevOpsOrgClient.Get method.
type AzureDevOpsOrgClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsOrgClientListOptions contains the optional parameters for the AzureDevOpsOrgClient.List method.
type AzureDevOpsOrgClientListOptions struct {
	// placeholder for future optional parameters
}

type AzureDevOpsOrgListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*AzureDevOpsOrg `json:"value,omitempty"`
}

// AzureDevOpsOrgMetadata - Org onboarding info.
type AzureDevOpsOrgMetadata struct {
	AutoDiscovery *AutoDiscovery `json:"autoDiscovery,omitempty"`

	// Gets or sets name of the AzureDevOps Org.
	Name     *string                       `json:"name,omitempty"`
	Projects []*AzureDevOpsProjectMetadata `json:"projects,omitempty"`
}

// AzureDevOpsOrgProperties - AzureDevOps Org properties.
type AzureDevOpsOrgProperties struct {
	AutoDiscovery     *AutoDiscovery     `json:"autoDiscovery,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

// AzureDevOpsProject - Azure DevOps Project Proxy Resource.
type AzureDevOpsProject struct {
	// AzureDevOps Project properties.
	Properties *AzureDevOpsProjectProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureDevOpsProjectClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsProjectClient.BeginCreateOrUpdate
// method.
type AzureDevOpsProjectClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsProjectClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsProjectClient.BeginUpdate
// method.
type AzureDevOpsProjectClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsProjectClientGetOptions contains the optional parameters for the AzureDevOpsProjectClient.Get method.
type AzureDevOpsProjectClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsProjectClientListOptions contains the optional parameters for the AzureDevOpsProjectClient.List method.
type AzureDevOpsProjectClientListOptions struct {
	// placeholder for future optional parameters
}

type AzureDevOpsProjectListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*AzureDevOpsProject `json:"value,omitempty"`
}

// AzureDevOpsProjectMetadata - Project onboarding info.
type AzureDevOpsProjectMetadata struct {
	AutoDiscovery *AutoDiscovery `json:"autoDiscovery,omitempty"`

	// Gets or sets name of the AzureDevOps Project.
	Name *string `json:"name,omitempty"`

	// Gets or sets repositories.
	Repos []*string `json:"repos,omitempty"`
}

// AzureDevOpsProjectProperties - AzureDevOps Project properties.
type AzureDevOpsProjectProperties struct {
	AutoDiscovery *AutoDiscovery `json:"autoDiscovery,omitempty"`

	// Gets or sets AzureDevOps Org Name.
	OrgName *string `json:"orgName,omitempty"`

	// Gets or sets AzureDevOps Project Id.
	ProjectID         *string            `json:"projectId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

// AzureDevOpsRepo - Azure DevOps Repo Proxy Resource.
type AzureDevOpsRepo struct {
	// AzureDevOps Repo properties.
	Properties *AzureDevOpsRepoProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AzureDevOpsRepoClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsRepoClient.BeginCreateOrUpdate
// method.
type AzureDevOpsRepoClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsRepoClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsRepoClient.BeginUpdate method.
type AzureDevOpsRepoClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AzureDevOpsRepoClientGetOptions contains the optional parameters for the AzureDevOpsRepoClient.Get method.
type AzureDevOpsRepoClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsRepoClientListByConnectorOptions contains the optional parameters for the AzureDevOpsRepoClient.ListByConnector
// method.
type AzureDevOpsRepoClientListByConnectorOptions struct {
	// placeholder for future optional parameters
}

// AzureDevOpsRepoClientListOptions contains the optional parameters for the AzureDevOpsRepoClient.List method.
type AzureDevOpsRepoClientListOptions struct {
	// placeholder for future optional parameters
}

type AzureDevOpsRepoListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*AzureDevOpsRepo `json:"value,omitempty"`
}

// AzureDevOpsRepoProperties - AzureDevOps Repo properties.
type AzureDevOpsRepoProperties struct {
	ActionableRemediation *ActionableRemediation `json:"actionableRemediation,omitempty"`

	// Gets or sets AzureDevOps Org Name.
	OrgName *string `json:"orgName,omitempty"`

	// Gets or sets AzureDevOps Project Name.
	ProjectName       *string            `json:"projectName,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// Gets or sets Azure DevOps repo id.
	RepoID *string `json:"repoId,omitempty"`

	// Gets or sets AzureDevOps repo url.
	RepoURL *string `json:"repoUrl,omitempty"`

	// Gets or sets AzureDevOps repo visibility, whether it is public or private etc.
	Visibility *string `json:"visibility,omitempty"`
}

// GitHubConnector - Represents an ARM resource for /subscriptions/xxx/resourceGroups/xxx/providers/Microsoft.SecurityDevOps/gitHubConnectors.
type GitHubConnector struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the ARM resource for /subscriptions/xxx/resourceGroups/xxx/providers/Microsoft.SecurityDevOps/gitHubConnectors.
	Properties *GitHubConnectorProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GitHubConnectorClientBeginCreateOrUpdateOptions contains the optional parameters for the GitHubConnectorClient.BeginCreateOrUpdate
// method.
type GitHubConnectorClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubConnectorClientBeginDeleteOptions contains the optional parameters for the GitHubConnectorClient.BeginDelete method.
type GitHubConnectorClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubConnectorClientBeginUpdateOptions contains the optional parameters for the GitHubConnectorClient.BeginUpdate method.
type GitHubConnectorClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubConnectorClientGetOptions contains the optional parameters for the GitHubConnectorClient.Get method.
type GitHubConnectorClientGetOptions struct {
	// placeholder for future optional parameters
}

// GitHubConnectorClientListByResourceGroupOptions contains the optional parameters for the GitHubConnectorClient.ListByResourceGroup
// method.
type GitHubConnectorClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// GitHubConnectorClientListBySubscriptionOptions contains the optional parameters for the GitHubConnectorClient.ListBySubscription
// method.
type GitHubConnectorClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

type GitHubConnectorListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*GitHubConnector `json:"value,omitempty"`
}

// GitHubConnectorProperties - Properties of the ARM resource for /subscriptions/xxx/resourceGroups/xxx/providers/Microsoft.SecurityDevOps/gitHubConnectors.
type GitHubConnectorProperties struct {
	// Gets or sets one-time OAuth code to exchange for refresh and access tokens.
	Code              *string            `json:"code,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

type GitHubConnectorStats struct {
	Properties *GitHubConnectorStatsProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GitHubConnectorStatsClientGetOptions contains the optional parameters for the GitHubConnectorStatsClient.Get method.
type GitHubConnectorStatsClientGetOptions struct {
	// placeholder for future optional parameters
}

type GitHubConnectorStatsListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*GitHubConnectorStats `json:"value,omitempty"`
}

type GitHubConnectorStatsProperties struct {
	// Gets or sets owners count.
	OwnersCount       *int64             `json:"ownersCount,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// Gets or sets repos count.
	ReposCount *int64 `json:"reposCount,omitempty"`
}

// GitHubOwner - GitHub repo owner Proxy Resource.
type GitHubOwner struct {
	// GitHub Repo Owner properties.
	Properties *GitHubOwnerProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GitHubOwnerClientBeginCreateOrUpdateOptions contains the optional parameters for the GitHubOwnerClient.BeginCreateOrUpdate
// method.
type GitHubOwnerClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubOwnerClientBeginUpdateOptions contains the optional parameters for the GitHubOwnerClient.BeginUpdate method.
type GitHubOwnerClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubOwnerClientGetOptions contains the optional parameters for the GitHubOwnerClient.Get method.
type GitHubOwnerClientGetOptions struct {
	// placeholder for future optional parameters
}

// GitHubOwnerClientListOptions contains the optional parameters for the GitHubOwnerClient.List method.
type GitHubOwnerClientListOptions struct {
	// placeholder for future optional parameters
}

type GitHubOwnerListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*GitHubOwner `json:"value,omitempty"`
}

// GitHubOwnerProperties - GitHub Repo Owner properties.
type GitHubOwnerProperties struct {
	// Gets or sets gitHub owner url.
	OwnerURL          *string            `json:"ownerUrl,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

// GitHubRepo - GitHub repo Proxy Resource.
type GitHubRepo struct {
	// GitHub Repo properties.
	Properties *GitHubRepoProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GitHubRepoClientBeginCreateOrUpdateOptions contains the optional parameters for the GitHubRepoClient.BeginCreateOrUpdate
// method.
type GitHubRepoClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubRepoClientBeginUpdateOptions contains the optional parameters for the GitHubRepoClient.BeginUpdate method.
type GitHubRepoClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GitHubRepoClientGetOptions contains the optional parameters for the GitHubRepoClient.Get method.
type GitHubRepoClientGetOptions struct {
	// placeholder for future optional parameters
}

// GitHubRepoClientListByConnectorOptions contains the optional parameters for the GitHubRepoClient.ListByConnector method.
type GitHubRepoClientListByConnectorOptions struct {
	// placeholder for future optional parameters
}

// GitHubRepoClientListOptions contains the optional parameters for the GitHubRepoClient.List method.
type GitHubRepoClientListOptions struct {
	// placeholder for future optional parameters
}

type GitHubRepoListResponse struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets list of resources.
	Value []*GitHubRepo `json:"value,omitempty"`
}

// GitHubRepoProperties - GitHub Repo properties.
type GitHubRepoProperties struct {
	// Gets or sets gitHub repo account id.
	AccountID *int64 `json:"accountId,omitempty"`

	// Gets or sets GitHub Owner Name.
	OwnerName         *string            `json:"ownerName,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// Gets or sets gitHub repo url.
	RepoURL *string `json:"repoUrl,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TargetBranchConfiguration - Branch onboarding info.
type TargetBranchConfiguration struct {
	// Gets or sets branches that should have annotations.
	// For Ignite, we will be supporting a single default branch configuration in the UX.
	Names []*string `json:"names,omitempty"`
}
