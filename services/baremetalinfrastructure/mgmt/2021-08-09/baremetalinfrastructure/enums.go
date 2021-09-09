package baremetalinfrastructure

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AzureBareMetalHardwareTypeNamesEnum enumerates the values for azure bare metal hardware type names enum.
type AzureBareMetalHardwareTypeNamesEnum string

const (
	// CiscoUCS ...
	CiscoUCS AzureBareMetalHardwareTypeNamesEnum = "Cisco_UCS"
	// HPE ...
	HPE AzureBareMetalHardwareTypeNamesEnum = "HPE"
)

// PossibleAzureBareMetalHardwareTypeNamesEnumValues returns an array of possible values for the AzureBareMetalHardwareTypeNamesEnum const type.
func PossibleAzureBareMetalHardwareTypeNamesEnumValues() []AzureBareMetalHardwareTypeNamesEnum {
	return []AzureBareMetalHardwareTypeNamesEnum{CiscoUCS, HPE}
}

// AzureBareMetalInstancePowerStateEnum enumerates the values for azure bare metal instance power state enum.
type AzureBareMetalInstancePowerStateEnum string

const (
	// Restarting ...
	Restarting AzureBareMetalInstancePowerStateEnum = "restarting"
	// Started ...
	Started AzureBareMetalInstancePowerStateEnum = "started"
	// Starting ...
	Starting AzureBareMetalInstancePowerStateEnum = "starting"
	// Stopped ...
	Stopped AzureBareMetalInstancePowerStateEnum = "stopped"
	// Stopping ...
	Stopping AzureBareMetalInstancePowerStateEnum = "stopping"
	// Unknown ...
	Unknown AzureBareMetalInstancePowerStateEnum = "unknown"
)

// PossibleAzureBareMetalInstancePowerStateEnumValues returns an array of possible values for the AzureBareMetalInstancePowerStateEnum const type.
func PossibleAzureBareMetalInstancePowerStateEnumValues() []AzureBareMetalInstancePowerStateEnum {
	return []AzureBareMetalInstancePowerStateEnum{Restarting, Started, Starting, Stopped, Stopping, Unknown}
}

// AzureBareMetalInstanceSizeNamesEnum enumerates the values for azure bare metal instance size names enum.
type AzureBareMetalInstanceSizeNamesEnum string

