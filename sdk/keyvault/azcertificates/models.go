//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcertificates

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azcertificates/internal/generated"
)

// Action - The action that will be executed.
type Action struct {
	// The type of the action.
	ActionType *CertificatePolicyAction `json:"action_type,omitempty"`
}

func (a *Action) toGenerated() *generated.Action {
	if a == nil {
		return nil
	}

	return &generated.Action{
		ActionType: (*generated.ActionType)(a.ActionType),
	}
}

// AdministratorDetails - Details of the organization administrator of the certificate issuer.
type AdministratorDetails struct {
	// Email address.
	EmailAddress *string `json:"email,omitempty"`

	// First name.
	FirstName *string `json:"first_name,omitempty"`

	// Last name.
	LastName *string `json:"last_name,omitempty"`

	// Phone number.
	Phone *string `json:"phone,omitempty"`
}

func administratorDetailsFromGenerated(g *generated.AdministratorDetails) *AdministratorDetails {
	if g == nil {
		return nil
	}

	return &AdministratorDetails{
		EmailAddress: g.EmailAddress,
		FirstName:    g.FirstName,
		LastName:     g.LastName,
		Phone:        g.Phone,
	}
}

func (a *AdministratorDetails) toGenerated() *generated.AdministratorDetails {
	if a == nil {
		return nil
	}

	return &generated.AdministratorDetails{
		EmailAddress: a.EmailAddress,
		FirstName:    a.FirstName,
		LastName:     a.LastName,
		Phone:        a.Phone,
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

// CertificateAttributes - The certificate management attributes.
type CertificateAttributes struct {
	// Determines whether the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Expiry date in UTC.
	Expires *time.Time `json:"exp,omitempty"`

	// Not before date in UTC.
	NotBefore *time.Time `json:"nbf,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32 `json:"recoverableDays,omitempty" azure:"ro"`

	// READ-ONLY; Reflects the deletion recovery level currently in effect for certificates in the current vault. If it contains
	// 'Purgeable', the certificate can be permanently deleted by a privileged user; otherwise,
	// only the system can purge the certificate, at the end of the retention interval.
	RecoveryLevel *string `json:"recoveryLevel,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

func (c *CertificateAttributes) toGenerated() *generated.CertificateAttributes {
	if c == nil {
		return nil
	}

	return &generated.CertificateAttributes{
		Created:         c.Created,
		Enabled:         c.Enabled,
		Expires:         c.Expires,
		NotBefore:       c.NotBefore,
		RecoverableDays: c.RecoverableDays,
		RecoveryLevel:   (*generated.DeletionRecoveryLevel)(c.RecoveryLevel),
		Updated:         c.Updated,
	}
}

func certificateAttributesFromGenerated(g *generated.CertificateAttributes) *CertificateAttributes {
	if g == nil {
		return nil
	}

	return &CertificateAttributes{
		Enabled:         g.Enabled,
		Expires:         g.Expires,
		NotBefore:       g.NotBefore,
		Created:         g.Created,
		Updated:         g.Updated,
		RecoverableDays: g.RecoverableDays,
		RecoveryLevel:   (*string)(g.RecoveryLevel),
	}
}

// CertificateBundle - A certificate bundle consists of a certificate (X509) plus its attributes.
type CertificateBundle struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// CER contents of x509 certificate.
	Cer []byte `json:"cer,omitempty"`

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string `json:"contentType,omitempty"`

	// Application specific metadata in the form of key-value pairs
	Tags map[string]string `json:"tags,omitempty"`

	// READ-ONLY; The certificate id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The key id.
	Kid *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy `json:"policy,omitempty" azure:"ro"`

	// READ-ONLY; The secret id.
	Sid *string `json:"sid,omitempty" azure:"ro"`

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty" azure:"ro"`
}

func certificateBundleFromGenerated(g *generated.CertificateBundle) CertificateBundle {
	if g == nil {
		return CertificateBundle{}
	}

	return CertificateBundle{
		Attributes:     certificateAttributesFromGenerated(g.Attributes),
		Cer:            g.Cer,
		ContentType:    g.ContentType,
		Tags:           convertGeneratedMap(g.Tags),
		ID:             g.ID,
		Kid:            g.Kid,
		Policy:         certificatePolicyFromGenerated(g.Policy),
		Sid:            g.Sid,
		X509Thumbprint: g.X509Thumbprint,
	}
}

// CertificateError - The key vault server error.
type CertificateError struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The key vault server error.
	InnerError *CertificateError `json:"innererror,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

func certificateErrorFromGenerated(g *generated.Error) *CertificateError {
	if g == nil {
		return nil
	}

	return &CertificateError{
		Code:       g.Code,
		Message:    g.Message,
		InnerError: certificateErrorFromGenerated(g.InnerError),
	}
}

// CertificateIssuerItem - The certificate issuer item containing certificate issuer metadata.
type CertificateIssuerItem struct {
	// Certificate Identifier.
	ID *string `json:"id,omitempty"`

	// The issuer provider.
	Provider *string `json:"provider,omitempty"`
}

func certificateIssuerItemFromGenerated(g *generated.CertificateIssuerItem) *CertificateIssuerItem {
	if g == nil {
		return nil
	}

	return &CertificateIssuerItem{ID: g.ID, Provider: g.Provider}
}

// CertificateItem - The certificate item containing certificate metadata.
type CertificateItem struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Certificate identifier.
	ID *string `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	// Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty"`
}

// CertificateOperation - A certificate operation is returned in case of asynchronous requests.
type CertificateOperation struct {
	// Indicates if cancellation was requested on the certificate operation.
	CancellationRequested *bool `json:"cancellation_requested,omitempty"`

	// The certificate signing request (CSR) that is being used in the certificate operation.
	Csr []byte `json:"csr,omitempty"`

	// Error encountered, if any, during the certificate operation.
	Error *CertificateError `json:"error,omitempty"`

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters `json:"issuer,omitempty"`

	// Identifier for the certificate operation.
	RequestID *string `json:"request_id,omitempty"`

	// Status of the certificate operation.
	Status *string `json:"status,omitempty"`

	// The status details of the certificate operation.
	StatusDetails *string `json:"status_details,omitempty"`

	// Location which contains the result of the certificate operation.
	Target *string `json:"target,omitempty"`

	// READ-ONLY; The certificate id.
	ID *string `json:"id,omitempty" azure:"ro"`
}

func certificateOperationFromGenerated(g generated.CertificateOperation) CertificateOperation {
	return CertificateOperation{
		CancellationRequested: g.CancellationRequested,
		Csr:                   g.Csr,
		Error:                 certificateErrorFromGenerated(g.Error),
		IssuerParameters:      issuerParametersFromGenerated(g.IssuerParameters),
		RequestID:             g.RequestID,
		Status:                g.Status,
		StatusDetails:         g.StatusDetails,
		Target:                g.Target,
		ID:                    g.ID,
	}
}

// CertificatePolicy - Management policy for a certificate.
type CertificatePolicy struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters `json:"issuer,omitempty"`

	// Properties of the key backing a certificate.
	KeyProperties *KeyProperties `json:"key_props,omitempty"`

	// Actions that will be performed by Key Vault over the lifetime of a certificate.
	LifetimeActions []*LifetimeAction `json:"lifetime_actions,omitempty"`

	// Properties of the secret backing a certificate.
	SecretProperties *SecretProperties `json:"secret_props,omitempty"`

	// Properties of the X509 component of a certificate.
	X509CertificateProperties *X509CertificateProperties `json:"x509_props,omitempty"`
}

