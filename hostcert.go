// Copyright © 2019 Du Yuyang.
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

package main

// Implement Host Certificate required interfaces

// ------------

// Implement the interface `sshKeyGenerator` to generate host key pairs

// Implement the interface `certSigner` to sign the key
// require Host CA to sign the public key
// return the host cert

// Implement the interface `trustedCerts` to write cert to server
// output the host cert to /etc//host-cert.pub

// Implement the interface `sshdConfiger` to ensure following exists in /etc/sshd_config
// HostCertificate /etc/ssh/host-cert.pub
