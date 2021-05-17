module github.com/Azure/azure-sdk-for-go/tools/profileBuilder

go 1.13

require (
	github.com/Azure/azure-sdk-for-go v54.2.1+incompatible
	github.com/Azure/azure-sdk-for-go/tools/internal v0.0.0
	github.com/spf13/cobra v1.1.3 // indirect
)

replace github.com/Azure/azure-sdk-for-go/tools/internal => ./../internal
