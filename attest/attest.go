// Copyright (c) Facebook, Inc. and its affiliates.
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

package attest

import (
	"io"

	"github.com/google/go-attestation/attest"
)

// EKData contains metadata for a TPM 2.0 Endorsement Key
type EKData struct {
	Certificate                 []byte // Complete ASN.1 DER content.
	PublicKey                   []byte
	IssuerCN                    string
	SubjectCN                   string
	SerialNumber                string
	SignatureAlgorithm          string
	PublicKeyAlgorithm          string
	HasCertInNVRAM              bool
	HasPublicKeyInNVRam         bool
	CertDownloadedFromVendorURL bool // this is expected to be false as we aren't downloading certificates yet
	VendorCertificateURL        string
}

// SecureHardwareVendorData represents metadata for the specific hardware backed key store available
// on the device
type SecureHardwareVendorData struct {
	EKs                    []EKData
	IsTPM20CompliantDevice bool
	VendorName             string
	VendorInfo             string
	Version                uint8
}

// Req represents the request to attest & certify a TPM key
type Req struct {
	AttestTPMHandle    *attest.TPM
	TransientKeyHandle *attest.Key
	TransientAKHandle  *attest.AK
	TPM                io.ReadWriteCloser
	KeyHandle          any
}

// Resp represents the response from the attestation process
type Resp struct {
	AttestationStatement string
	CertificationParams  string
	PublicKey            []byte
}

// Attestor is the interface which performs attestation
type Attestor interface {
	Attest(*Req) (*Resp, error)
}
