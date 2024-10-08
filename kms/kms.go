// Copyright (C) 2022-2023 Nicola Murino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kms provides Key Management Services support
package kms

// SecretStatus defines the statuses of a Secret object
type SecretStatus = string

const (
	// SecretStatusPlain means the secret is in plain text and must be encrypted
	SecretStatusPlain SecretStatus = "Plain"
	// SecretStatusAES256GCM means the secret is encrypted using AES-256-GCM
	SecretStatusAES256GCM SecretStatus = "AES-256-GCM"
	// SecretStatusSecretBox means the secret is encrypted using a locally provided symmetric key
	SecretStatusSecretBox SecretStatus = "Secretbox"
	// SecretStatusGCP means we use keys from Google Cloud Platform’s Key Management Service
	// (GCP KMS) to keep information secret
	SecretStatusGCP SecretStatus = "GCP"
	// SecretStatusAWS means we use customer master keys from Amazon Web Service’s
	// Key Management Service (AWS KMS) to keep information secret
	SecretStatusAWS SecretStatus = "AWS"
	// SecretStatusVaultTransit means we use the transit secrets engine in Vault
	// to keep information secret
	SecretStatusVaultTransit SecretStatus = "VaultTransit"
	// SecretStatusAzureKeyVault means we use Azure KeyVault to keep information secret
	SecretStatusAzureKeyVault SecretStatus = "AzureKeyVault"
	// SecretStatusOracleKeyVault means we use Oracle Key Vault to keep information secret
	SecretStatusOracleKeyVault SecretStatus = "OracleKeyVault"
	// SecretStatusRedacted means the secret is redacted
	SecretStatusRedacted SecretStatus = "Redacted"
)

// Scheme defines the supported URL scheme
type Scheme = string

// supported URL schemes
const (
	SchemeLocal          Scheme = "local"
	SchemeBuiltin        Scheme = "builtin"
	SchemeAWS            Scheme = "awskms"
	SchemeGCP            Scheme = "gcpkms"
	SchemeVaultTransit   Scheme = "hashivault"
	SchemeAzureKeyVault  Scheme = "azurekeyvault"
	SchemeOracleKeyVault Scheme = "ocikeyvault"
)

// BaseSecret defines the base struct shared among all the secret providers
type BaseSecret struct {
	Status         SecretStatus `json:"status,omitempty"`
	Payload        string       `json:"payload,omitempty"`
	Key            string       `json:"key,omitempty"`
	AdditionalData string       `json:"additional_data,omitempty"`
	// 1 means encrypted using a master key
	Mode int `json:"mode,omitempty"`
}