func (c *CertificatePolicy) toGeneratedCertificateCreateParameters() *generated.CertificatePolicy {
	if c == nil {
		return nil
	}
	var la []*generated.LifetimeAction
	for _, l := range c.LifetimeActions {
		la = append(la, l.toGenerated())
	}

	return &generated.CertificatePolicy{
		Attributes:                c.Attributes.toGenerated(),
		IssuerParameters:          c.IssuerParameters.toGenerated(),
		KeyProperties:             c.KeyProperties.toGenerated(),
		SecretProperties:          c.SecretProperties.toGenerated(),
		LifetimeActions:           la,
		X509CertificateProperties: c.X509CertificateProperties.toGenerated(),
	}
}

func certificatePolicyFromGenerated(g *generated.CertificatePolicy) *CertificatePolicy {
	if g == nil {
		return nil
	}

	var la []*LifetimeAction
	for _, l := range g.LifetimeActions {
		la = append(la, lifetimeActionFromGenerated(l))
	}

	return &CertificatePolicy{
		Attributes:                certificateAttributesFromGenerated(g.Attributes),
		IssuerParameters:          issuerParametersFromGenerated(g.IssuerParameters),
		KeyProperties:             keyPropertiesFromGenerated(g.KeyProperties),
		LifetimeActions:           la,
		SecretProperties:          &SecretProperties{ContentType: g.SecretProperties.ContentType},
		X509CertificateProperties: x509CertificatePropertiesFromGenerated(g.X509CertificateProperties),
	}
}

