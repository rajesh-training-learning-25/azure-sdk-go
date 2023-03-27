//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiPolicies.json
func ExampleAPIPolicyClient_ListByAPI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIPolicyClient().ListByAPI(ctx, "rg1", "apimService1", "5600b59475ff190048040001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyCollection = armapimanagement.PolicyCollection{
	// 	Count: to.Ptr[int64](1),
	// 	Value: []*armapimanagement.PolicyContract{
	// 		{
	// 			Name: to.Ptr("policy"),
	// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/policies"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5600b59475ff190048040001/policies/policy"),
	// 			Properties: &armapimanagement.PolicyContractProperties{
	// 				Value: to.Ptr("<!--\r\n    IMPORTANT:\r\n    - Policy elements can appear only within the <inbound>, <outbound>, <backend> section elements.\r\n    - Only the <forward-request> policy element can appear within the <backend> section element.\r\n    - To apply a policy to the incoming request (before it is forwarded to the backend service), place a corresponding policy element within the <inbound> section element.\r\n    - To apply a policy to the outgoing response (before it is sent back to the caller), place a corresponding policy element within the <outbound> section element.\r\n    - To add a policy position the cursor at the desired insertion point and click on the round button associated with the policy.\r\n    - To remove a policy, delete the corresponding policy statement from the policy document.\r\n    - Position the <base> element within a section element to inherit all policies from the corresponding section element in the enclosing scope.\r\n    - Remove the <base> element to prevent inheriting policies from the corresponding section element in the enclosing scope.\r\n    - Policies are applied in the order of their appearance, from the top down.\r\n-->\r\n<policies>\r\n  <inbound>\r\n    <quota-by-key calls=\"5\" bandwidth=\"2\" renewal-period=\"&#x9;P3Y6M4DT12H30M5S\" counter-key=\"ba\" />\r\n    <base />\r\n  </inbound>\r\n  <backend>\r\n    <base />\r\n  </backend>\r\n  <outbound>\r\n    <log-to-eventhub logger-id=\"apimService1\" partition-key=\"@(context.Subscription.Id)\">\r\n@{\r\n	Random Random = new Random();\r\n				const string Chars = \"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz \";                \r\n                return string.Join(\",\", DateTime.UtcNow, new string(\r\n                    Enumerable.Repeat(Chars, Random.Next(2150400))\r\n                              .Select(s =&gt; s[Random.Next(s.Length)])\r\n                              .ToArray()));\r\n          }                           \r\n                        </log-to-eventhub>\r\n    <base />\r\n  </outbound>\r\n</policies>"),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiPolicy.json
func ExampleAPIPolicyClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAPIPolicyClient().GetEntityTag(ctx, "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", armapimanagement.PolicyIDNamePolicy, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiPolicy.json
func ExampleAPIPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIPolicyClient().Get(ctx, "rg1", "apimService1", "5600b59475ff190048040001", armapimanagement.PolicyIDNamePolicy, &armapimanagement.APIPolicyClientGetOptions{Format: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyContract = armapimanagement.PolicyContract{
	// 	Name: to.Ptr("policy"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/policies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5600b59475ff190048040001/policies/policy"),
	// 	Properties: &armapimanagement.PolicyContractProperties{
	// 		Value: to.Ptr("<!--\r\n    IMPORTANT:\r\n    - Policy elements can appear only within the <inbound>, <outbound>, <backend> section elements.\r\n    - Only the <forward-request> policy element can appear within the <backend> section element.\r\n    - To apply a policy to the incoming request (before it is forwarded to the backend service), place a corresponding policy element within the <inbound> section element.\r\n    - To apply a policy to the outgoing response (before it is sent back to the caller), place a corresponding policy element within the <outbound> section element.\r\n    - To add a policy position the cursor at the desired insertion point and click on the round button associated with the policy.\r\n    - To remove a policy, delete the corresponding policy statement from the policy document.\r\n    - Position the <base> element within a section element to inherit all policies from the corresponding section element in the enclosing scope.\r\n    - Remove the <base> element to prevent inheriting policies from the corresponding section element in the enclosing scope.\r\n    - Policies are applied in the order of their appearance, from the top down.\r\n-->\r\n<policies>\r\n  <inbound>\r\n    <quota-by-key calls=\"5\" bandwidth=\"2\" renewal-period=\"&#x9;P3Y6M4DT12H30M5S\" counter-key=\"ba\" />\r\n    <base />\r\n  </inbound>\r\n  <backend>\r\n    <base />\r\n  </backend>\r\n  <outbound>\r\n    <log-to-eventhub logger-id=\"apimService1\" partition-key=\"@(context.Subscription.Id)\">\r\n@{\r\n	Random Random = new Random();\r\n				const string Chars = \"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz \";                \r\n                return string.Join(\",\", DateTime.UtcNow, new string(\r\n                    Enumerable.Repeat(Chars, Random.Next(2150400))\r\n                              .Select(s =&gt; s[Random.Next(s.Length)])\r\n                              .ToArray()));\r\n          }                           \r\n                        </log-to-eventhub>\r\n    <base />\r\n  </outbound>\r\n</policies>"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiPolicy.json
func ExampleAPIPolicyClient_CreateOrUpdate_apiManagementCreateApiPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIPolicyClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5600b57e7e8880006a040001", armapimanagement.PolicyIDNamePolicy, armapimanagement.PolicyContract{
		Properties: &armapimanagement.PolicyContractProperties{
			Format: to.Ptr(armapimanagement.PolicyContentFormatXML),
			Value:  to.Ptr("<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>"),
		},
	}, &armapimanagement.APIPolicyClientCreateOrUpdateOptions{IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyContract = armapimanagement.PolicyContract{
	// 	Name: to.Ptr("policy"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/policies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5600b57e7e8880006a040001/policies/policy"),
	// 	Properties: &armapimanagement.PolicyContractProperties{
	// 		Value: to.Ptr("<policies>\r\n  <inbound />\r\n  <backend>\r\n    <forward-request />\r\n  </backend>\r\n  <outbound />\r\n</policies>"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiPolicyNonXmlEncoded.json
func ExampleAPIPolicyClient_CreateOrUpdate_apiManagementCreateApiPolicyNonXmlEncoded() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIPolicyClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5600b57e7e8880006a040001", armapimanagement.PolicyIDNamePolicy, armapimanagement.PolicyContract{
		Properties: &armapimanagement.PolicyContractProperties{
			Format: to.Ptr(armapimanagement.PolicyContentFormatRawxml),
			Value:  to.Ptr("<policies>\r\n     <inbound>\r\n     <base />\r\n  <set-header name=\"newvalue\" exists-action=\"override\">\r\n   <value>\"@(context.Request.Headers.FirstOrDefault(h => h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n      </policies>"),
		},
	}, &armapimanagement.APIPolicyClientCreateOrUpdateOptions{IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyContract = armapimanagement.PolicyContract{
	// 	Name: to.Ptr("policy"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/operations/policies"),
	// 	ID: to.Ptr("/subscriptions/4c1a3bc6-89f9-46fe-a175-5d8984b25095/resourcegroups/Api-DF-West-US/providers/Microsoft.ApiManagement/service/samirmsiservice2/apis/echo-api/operations/create-resource/policies/policy"),
	// 	Properties: &armapimanagement.PolicyContractProperties{
	// 		Value: to.Ptr("<policies>\r\n  <inbound>\r\n    <base />\r\n    <set-header name=\"newvalue\" exists-action=\"override\">\r\n      <value>\"@(context.Request.Headers.FirstOrDefault(h =&gt; h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n</policies>"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiPolicy.json
func ExampleAPIPolicyClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAPIPolicyClient().Delete(ctx, "rg1", "apimService1", "loggerId", armapimanagement.PolicyIDNamePolicy, "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
