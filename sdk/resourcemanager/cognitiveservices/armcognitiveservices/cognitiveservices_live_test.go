//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcognitiveservices_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/stretchr/testify/suite"
)

type CognitiveservicesTestSuite struct {
	suite.Suite

	ctx                           context.Context
	cred                          azcore.TokenCredential
	options                       *arm.ClientOptions
	accountId                     string
	accountName                   string
	commitmentPlanAssociationName string
	commitmentPlanName            string
	deploymentName                string
	location                      string
	resourceGroupName             string
	subscriptionId                string
}

func (testsuite *CognitiveservicesTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cognitiveservices/armcognitiveservices/testdata")

	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.accountName, _ = recording.GenerateAlphaNumericID(testsuite.T(), "accountn", 14, false)
	testsuite.commitmentPlanAssociationName, _ = recording.GenerateAlphaNumericID(testsuite.T(), "commitmeass", 17, false)
	testsuite.commitmentPlanName, _ = recording.GenerateAlphaNumericID(testsuite.T(), "commitme", 14, false)
	testsuite.deploymentName, _ = recording.GenerateAlphaNumericID(testsuite.T(), "deployme", 14, false)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *CognitiveservicesTestSuite) TearDownSuite() {
	testsuite.Cleanup()
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestCognitiveservicesTestSuite(t *testing.T) {
	suite.Run(t, new(CognitiveservicesTestSuite))
}

func (testsuite *CognitiveservicesTestSuite) Prepare() {
	var err error
	// From step Accounts_Create
	fmt.Println("Call operation: Accounts_Create")
	accountsClient, err := armcognitiveservices.NewAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	accountsClientCreateResponsePoller, err := accountsClient.BeginCreate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcognitiveservices.Account{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr("CognitiveServices"),
		Location: to.Ptr(testsuite.location),
		Properties: &armcognitiveservices.AccountProperties{
			CustomSubDomainName: to.Ptr(testsuite.accountName),
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
		},
	}, nil)
	testsuite.Require().NoError(err)
	var accountsClientCreateResponse *armcognitiveservices.AccountsClientCreateResponse
	accountsClientCreateResponse, err = testutil.PollForTest(testsuite.ctx, accountsClientCreateResponsePoller)
	testsuite.Require().NoError(err)
	testsuite.accountId = *accountsClientCreateResponse.ID
}