// Contact - The contact information for the vault certificates.
type Contact struct {
	// Email address.
	EmailAddress *string `json:"email,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	// Phone number.
	Phone *string `json:"phone,omitempty"`
}

func contactListFromGenerated(g []*generated.Contact) []*Contact {
	var ret []*Contact
	for _, c := range g {
		ret = append(ret, &Contact{
			EmailAddress: c.EmailAddress,
			Name:         c.Name,
			Phone:        c.Phone,
		})
	}
	return ret
}

// Contacts - The contacts for the vault certificates.
type Contacts struct {
	// The contact list for the vault certificates.
	ContactList []*Contact `json:"contacts,omitempty"`

	// READ-ONLY; Identifier for the contacts collection.
	ID *string `json:"id,omitempty" azure:"ro"`
}

func (c *Contacts) toGenerated() generated.Contacts {
	if c == nil {
		return generated.Contacts{}
	}

	var contacts []*generated.Contact
	for _, contact := range c.ContactList {
		contacts = append(contacts, &generated.Contact{
			EmailAddress: contact.EmailAddress,
			Name:         contact.Name,
			Phone:        contact.Phone,
		})
	}

	return generated.Contacts{
		ID:          c.ID,
		ContactList: contacts,
	}
}

// DeletedCertificateBundle - A Deleted Certificate consisting of its previous id, attributes and its tags, as well as information on when it will be purged.
type DeletedCertificateBundle struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// CER contents of x509 certificate.
	Cer []byte `json:"cer,omitempty"`

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string `json:"contentType,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs
	Tags map[string]string `json:"tags,omitempty"`

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The certificate id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The key id.
	Kid *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy `json:"policy,omitempty" azure:"ro"`

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`

	// READ-ONLY; The secret id.
	Sid *string `json:"sid,omitempty" azure:"ro"`

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty" azure:"ro"`
}

