//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-09-01/examples/ListOperations.json
func ExampleCertificateRegistrationProviderClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewCertificateRegistrationProviderClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListOperationsPager(nil)
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
		// page.CsmOperationCollection = armappservice.CsmOperationCollection{
		// 	Value: []*armappservice.CsmOperationDescription{
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates/Write"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Add a new certificate or update an existing one"),
		// 				Operation: to.Ptr("Add or Update Certificate"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("Certificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/Write"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Add a new certificateOrder or update an existing one"),
		// 				Operation: to.Ptr("Add or Update AppServiceCertificate"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates/Delete"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Delete an existing certificate"),
		// 				Operation: to.Ptr("Delete Certificate"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("Certificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/Delete"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Delete an existing AppServiceCertificate"),
		// 				Operation: to.Ptr("Delete AppServiceCertificate"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/Read"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Get the list of CertificateOrders"),
		// 				Operation: to.Ptr("Get CertificateOrders"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates/Read"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Get the list of certificates"),
		// 				Operation: to.Ptr("Get Certificates"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("Certificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/reissue/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Reissue an existing certificateorder"),
		// 				Operation: to.Ptr("Reissue certificateorder"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/renew/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Renew an existing certificateorder"),
		// 				Operation: to.Ptr("Renew certificateorder"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/retrieveCertificateActions/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Retrieve the list of certificate actions"),
		// 				Operation: to.Ptr("Certificateorder actions"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/retrieveEmailHistory/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Retrieve certificate email history"),
		// 				Operation: to.Ptr("Certificateorder email history"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/resendEmail/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Resend certificate email"),
		// 				Operation: to.Ptr("Resend Certificateorder email"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/verifyDomainOwnership/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Verify domain ownership"),
		// 				Operation: to.Ptr("Verify domain ownership"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/resendRequestEmails/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Resend request emails to another email address"),
		// 				Operation: to.Ptr("Resend request emails to another email address"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/resendRequestEmails/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Retrieve site seal for an issued App Service Certificate"),
		// 				Operation: to.Ptr("Retrieve site seal for an issued App Service Certificate"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/provisionGlobalAppServicePrincipalInUserTenant/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Provision service principal for service app principal"),
		// 				Operation: to.Ptr("Provision service principal"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/validateCertificateRegistrationInformation/Action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Validate certificate purchase object without submitting it"),
		// 				Operation: to.Ptr("Certificate Purchase Info Validation"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("AppServiceCertificate"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CertificateRegistration/register/action"),
		// 			Display: &armappservice.CsmOperationDisplay{
		// 				Description: to.Ptr("Register the Microsoft Certificates resource provider for the subscription"),
		// 				Operation: to.Ptr("Register Microsoft Certificates resource provider"),
		// 				Provider: to.Ptr("Microsoft Certificates"),
		// 				Resource: to.Ptr("Microsoft Certificates resource provider"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
