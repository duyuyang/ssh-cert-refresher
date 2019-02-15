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

// Implement User Certificate required interfaces

// ------------

import (
	"errors"
	"log"
	"net"
)

// dnsCA stores CA key pairs in DNS txt record
type dnsCA struct {
	DNS string
	err error
}

// getUserCAKey implement userCAKey to retrieve user CA public key from DNS txt record
func (dc *dnsCA) getUserCAKey() (string, error) {
	log.Printf("pull cert from DNS txt: %v", dc.DNS)
	txtrecords, _ := net.LookupTXT(dc.DNS)

	for _, txt := range txtrecords {
		log.Println(txt)
	}
	return "cert", nil
}

// paramCA stores CA key pairs in AWS Parameter Store
type paramCA struct {
	err error
}

// getUserCAkey implement userCAKey to retrieve user CA public from AWS parameter store
func (dc *paramCA) getUserCAKey() (string, error) {
	log.Println("pull cert from AWS parameter store")
	return "cert", nil
}

// cert stores certificate data
type userCert struct {
	data string
	err  error
}

// useTrustedCerts implement `trustedCerts` write the User CA Pub Key into /etc/ssh/trusted_certs
func (c *userCert) useTrustedCerts() error {
	if c.err != nil {
		return errors.New("not getting response")
	}
	log.Println("write trusted cert")
	return nil
}

// sshdConfig stores sshd_config required content
type userSSHdConfig struct {
	path    string
	file    string
	content string
	err     error
}

// ensureSSHdCfg implement `sshdConfiger` to edit sshd_config
func (u *userSSHdConfig) ensureSSHdCfg() error {
	return nil
}
