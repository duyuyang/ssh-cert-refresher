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

// userDriver stores interfaces used for User Certificate implementation
type userDriver struct {
	iUserCAKey    userCAKey
	iTrustedCerts trustedCerts
}

// hostDriver stores interfacs used for Host Certificate implementation
type hostDriver struct {
	iSSHKeyGenerator sshKeyGenerator
	iCertSigner      certSigner
}

// driver stores interfaces used for generic implementations
type driver struct {
	iSSHdConfiger    sshdConfiger
	iServerValidator serverValidator
}

// userCAKey is the interface to retrieve CA Public Key
type userCAKey interface {
	getUserCAKey() (string, error)
}

// trustedCerts is the interface to write cert to the server
type trustedCerts interface {
	useTrustedCerts() error
}

// sshdConfiger is the interface to edit /etc//sshd_config
type sshdConfiger interface {
	ensureSSHdCfg() error
	restartSSHd() error
}

// certRefresher is the interface to loop the refresher
type certRefresher interface {
	refreshCert()
}

// sshKeyGenerator is the interface to generate SSH key pairs
type sshKeyGenerator interface {
	genSSHKeys() (priKey string, pubKey string)
}

// certSigner is the interface to sign the public key
type certSigner interface {
	signPubKey() (cert string, err error)
}

// sshdRestarter is the interface to restart sshd service
// type sshdRestarter interface {
// 	restartSSHd() error
// }

// serverValidator is the interface to run basic checks on the server
type serverValidator interface {
	validateServer() error
}
