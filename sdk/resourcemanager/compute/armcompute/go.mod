module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v59.0.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.20.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.11.0
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.8.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v0.2.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.11.0 => ../../../azidentity
