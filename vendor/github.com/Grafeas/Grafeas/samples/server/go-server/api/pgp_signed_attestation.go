/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// An attestation wrapper with a PGP-compatible signature. This message only supports ATTACHED signatures, where the payload that is signed is included alongside the signature itself in the same file.
type PgpSignedAttestation struct {
	// The raw content of the signature, as output by gpg or equivalent.  Since this message only supports attached signatures, the payload that was signed must be attached. While the signature format supported is dependent on the verification implementation, currently only ASCII-armored (`--armor` to gpg), non-clearsigned (`--sign` rather than `--clearsign` to gpg) are supported. Concretely, `gpg --sign --armor --output=signature.gpg payload.json` will create the signature content expected in this field in `signature.gpg` for the `payload.json` attestation payload.
	Signature string `json:"signature,omitempty"`

	// Type (e.g. schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (e.g. by validating a JSON schema).
	ContentType string `json:"contentType,omitempty"`

	// The ID of the key, as output by `gpg --list-keys`.  This should be 8 hexidecimal digits, capitalized.  e.g. $ gpg --list-keys pub 2048R/A663AEEA 2017-08-01 ui Fake Name <example-attesting-user@google.com> In the above example, the `key_id` is \"A663AEEA\". Note that in practice this ID is the last 64 bits of the key fingerprint.
	PgpKeyId string `json:"pgpKeyId,omitempty"`
}
