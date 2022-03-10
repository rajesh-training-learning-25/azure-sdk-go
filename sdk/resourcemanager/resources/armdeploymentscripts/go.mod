module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts

go 1.16

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.22.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.13.1
	github.com/Azure/azure-sdk-for-go/sdk/internal v0.9.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal v0.0.0-00010101000000-000000000000
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi v0.3.1
	github.com/stretchr/testify v1.7.0
)

replace github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal => ../../internal
