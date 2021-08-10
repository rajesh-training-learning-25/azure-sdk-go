// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const SnapshotTimeFormat = "2006-01-02T15:04:05.0000000Z07:00"

// TableSASSignatureValues is used to generate a Shared Access Signature (SAS) for an Azure Storage container or blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/constructing-a-service-sas
type TableSASSignatureValues struct {
	Version      string      `param:"sv"`  // If not specified, this defaults to SASVersion
	Protocol     SASProtocol `param:"spr"` // See the SASProtocol* constants
	StartTime    time.Time   `param:"st"`  // Not specified if IsZero
	ExpiryTime   time.Time   `param:"se"`  // Not specified if IsZero
	SnapshotTime time.Time
	Permissions  string  `param:"sp"` // Create by initializing a ContainerSASPermissions or TableSASPermissions and then call String()
	IPRange      IPRange `param:"sip"`
	Identifier   string  `param:"si"`
	TableName    string
}

// NewSASQueryParameters uses an account's StorageAccountCredential to sign this signature values to produce
// the proper SAS query parameters.
// See: StorageAccountCredential. Compatible with both UserDelegationCredential and SharedKeyCredential
func (v TableSASSignatureValues) NewSASQueryParameters(credential *SharedKeyCredential) (SASQueryParameters, error) {
	resource := ""
	// if credential == nil {
	// 	return SASQueryParameters{}, fmt.Errorf("cannot sign SAS query without StorageAccountCredential")
	// }

	if v.Version != "" {
		// resource = "bv"
		//Make sure the permission characters are in the correct order
		perms := &TableSASPermissions{}
		if err := perms.Parse(v.Permissions); err != nil {
			return SASQueryParameters{}, err
		}
		v.Permissions = perms.String()
	} else if v.TableName == "" {
		// Make sure the permission characters are in the correct order
		perms := &TableSASPermissions{}
		if err := perms.Parse(v.Permissions); err != nil {
			return SASQueryParameters{}, err
		}
		v.Permissions = perms.String()
	} else {
		// resource = "t"
		// Make sure the permission characters are in the correct order
		perms := &TableSASPermissions{}
		if err := perms.Parse(v.Permissions); err != nil {
			return SASQueryParameters{}, err
		}
		v.Permissions = perms.String()
	}
	if v.Version == "" {
		v.Version = SASVersion
	}
	startTime, expiryTime, snapshotTime := FormatTimesForSASSigning(v.StartTime, v.ExpiryTime, v.SnapshotTime)
	_ = snapshotTime

	signedIdentifier := v.Identifier

	p := SASQueryParameters{
		// Common SAS parameters
		version:     v.Version,
		protocol:    v.Protocol,
		startTime:   v.StartTime,
		expiryTime:  v.ExpiryTime,
		permissions: v.Permissions,
		ipRange:     v.IPRange,
		tableName:   v.TableName,

		// Container/Blob-specific SAS parameters
		resource:   resource,
		identifier: v.Identifier,
	}

	startPartitionKey := ""
	startRowKey := ""
	endPartitionKey := ""
	endRowKey := ""

	canonicalName := "/" + "table" + "/" + credential.AccountName() + "/" + v.TableName
	fmt.Println("CANONICAL NAME ", canonicalName)

	// String to sign: http://msdn.microsoft.com/en-us/library/azure/dn140255.aspx
	stringToSign := strings.Join([]string{
		v.Permissions,
		startTime,
		expiryTime,
		// getCanonicalName(credential.AccountName(), v.TableName),
		canonicalName,
		signedIdentifier,
		v.IPRange.String(),
		string(v.Protocol),
		v.Version,
		startPartitionKey,
		startRowKey,
		endPartitionKey,
		endRowKey,
	},
		// resource,
		// snapshotTime}, // signed timestamp
		"\n",
	)

	// TODO: Add start/end pk/rk
	fmt.Println("String to sign\n", stringToSign, "END")

	signature, err := credential.ComputeHMACSHA256(stringToSign)
	p.signature = signature
	return p, err
}

// The ContainerSASPermissions type simplifies creating the permissions string for an Azure Storage container SAS.
// Initialize an instance of this type and then call its String method to set TableSASSignatureValues's Permissions field.
type ContainerSASPermissions struct {
	Read, Add, Create, Write, Delete, List bool
}

// String produces the SAS permissions string for an Azure Storage container.
// Call this method to set TableSASSignatureValues's Permissions field.
func (p ContainerSASPermissions) String() string {
	var b bytes.Buffer
	if p.Read {
		b.WriteRune('r')
	}
	if p.Add {
		b.WriteRune('a')
	}
	if p.Create {
		b.WriteRune('c')
	}
	if p.Write {
		b.WriteRune('w')
	}
	if p.Delete {
		b.WriteRune('d')
	}
	if p.List {
		b.WriteRune('l')
	}
	return b.String()
}

// Parse initializes the ContainerSASPermissions's fields from a string.
func (p *ContainerSASPermissions) Parse(s string) error {
	*p = ContainerSASPermissions{} // Clear the flags
	for _, r := range s {
		switch r {
		case 'r':
			p.Read = true
		case 'a':
			p.Add = true
		case 'c':
			p.Create = true
		case 'w':
			p.Write = true
		case 'd':
			p.Delete = true
		case 'l':
			p.List = true
		default:
			return fmt.Errorf("Invalid permission: '%v'", r)
		}
	}
	return nil
}

// The TableSASPermissions type simplifies creating the permissions string for an Azure Storage blob SAS.
// Initialize an instance of this type and then call its String method to set TableSASSignatureValues's Permissions field.
type TableSASPermissions struct{ Read, Add, Create, Write, Delete, DeletePreviousVersion bool }

// String produces the SAS permissions string for an Azure Storage blob.
// Call this method to set TableSASSignatureValues's Permissions field.
func (p TableSASPermissions) String() string {
	var b bytes.Buffer
	if p.Read {
		b.WriteRune('r')
	}
	if p.Add {
		b.WriteRune('a')
	}
	if p.Create {
		b.WriteRune('c')
	}
	if p.Write {
		b.WriteRune('w')
	}
	if p.Delete {
		b.WriteRune('d')
	}
	if p.DeletePreviousVersion {
		b.WriteRune('x')
	}
	return b.String()
}

// Parse initializes the TableSASPermissions's fields from a string.
func (p *TableSASPermissions) Parse(s string) error {
	*p = TableSASPermissions{} // Clear the flags
	for _, r := range s {
		switch r {
		case 'r':
			p.Read = true
		case 'a':
			p.Add = true
		case 'c':
			p.Create = true
		case 'w':
			p.Write = true
		case 'd':
			p.Delete = true
		case 'x':
			p.DeletePreviousVersion = true
		default:
			return fmt.Errorf("invalid permission: '%v'", r)
		}
	}
	return nil
}
