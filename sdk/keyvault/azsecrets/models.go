//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azsecrets

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets/internal"
)

// DeletedSecretBundle - A Deleted Secret consisting of its previous id, attributes and its tags, as well as information on when it will be purged.
type DeletedSecretBundle struct {
	SecretBundle
	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// SecretBundle - A secret consisting of a value, id and its attributes.
type SecretBundle struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// The content type of the secret.
	ContentType *string `json:"contentType,omitempty"`

	// The secret id.
	ID *string `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// The secret value.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV certificate.
	Kid *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

func secretBundleFromGenerated(i internal.SecretBundle) SecretBundle {
	return SecretBundle{
		Attributes:  secretAttributesFromGenerated(i.Attributes),
		ContentType: i.ContentType,
		ID:          i.ID,
		Tags:        i.Tags,
		Value:       i.Value,
		Kid:         i.Kid,
		Managed:     i.Managed,
	}
}

// SecretAttributes - The secret management attributes.
type SecretAttributes struct {
	Attributes
	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32 `json:"recoverableDays,omitempty" azure:"ro"`

	// READ-ONLY; Reflects the deletion recovery level currently in effect for secrets in the current vault. If it contains 'Purgeable', the secret can be permanently
	// deleted by a privileged user; otherwise, only the
	// system can purge the secret, at the end of the retention interval.
	RecoveryLevel *DeletionRecoveryLevel `json:"recoveryLevel,omitempty" azure:"ro"`
}

func (s SecretAttributes) toGenerated() *internal.SecretAttributes {
	return &internal.SecretAttributes{
		RecoverableDays: s.RecoverableDays,
		RecoveryLevel:   s.RecoveryLevel.toGenerated().ToPtr(),
		Attributes: internal.Attributes{
			Enabled:   s.Enabled,
			Expires:   s.Expires,
			NotBefore: s.NotBefore,
			Created:   s.Created,
			Updated:   s.Updated,
		},
	}
}

// create a SecretAttributes object from an internal.SecretAttributes object
func secretAttributesFromGenerated(i *internal.SecretAttributes) *SecretAttributes {
	if i == nil {
		return nil
	}
	return &SecretAttributes{
		RecoverableDays: i.RecoverableDays,
		RecoveryLevel:   deletionRecoveryLevelFromGenerated(*i.RecoveryLevel).ToPtr(),
		Attributes: Attributes{
			Enabled:   i.Enabled,
			Expires:   i.Expires,
			NotBefore: i.NotBefore,
			Created:   i.Created,
			Updated:   i.Updated,
		},
	}
}

// Attributes - The object attributes managed by the KeyVault service.
type Attributes struct {
	// Determines whether the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Expiry date in UTC.
	Expires *time.Time `json:"exp,omitempty"`

	// Not before date in UTC.
	NotBefore *time.Time `json:"nbf,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

// SecretItem - The secret item containing secret metadata.
type SecretItem struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// Secret identifier.
	ID *string `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a key backing a certificate, then managed will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

// create a SecretItem from the internal.SecretItem model
func secretItemFromGenerated(i *internal.SecretItem) *SecretItem {
	if i == nil {
		return nil
	}

	return &SecretItem{
		Attributes:  secretAttributesFromGenerated(i.Attributes),
		ContentType: i.ContentType,
		ID:          i.ID,
		Tags:        i.Tags,
		Managed:     i.Managed,
	}
}

// DeletedSecretItem - The deleted secret item containing metadata about the deleted secret.
type DeletedSecretItem struct {
	SecretItem
	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

func deletedSecretItemFromGenerated(i *internal.DeletedSecretItem) *DeletedSecretItem {
	if i == nil {
		return nil
	}

	return &DeletedSecretItem{
		RecoveryID:         i.RecoveryID,
		DeletedDate:        i.DeletedDate,
		ScheduledPurgeDate: i.ScheduledPurgeDate,
		SecretItem:         *secretItemFromGenerated(&i.SecretItem),
	}
}

// BackupSecretResult - The backup secret result, containing the backup blob.
type BackupSecretResult struct {
	// READ-ONLY; The backup blob containing the backed up secret.
	Value []byte `json:"value,omitempty" azure:"ro"`
}