const (
	// S112 ...
	S112 AzureBareMetalInstanceSizeNamesEnum = "S112"
	// S144 ...
	S144 AzureBareMetalInstanceSizeNamesEnum = "S144"
	// S144m ...
	S144m AzureBareMetalInstanceSizeNamesEnum = "S144m"
	// S192 ...
	S192 AzureBareMetalInstanceSizeNamesEnum = "S192"
	// S192m ...
	S192m AzureBareMetalInstanceSizeNamesEnum = "S192m"
	// S192xm ...
	S192xm AzureBareMetalInstanceSizeNamesEnum = "S192xm"
	// S224 ...
	S224 AzureBareMetalInstanceSizeNamesEnum = "S224"
	// S224m ...
	S224m AzureBareMetalInstanceSizeNamesEnum = "S224m"
	// S224om ...
	S224om AzureBareMetalInstanceSizeNamesEnum = "S224om"
	// S224oo ...
	S224oo AzureBareMetalInstanceSizeNamesEnum = "S224oo"
	// S224oom ...
	S224oom AzureBareMetalInstanceSizeNamesEnum = "S224oom"
	// S224ooo ...
	S224ooo AzureBareMetalInstanceSizeNamesEnum = "S224ooo"
	// S384 ...
	S384 AzureBareMetalInstanceSizeNamesEnum = "S384"
	// S384m ...
	S384m AzureBareMetalInstanceSizeNamesEnum = "S384m"
	// S384xm ...
	S384xm AzureBareMetalInstanceSizeNamesEnum = "S384xm"
	// S384xxm ...
	S384xxm AzureBareMetalInstanceSizeNamesEnum = "S384xxm"
	// S448 ...
	S448 AzureBareMetalInstanceSizeNamesEnum = "S448"
	// S448m ...
	S448m AzureBareMetalInstanceSizeNamesEnum = "S448m"
	// S448om ...
	S448om AzureBareMetalInstanceSizeNamesEnum = "S448om"
	// S448oo ...
	S448oo AzureBareMetalInstanceSizeNamesEnum = "S448oo"
	// S448oom ...
	S448oom AzureBareMetalInstanceSizeNamesEnum = "S448oom"
	// S448ooo ...
	S448ooo AzureBareMetalInstanceSizeNamesEnum = "S448ooo"
	// S576m ...
	S576m AzureBareMetalInstanceSizeNamesEnum = "S576m"
	// S576xm ...
	S576xm AzureBareMetalInstanceSizeNamesEnum = "S576xm"
	// S672 ...
	S672 AzureBareMetalInstanceSizeNamesEnum = "S672"
	// S672m ...
	S672m AzureBareMetalInstanceSizeNamesEnum = "S672m"
	// S672om ...
	S672om AzureBareMetalInstanceSizeNamesEnum = "S672om"
	// S672oo ...
	S672oo AzureBareMetalInstanceSizeNamesEnum = "S672oo"
	// S672oom ...
	S672oom AzureBareMetalInstanceSizeNamesEnum = "S672oom"
	// S672ooo ...
	S672ooo AzureBareMetalInstanceSizeNamesEnum = "S672ooo"
	// S72 ...
	S72 AzureBareMetalInstanceSizeNamesEnum = "S72"
	// S72m ...
	S72m AzureBareMetalInstanceSizeNamesEnum = "S72m"
	// S768 ...
	S768 AzureBareMetalInstanceSizeNamesEnum = "S768"
	// S768m ...
	S768m AzureBareMetalInstanceSizeNamesEnum = "S768m"
	// S768xm ...
	S768xm AzureBareMetalInstanceSizeNamesEnum = "S768xm"
	// S896 ...
	S896 AzureBareMetalInstanceSizeNamesEnum = "S896"
	// S896m ...
	S896m AzureBareMetalInstanceSizeNamesEnum = "S896m"
	// S896om ...
	S896om AzureBareMetalInstanceSizeNamesEnum = "S896om"
	// S896oo ...
	S896oo AzureBareMetalInstanceSizeNamesEnum = "S896oo"
	// S896oom ...
	S896oom AzureBareMetalInstanceSizeNamesEnum = "S896oom"
	// S896ooo ...
	S896ooo AzureBareMetalInstanceSizeNamesEnum = "S896ooo"
	// S96 ...
	S96 AzureBareMetalInstanceSizeNamesEnum = "S96"
	// S960m ...
	S960m AzureBareMetalInstanceSizeNamesEnum = "S960m"
)

// PossibleAzureBareMetalInstanceSizeNamesEnumValues returns an array of possible values for the AzureBareMetalInstanceSizeNamesEnum const type.
func PossibleAzureBareMetalInstanceSizeNamesEnumValues() []AzureBareMetalInstanceSizeNamesEnum {
	return []AzureBareMetalInstanceSizeNamesEnum{S112, S144, S144m, S192, S192m, S192xm, S224, S224m, S224om, S224oo, S224oom, S224ooo, S384, S384m, S384xm, S384xxm, S448, S448m, S448om, S448oo, S448oom, S448ooo, S576m, S576xm, S672, S672m, S672om, S672oo, S672oom, S672ooo, S72, S72m, S768, S768m, S768xm, S896, S896m, S896om, S896oo, S896oom, S896ooo, S96, S960m}
}

// AzureBareMetalProvisioningStatesEnum enumerates the values for azure bare metal provisioning states enum.
type AzureBareMetalProvisioningStatesEnum string

const (
	// Accepted ...
	Accepted AzureBareMetalProvisioningStatesEnum = "Accepted"
	// Creating ...
	Creating AzureBareMetalProvisioningStatesEnum = "Creating"
	// Deleting ...
	Deleting AzureBareMetalProvisioningStatesEnum = "Deleting"
	// Failed ...
	Failed AzureBareMetalProvisioningStatesEnum = "Failed"
	// Migrating ...
	Migrating AzureBareMetalProvisioningStatesEnum = "Migrating"
	// Succeeded ...
	Succeeded AzureBareMetalProvisioningStatesEnum = "Succeeded"
	// Updating ...
	Updating AzureBareMetalProvisioningStatesEnum = "Updating"
)

// PossibleAzureBareMetalProvisioningStatesEnumValues returns an array of possible values for the AzureBareMetalProvisioningStatesEnum const type.
func PossibleAzureBareMetalProvisioningStatesEnumValues() []AzureBareMetalProvisioningStatesEnum {
	return []AzureBareMetalProvisioningStatesEnum{Accepted, Creating, Deleting, Failed, Migrating, Succeeded, Updating}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}