// Microsoft.CognitiveServices/accounts/{accountName}
func (testsuite *CognitiveservicesTestSuite) TestAccounts() {
	var err error
	// From step Accounts_List
	fmt.Println("Call operation: Accounts_List")
	accountsClient, err := armcognitiveservices.NewAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	accountsClientNewListPager := accountsClient.NewListPager(nil)
	for accountsClientNewListPager.More() {
		_, err := accountsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Accounts_ListByResourceGroup
	fmt.Println("Call operation: Accounts_ListByResourceGroup")
	accountsClientNewListByResourceGroupPager := accountsClient.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for accountsClientNewListByResourceGroupPager.More() {
		_, err := accountsClientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Accounts_Get
	fmt.Println("Call operation: Accounts_Get")
	_, err = accountsClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)

	// From step Accounts_ListUsages
	fmt.Println("Call operation: Accounts_ListUsages")
	_, err = accountsClient.ListUsages(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, &armcognitiveservices.AccountsClientListUsagesOptions{Filter: nil})
	testsuite.Require().NoError(err)

	// From step Accounts_ListSkus
	fmt.Println("Call operation: Accounts_ListSkus")
	_, err = accountsClient.ListSKUs(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)

	// From step Accounts_ListModels
	fmt.Println("Call operation: Accounts_ListModels")
	accountsClientNewListModelsPager := accountsClient.NewListModelsPager(testsuite.resourceGroupName, testsuite.accountName, nil)
	for accountsClientNewListModelsPager.More() {
		_, err := accountsClientNewListModelsPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Accounts_Update
	fmt.Println("Call operation: Accounts_Update")
	accountsClientUpdateResponsePoller, err := accountsClient.BeginUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcognitiveservices.Account{
		Location: to.Ptr(testsuite.location),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, accountsClientUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step Accounts_RegenerateKey
	fmt.Println("Call operation: Accounts_RegenerateKey")
	_, err = accountsClient.RegenerateKey(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcognitiveservices.RegenerateKeyParameters{
		KeyName: to.Ptr(armcognitiveservices.KeyNameKey2),
	}, nil)
	testsuite.Require().NoError(err)

	// From step Accounts_ListKeys
	fmt.Println("Call operation: Accounts_ListKeys")
	_, err = accountsClient.ListKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.CognitiveServices/accounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}
func (testsuite *CognitiveservicesTestSuite) TestPrivateEndpointConnections() {
	var privateEndpointConnectionName string
	var err error
	// From step Create_PrivateEndpoint
	template := map[string]any{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"parameters": map[string]any{
			"accountId": map[string]any{
				"type":         "string",
				"defaultValue": testsuite.accountId,
			},
			"location": map[string]any{
				"type":         "string",
				"defaultValue": testsuite.location,
			},
			"networkInterfaceName": map[string]any{
				"type":         "string",
				"defaultValue": "epcognitiveservice-nic",
			},
			"privateEndpointName": map[string]any{
				"type":         "string",
				"defaultValue": "epcognitiveservice",
			},
			"virtualNetworksName": map[string]any{
				"type":         "string",
				"defaultValue": "cognitiveservicevnet",
			},
		},
		"resources": []any{
			map[string]any{
				"name":       "[parameters('virtualNetworksName')]",
				"type":       "Microsoft.Network/virtualNetworks",
				"apiVersion": "2020-11-01",
				"location":   "[parameters('location')]",
				"properties": map[string]any{
					"addressSpace": map[string]any{
						"addressPrefixes": []any{
							"10.0.0.0/16",
						},
					},
					"enableDdosProtection": false,
					"subnets": []any{
						map[string]any{
							"name": "default",
							"properties": map[string]any{
								"addressPrefix":                     "10.0.0.0/24",
								"delegations":                       []any{},
								"privateEndpointNetworkPolicies":    "Disabled",
								"privateLinkServiceNetworkPolicies": "Enabled",
							},
						},
					},
					"virtualNetworkPeerings": []any{},
				},
			},
			map[string]any{
				"name":       "[parameters('networkInterfaceName')]",
				"type":       "Microsoft.Network/networkInterfaces",
				"apiVersion": "2020-11-01",
				"dependsOn": []any{
					"[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('virtualNetworksName'), 'default')]",
				},
				"location": "[parameters('location')]",
				"properties": map[string]any{
					"dnsSettings": map[string]any{
						"dnsServers": []any{},
					},
					"enableIPForwarding": false,
					"ipConfigurations": []any{
						map[string]any{
							"name": "privateEndpointIpConfig",
							"properties": map[string]any{
								"primary":                   true,
								"privateIPAddress":          "10.0.0.4",
								"privateIPAddressVersion":   "IPv4",
								"privateIPAllocationMethod": "Dynamic",
								"subnet": map[string]any{
									"id": "[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('virtualNetworksName'), 'default')]",
								},
							},
						},
					},
				},
			},
			map[string]any{
				"name":       "[parameters('privateEndpointName')]",
				"type":       "Microsoft.Network/privateEndpoints",
				"apiVersion": "2020-11-01",
				"dependsOn": []any{
					"[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('virtualNetworksName'), 'default')]",
				},
				"location": "[parameters('location')]",
				"properties": map[string]any{
					"customDnsConfigs":                    []any{},
					"manualPrivateLinkServiceConnections": []any{},
					"privateLinkServiceConnections": []any{
						map[string]any{
							"name": "[parameters('privateEndpointName')]",
							"properties": map[string]any{
								"groupIds": []any{
									"account",
								},
								"privateLinkServiceConnectionState": map[string]any{
									"description":     "Auto-Approved",
									"actionsRequired": "None",
									"status":          "Approved",
								},
								"privateLinkServiceId": "[parameters('accountId')]",
							},
						},
					},
					"subnet": map[string]any{
						"id": "[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('virtualNetworksName'), 'default')]",
					},
				},
			},
			map[string]any{
				"name":       "[concat(parameters('virtualNetworksName'), '/default')]",
				"type":       "Microsoft.Network/virtualNetworks/subnets",
				"apiVersion": "2020-11-01",
				"dependsOn": []any{
					"[resourceId('Microsoft.Network/virtualNetworks', parameters('virtualNetworksName'))]",
				},
				"properties": map[string]any{
					"addressPrefix":                     "10.0.0.0/24",
					"delegations":                       []any{},
					"privateEndpointNetworkPolicies":    "Disabled",
					"privateLinkServiceNetworkPolicies": "Enabled",
				},
			},
		},
		"variables": map[string]any{},
	}
	deployment := armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Template: template,
			Mode:     to.Ptr(armresources.DeploymentModeIncremental),
		},
	}
	_, err = testutil.CreateDeployment(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName, "Create_PrivateEndpoint", &deployment)
	testsuite.Require().NoError(err)

	// From step PrivateEndpointConnections_List
	fmt.Println("Call operation: PrivateEndpointConnections_List")
	privateEndpointConnectionsClient, err := armcognitiveservices.NewPrivateEndpointConnectionsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	privateEndpointConnectionsClientListResponse, err := privateEndpointConnectionsClient.List(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)
	privateEndpointConnectionName = *privateEndpointConnectionsClientListResponse.Value[0].Name
	// accountName/privateEndpointConnectionName
	_, privateEndpointConnectionName, _ = strings.Cut(privateEndpointConnectionName, "/")

	// From step PrivateEndpointConnections_CreateOrUpdate
	fmt.Println("Call operation: PrivateEndpointConnections_CreateOrUpdate")
	privateEndpointConnectionsClientCreateOrUpdateResponsePoller, err := privateEndpointConnectionsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, privateEndpointConnectionName, armcognitiveservices.PrivateEndpointConnection{
		Properties: &armcognitiveservices.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcognitiveservices.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armcognitiveservices.PrivateEndpointServiceConnectionStatusRejected),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, privateEndpointConnectionsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step PrivateEndpointConnections_Get
	fmt.Println("Call operation: PrivateEndpointConnections_Get")
	_, err = privateEndpointConnectionsClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, privateEndpointConnectionName, nil)
	testsuite.Require().NoError(err)

	// From step PrivateLinkResources_List
	fmt.Println("Call operation: PrivateLinkResources_List")
	privateLinkResourcesClient, err := armcognitiveservices.NewPrivateLinkResourcesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = privateLinkResourcesClient.List(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)

	// From step PrivateEndpointConnections_Delete
	fmt.Println("Call operation: PrivateEndpointConnections_Delete")
	privateEndpointConnectionsClientDeleteResponsePoller, err := privateEndpointConnectionsClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, privateEndpointConnectionName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, privateEndpointConnectionsClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.CognitiveServices/operations
func (testsuite *CognitiveservicesTestSuite) TestOperations() {
	var err error
	// From step Operations_List
	fmt.Println("Call operation: Operations_List")
	operationsClient, err := armcognitiveservices.NewOperationsClient(testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	operationsClientNewListPager := operationsClient.NewListPager(nil)
	for operationsClientNewListPager.More() {
		_, err := operationsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}
}

// Microsoft.CognitiveServices/checkDomainAvailability
func (testsuite *CognitiveservicesTestSuite) TestCheckDomainAvailability() {
	var err error
	// From step CheckDomainAvailability
	fmt.Println("Call operation: CheckDomainAvailability")
	managementClient, err := armcognitiveservices.NewManagementClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = managementClient.CheckDomainAvailability(testsuite.ctx, armcognitiveservices.CheckDomainAvailabilityParameter{
		Type:          to.Ptr("Microsoft.CognitiveServices/accounts"),
		SubdomainName: to.Ptr("contosodemoapp1"),
	}, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability
func (testsuite *CognitiveservicesTestSuite) TestCheckSkuAvailability() {
	var err error
	// From step CheckSkuAvailability
	fmt.Println("Call operation: CheckSkuAvailability")
	managementClient, err := armcognitiveservices.NewManagementClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = managementClient.CheckSKUAvailability(testsuite.ctx, testsuite.location, armcognitiveservices.CheckSKUAvailabilityParameter{
		Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		Kind: to.Ptr("Face"),
		SKUs: []*string{
			to.Ptr("S0")},
	}, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.CognitiveServices/locations/{location}/commitmentTiers
func (testsuite *CognitiveservicesTestSuite) TestCommitmentTiers() {
	var err error
	// From step CommitmentTiers_List
	fmt.Println("Call operation: CommitmentTiers_List")
	commitmentTiersClient, err := armcognitiveservices.NewCommitmentTiersClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	commitmentTiersClientNewListPager := commitmentTiersClient.NewListPager(testsuite.location, nil)
	for commitmentTiersClientNewListPager.More() {
		_, err := commitmentTiersClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}
}

// Microsoft.CognitiveServices/skus
func (testsuite *CognitiveservicesTestSuite) TestResourceSkus() {
	var err error
	// From step ResourceSkus_List
	fmt.Println("Call operation: ResourceSkus_List")
	resourceSKUsClient, err := armcognitiveservices.NewResourceSKUsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	resourceSKUsClientNewListPager := resourceSKUsClient.NewListPager(nil)
	for resourceSKUsClientNewListPager.More() {
		_, err := resourceSKUsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}
}

func (testsuite *CognitiveservicesTestSuite) Cleanup() {
	var err error
	// From step Accounts_Delete
	fmt.Println("Call operation: Accounts_Delete")
	accountsClient, err := armcognitiveservices.NewAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	accountsClientDeleteResponsePoller, err := accountsClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, accountsClientDeleteResponsePoller)
	testsuite.Require().NoError(err)

	// From step DeletedAccounts_List
	fmt.Println("Call operation: DeletedAccounts_List")
	deletedAccountsClient, err := armcognitiveservices.NewDeletedAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	deletedAccountsClientNewListPager := deletedAccountsClient.NewListPager(nil)
	for deletedAccountsClientNewListPager.More() {
		_, err := deletedAccountsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step DeletedAccounts_Get
	fmt.Println("Call operation: DeletedAccounts_Get")
	_, err = deletedAccountsClient.Get(testsuite.ctx, testsuite.location, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)

	// From step DeletedAccounts_Purge
	fmt.Println("Call operation: DeletedAccounts_Purge")
	deletedAccountsClientPurgeResponsePoller, err := deletedAccountsClient.BeginPurge(testsuite.ctx, testsuite.location, testsuite.resourceGroupName, testsuite.accountName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, deletedAccountsClientPurgeResponsePoller)
	testsuite.Require().NoError(err)
}
