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
	"io/ioutil"
	"log"
	"net"
	"os"
)

// dnsCA stores CA key pairs in DNS txt record
type dnsCA struct {
	DNS string
	err error
}

// getUserCAKey implement userCAKey to retrieve user CA public key from DNS txt record
func (dc *dnsCA) getUserCAKey() (string, error) {
	// Assume this DNS record stores one public key only
	log.Printf("pull cert from DNS txt: %v", dc.DNS)
	if txtrecords, err := net.LookupTXT(dc.DNS); err != nil {
		return "", errors.New("Failed to reach DNS record")
	} else {
		if len(txtrecords) > 0 {
			return txtrecords[0], nil
		}
	}
	return "", errors.New("DNS record is empty")
}

// paramCA stores CA key pairs in AWS Parameter Store
type paramCA struct {
	err error
}

// getUserCAkey implement userCAKey to retrieve user CA public from AWS parameter store
func (dc *paramCA) getUserCAKey() (string, error) {
	log.Println("pull cert from AWS parameter store")
	return "", errors.New("Not implement yet")
}

// cert stores certificate data
type userCert struct {
	data string
	err  error
}

// useTrustedCerts implement `trustedCerts` write the User CA Pub Key into /etc/ssh/trusted_certs
func (c *userCert) useTrustedCerts() error {
	if c.err != nil {
		return errors.New("Error with User Cert")
	}
	b := []byte(c.data + "\n")
	err := ioutil.WriteFile(sshdCfgPathMnt+"trusted_cert", b, 0644)
	if err != nil {
		log.Printf("error write to trusted_cert %v", err)
		return err
	}
	log.Printf("Override trusted_cert with %v", c.data)

	return nil
}

// setData implement `trustedCerts` to update cert object
func (c *userCert) setCert(data string, err error) {
	c.data = data
	c.err = err
}

// sshdConfig stores sshd_config required content
type userSSHdConfig struct {
	sshPath    string
	sshMntPath string
	file       string
	content    string
}

func (u *userSSHdConfig) setUserSSHdConfig(sshPath string, sshMntPath string, file string, content string) {
	u.sshPath = sshPath
	u.sshMntPath = sshMntPath
	u.file = file
	u.content = content
}

// ensureSSHdCfg implement `sshdConfiger` to edit sshd_config
func (u *userSSHdConfig) ensureSSHdCfg() error {
	// Assume run this function once
	log.Println("Add TrustedUserCAKeys /etc/ssh/trusted_cert to " + u.sshPath + "sshd_config")
	file, err := os.OpenFile(u.sshMntPath+u.file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(u.content + u.sshPath + u.file + "\n")); err != nil {
		return errors.New("Failed to edit sshd_config")
	}

	return nil
}
