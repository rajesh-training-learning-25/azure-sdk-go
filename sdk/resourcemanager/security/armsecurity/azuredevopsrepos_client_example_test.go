//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListAzureDevOpsRepos_example.json
func ExampleAzureDevOpsReposClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsReposClient().NewListPager("myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AzureDevOpsRepositoryListResponse = armsecurity.AzureDevOpsRepositoryListResponse{
		// 	Value: []*armsecurity.AzureDevOpsRepository{
		// 		{
		// 			Name: to.Ptr("myAzDevOpsRepo"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs/projects/repos"),
		// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg/projects/myAzDevOpsProject/repos/myAzDevOpsRepo"),
		// 			Properties: &armsecurity.AzureDevOpsRepositoryProperties{
		// 				ActionableRemediation: &armsecurity.ActionableRemediation{
		// 					State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
		// 				},
		// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
		// 				ParentOrgName: to.Ptr("myAzDevOpsOrg"),
		// 				ParentProjectName: to.Ptr("myAzDevOpsProject"),
		// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
		// 				RepoID: to.Ptr("cb64ab91-c9ba-46df-b44c-c769358bccdf"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetAzureDevOpsRepos_example.json
func ExampleAzureDevOpsReposClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureDevOpsReposClient().Get(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject", "myAzDevOpsRepo", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsRepository = armsecurity.AzureDevOpsRepository{
	// 	Name: to.Ptr("myAzDevOpsRepo"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs/projects/repos"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg/projects/myAzDevOpsProject/repos/myAzDevOpsRepo"),
	// 	Properties: &armsecurity.AzureDevOpsRepositoryProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ParentOrgName: to.Ptr("myAzDevOpsOrg"),
	// 		ParentProjectName: to.Ptr("myAzDevOpsProject"),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 		RepoID: to.Ptr("cb64ab91-c9ba-46df-b44c-c769358bccdf"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/CreateOrUpdateAzureDevOpsRepos_example.json
func ExampleAzureDevOpsReposClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsReposClient().BeginCreateOrUpdate(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject", "myAzDevOpsRepo", armsecurity.AzureDevOpsRepository{
		Properties: &armsecurity.AzureDevOpsRepositoryProperties{
			ActionableRemediation: &armsecurity.ActionableRemediation{
				State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
			},
			OnboardingState: to.Ptr(armsecurity.OnboardingStateNotApplicable),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsRepository = armsecurity.AzureDevOpsRepository{
	// 	Name: to.Ptr("myAzDevOpsRepo"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs/projects/repos"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg/projects/myAzDevOpsProject/repos/myAzDevOpsRepo"),
	// 	Properties: &armsecurity.AzureDevOpsRepositoryProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ParentOrgName: to.Ptr("myAzDevOpsOrg"),
	// 		ParentProjectName: to.Ptr("myAzDevOpsProject"),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 		RepoID: to.Ptr("cb64ab91-c9ba-46df-b44c-c769358bccdf"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c63d685a0b03817a504b04be938ce46d06ac19/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/UpdateAzureDevOpsRepos_example.json
func ExampleAzureDevOpsReposClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsReposClient().BeginUpdate(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject", "myAzDevOpsRepo", armsecurity.AzureDevOpsRepository{
		Properties: &armsecurity.AzureDevOpsRepositoryProperties{
			ActionableRemediation: &armsecurity.ActionableRemediation{
				State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
			},
			OnboardingState: to.Ptr(armsecurity.OnboardingStateNotApplicable),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsRepository = armsecurity.AzureDevOpsRepository{
	// 	Name: to.Ptr("myAzDevOpsRepo"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs/projects/repos"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg/projects/myAzDevOpsProject/repos/myAzDevOpsRepo"),
	// 	Properties: &armsecurity.AzureDevOpsRepositoryProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ParentOrgName: to.Ptr("myAzDevOpsOrg"),
	// 		ParentProjectName: to.Ptr("myAzDevOpsProject"),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 		RepoID: to.Ptr("cb64ab91-c9ba-46df-b44c-c769358bccdf"),
	// 	},
	// }
}