// DeletedCertificateItem - The deleted certificate item containing metadata about the deleted certificate.
type DeletedCertificateItem struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Certificate identifier.
	ID *string `json:"id,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	// Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty"`

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// IssuerAttributes - The attributes of an issuer managed by the Key Vault service.
type IssuerAttributes struct {
	// Determines whether the issuer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

func issuerAttributesFromGenerated(g *generated.IssuerAttributes) *IssuerAttributes {
	if g == nil {
		return nil
	}

	return &IssuerAttributes{
		Enabled: g.Enabled,
		Created: g.Created,
		Updated: g.Updated,
	}
}

func (i *IssuerAttributes) toGenerated() *generated.IssuerAttributes {
	if i == nil {
		return nil
	}

	return &generated.IssuerAttributes{
		Enabled: i.Enabled,
		Created: i.Created,
		Updated: i.Updated,
	}
}

// CertificateIssuer - The issuer for Key Vault certificate.
type CertificateIssuer struct {
	// Attributes of the issuer object.
	Attributes *IssuerAttributes `json:"attributes,omitempty"`

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials `json:"credentials,omitempty"`

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`

	// The issuer provider.
	Provider *string `json:"provider,omitempty"`

	// READ-ONLY; Identifier for the issuer object.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// IssuerCredentials - The credentials to be used for the certificate issuer.
type IssuerCredentials struct {
	// The user name/account name/account id.
	AccountID *string `json:"account_id,omitempty"`

	// The password/secret/account key.
	Password *string `json:"pwd,omitempty"`
}

func issuerCredentialsFromGenerated(g *generated.IssuerCredentials) *IssuerCredentials {
	if g == nil {
		return nil
	}

	return &IssuerCredentials{AccountID: g.AccountID, Password: g.Password}
}

func (i *IssuerCredentials) toGenerated() *generated.IssuerCredentials {
	if i == nil {
		return nil
	}

	return &generated.IssuerCredentials{
		AccountID: i.AccountID,
		Password:  i.Password,
	}
}

// IssuerParameters - Parameters for the issuer of the X509 component of a certificate.
type IssuerParameters struct {
	// Indicates if the certificates generated under this policy should be published to certificate transparency logs.
	CertificateTransparency *bool `json:"cert_transparency,omitempty"`

	// Certificate type as supported by the provider (optional); for example 'OV-SSL', 'EV-SSL'
	CertificateType *string `json:"cty,omitempty"`

	// Name of the referenced issuer object or reserved names; for example, 'Self' or 'Unknown'.
	Name *string `json:"name,omitempty"`
}

func (i *IssuerParameters) toGenerated() *generated.IssuerParameters {
	if i == nil {
		return &generated.IssuerParameters{}
	}

	return &generated.IssuerParameters{
		CertificateTransparency: i.CertificateTransparency,
		CertificateType:         i.CertificateType,
		Name:                    i.Name,
	}
}

func issuerParametersFromGenerated(g *generated.IssuerParameters) *IssuerParameters {
	if g == nil {
		return nil
	}

	return &IssuerParameters{
		CertificateTransparency: g.CertificateTransparency,
		CertificateType:         g.CertificateType,
		Name:                    g.Name,
	}
}

// KeyProperties - Properties of the key pair backing a certificate.
type KeyProperties struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Curve *CertificateKeyCurveName `json:"crv,omitempty"`

	// Indicates if the private key can be exported.
	Exportable *bool `json:"exportable,omitempty"`

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"key_size,omitempty"`

	// The type of key pair to be used for the certificate.
	KeyType *CertificateKeyType `json:"kty,omitempty"`

	// Indicates if the same key pair will be used on certificate renewal.
	ReuseKey *bool `json:"reuse_key,omitempty"`
}

func (k *KeyProperties) toGenerated() *generated.KeyProperties {
	if k == nil {
		return nil
	}

	return &generated.KeyProperties{
		Curve:      (*generated.JSONWebKeyCurveName)(k.Curve),
		Exportable: k.Exportable,
		KeySize:    k.KeySize,
		KeyType:    (*generated.JSONWebKeyType)(k.KeyType),
		ReuseKey:   k.ReuseKey,
	}
}

func keyPropertiesFromGenerated(g *generated.KeyProperties) *KeyProperties {
	if g == nil {
		return nil
	}

	return &KeyProperties{
		Curve:      (*CertificateKeyCurveName)(g.Curve),
		Exportable: g.Exportable,
		KeySize:    g.KeySize,
		KeyType:    (*CertificateKeyType)(g.KeyType),
		ReuseKey:   g.ReuseKey,
	}
}

// LifetimeAction - Action and its trigger that will be performed by Key Vault over the lifetime of a certificate.
type LifetimeAction struct {
	// The action that will be executed.
	Action *Action `json:"action,omitempty"`

	// The condition that will execute the action.
	Trigger *Trigger `json:"trigger,omitempty"`
}

func (l LifetimeAction) toGenerated() *generated.LifetimeAction {
	return &generated.LifetimeAction{
		Action:  l.Action.toGenerated(),
		Trigger: l.Trigger.toGenerated(),
	}
}

func lifetimeActionFromGenerated(g *generated.LifetimeAction) *LifetimeAction {
	if g == nil {
		return nil
	}

	return &LifetimeAction{
		Action: &Action{
			ActionType: (*CertificatePolicyAction)(g.Action.ActionType),
		},
		Trigger: &Trigger{
			DaysBeforeExpiry:   g.Trigger.DaysBeforeExpiry,
			LifetimePercentage: g.Trigger.LifetimePercentage,
		},
	}
}

// OrganizationDetails - Details of the organization of the certificate issuer.
type OrganizationDetails struct {
	// Details of the organization administrator.
	AdminDetails []*AdministratorDetails `json:"admin_details,omitempty"`

	// Id of the organization.
	ID *string `json:"id,omitempty"`
}

func organizationDetailsFromGenerated(g *generated.OrganizationDetails) *OrganizationDetails {
	if g == nil {
		return nil
	}
	var ret []*AdministratorDetails

	for _, a := range g.AdminDetails {
		ret = append(ret, administratorDetailsFromGenerated(a))
	}

	return &OrganizationDetails{
		ID:           g.ID,
		AdminDetails: ret,
	}
}

func (o *OrganizationDetails) toGenerated() *generated.OrganizationDetails {
	if o == nil {
		return nil
	}

	var adminDetails []*generated.AdministratorDetails
	for _, a := range o.AdminDetails {
		adminDetails = append(adminDetails, a.toGenerated())
	}

	return &generated.OrganizationDetails{
		ID:           o.ID,
		AdminDetails: adminDetails,
	}
}

// SecretProperties - Properties of the key backing a certificate.
type SecretProperties struct {
	// The media type (MIME type).
	ContentType *string `json:"contentType,omitempty"`
}

func (s *SecretProperties) toGenerated() *generated.SecretProperties {
	if s == nil {
		return nil
	}

	return &generated.SecretProperties{
		ContentType: s.ContentType,
	}
}

// SubjectAlternativeNames - The subject alternate names of a X509 object.
type SubjectAlternativeNames struct {
	// Domain names.
	DNSNames []*string `json:"dns_names,omitempty"`

	// Email addresses.
	Emails []*string `json:"emails,omitempty"`

	// User principal names.
	Upns []*string `json:"upns,omitempty"`
}

func (s *SubjectAlternativeNames) toGenerated() *generated.SubjectAlternativeNames {
	if s == nil {
		return &generated.SubjectAlternativeNames{}
	}

	return &generated.SubjectAlternativeNames{
		DNSNames: s.DNSNames,
		Emails:   s.Emails,
		Upns:     s.Upns,
	}
}

func subjectAlternativeNamesFromGenerated(g *generated.SubjectAlternativeNames) *SubjectAlternativeNames {
	if g == nil {
		return nil
	}

	return &SubjectAlternativeNames{
		DNSNames: g.DNSNames,
		Emails:   g.Emails,
		Upns:     g.Upns,
	}
}

// Trigger - A condition to be satisfied for an action to be executed.
type Trigger struct {
	// Days before expiry to attempt renewal. Value should be between 1 and validityinmonths multiplied by 27. If validityinmonths is 36, then value should
	// be between 1 and 972 (36 * 27).
	DaysBeforeExpiry *int32 `json:"days_before_expiry,omitempty"`

	// Percentage of lifetime at which to trigger. Value should be between 1 and 99.
	LifetimePercentage *int32 `json:"lifetime_percentage,omitempty"`
}

func (t *Trigger) toGenerated() *generated.Trigger {
	if t == nil {
		return nil
	}

	return &generated.Trigger{
		DaysBeforeExpiry:   t.DaysBeforeExpiry,
		LifetimePercentage: t.LifetimePercentage,
	}
}

// X509CertificateProperties - Properties of the X509 component of a certificate.
type X509CertificateProperties struct {
	// The enhanced key usage.
	Ekus []*string `json:"ekus,omitempty"`

	// List of key usages.
	KeyUsage []*CerificateKeyUsage `json:"key_usage,omitempty"`

	// The subject name. Should be a valid X509 distinguished Name.
	Subject *string `json:"subject,omitempty"`

	// The subject alternative names.
	SubjectAlternativeNames *SubjectAlternativeNames `json:"sans,omitempty"`

	// The duration that the certificate is valid in months.
	ValidityInMonths *int32 `json:"validity_months,omitempty"`
}

func (x *X509CertificateProperties) toGenerated() *generated.X509CertificateProperties {
	if x == nil {
		return &generated.X509CertificateProperties{}
	}

	var keyUsage []*generated.KeyUsageType
	for _, k := range x.KeyUsage {
		keyUsage = append(keyUsage, (*generated.KeyUsageType)(k))
	}

	return &generated.X509CertificateProperties{
		Ekus:                    x.Ekus,
		KeyUsage:                keyUsage,
		Subject:                 x.Subject,
		SubjectAlternativeNames: x.SubjectAlternativeNames.toGenerated(),
		ValidityInMonths:        x.ValidityInMonths,
	}
}

func x509CertificatePropertiesFromGenerated(g *generated.X509CertificateProperties) *X509CertificateProperties {
	if g == nil {
		return nil
	}

	var ku []*CerificateKeyUsage
	for _, k := range g.KeyUsage {
		ku = append(ku, (*CerificateKeyUsage)(k))
	}

	return &X509CertificateProperties{
		Ekus:                    g.Ekus,
		Subject:                 g.Subject,
		KeyUsage:                ku,
		SubjectAlternativeNames: subjectAlternativeNamesFromGenerated(g.SubjectAlternativeNames),
		ValidityInMonths:        g.ValidityInMonths,
	}
}

func convertToGeneratedMap(m map[string]string) map[string]*string {
	if m == nil {
		return nil
	}

	ret := make(map[string]*string)
	for k, v := range m {
		ret[k] = &v
	}
	return ret
}

func convertGeneratedMap(m map[string]*string) map[string]string {
	if m == nil {
		return nil
	}

	ret := make(map[string]string)
	for k, v := range m {
		ret[k] = *v
	}
	return ret
}
