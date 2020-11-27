package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type X509V3ExtensionsType struct {
	BasicConstraints string `json:"basic_constraints,omitempty"`
	NameConstraints string `json:"name_constraints,omitempty"`
	PolicyConstraints string `json:"policy_constraints,omitempty"`
	KeyUsage string `json:"key_usage,omitempty"`
	ExtendedKeyUsage string `json:"extended_key_usage,omitempty"`
	SubjectKeyIdentifier string `json:"subject_key_identifier,omitempty"`
	AuthorityKeyIdentifier string `json:"authority_key_identifier,omitempty"`
	SubjectAlternativeName string `json:"subject_alternative_name,omitempty"`
	IssuerAlternativeName string `json:"issuer_alternative_name,omitempty"`
	SubjectDirectoryAttributes string `json:"subject_directory_attributes,omitempty"`
	CrlDistributionPoints string `json:"crl_distribution_points,omitempty"`
	InhibitAnyPolicy string `json:"inhibit_any_policy,omitempty"`
	PrivateKeyUsagePeriodNotBefore time.Time `json:"private_key_usage_period_not_before,omitempty"`
	PrivateKeyUsagePeriodNotAfter time.Time `json:"private_key_usage_period_not_after,omitempty"`
	CertificatePolicies string `json:"certificate_policies,omitempty"`
	PolicyMappings string `json:"policy_mappings,omitempty"`
}
type X509Certificate struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	IsSelfSigned bool `json:"is_self_signed,omitempty"`
	Hashes map[string]interface{} `json:"hashes,omitempty"`
	Version string `json:"version,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	SignatureAlgorithm string `json:"signature_algorithm,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	ValidityNotBefore time.Time `json:"validity_not_before,omitempty"`
	ValidityNotAfter time.Time `json:"validity_not_after,omitempty"`
	Subject string `json:"subject,omitempty"`
	SubjectPublicKeyAlgorithm string `json:"subject_public_key_algorithm,omitempty"`
	SubjectPublicKeyModulus string `json:"subject_public_key_modulus,omitempty"`
	SubjectPublicKeyExponent string `json:"subject_public_key_exponent,omitempty"`
	X509V3Extensions X509V3ExtensionsType `json:"x509_v3_extensions,omitempty"`
}

func printX509Certificate(x509Certificate X509Certificate) {
	fmt.Println("X509 Certificate:\n", marshalX509Certificate(x509Certificate))
}

func marshalX509Certificate(x509Certificate X509Certificate) (jsonData string){
	data, e := json.MarshalIndent(x509Certificate, "", "  ")
	check(e)
	return string(data)
}

func unmarshalX509Certificate(obj json.RawMessage) (x509Certificate X509Certificate) {
	json.Unmarshal(obj, &x509Certificate)
	return x509Certificate
}
