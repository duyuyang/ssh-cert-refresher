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

// Implement the Refresher

// ------------

import (
	"log"
	"time"
)

// sshdConfig Implement the ensureSSHdCfg function
func sshdConfig(sc sshdConfiger) {
	sc.ensureSSHdCfg()
}

// refresh Implement the refreshCert function
func refresh(re certRefresher) {
	re.ensureSSHdCfg()
	re.refreshCert()
}

// server stores services required running on the server
type server struct {
	service string
	ensure  string
}

// validateServer implements the interface `serverValidator`
func (svr *server) validateServer() error {
	return nil
}

// restartSSHd implements the interface `sshdRestarter`
func (svr *server) restartSSHd() error {
	return nil
}

// defaultRefresher stores the data required by certRefresher
type defaultDriver struct {
	driver     *driver
	userDriver *userDriver
}

func (d *defaultDriver) ensureSSHdCfg() error {
	// Assume run this function once

	return nil
}

// refreshCert implements the interface `certRefresher`
func (d *defaultDriver) refreshCert() {

	// refreshUserCert is an infinite loop delay time.Duration
	for {
		log.Println("looping through refreshing the cert")

		// get the ca kay to trust file
		cert, err := d.userDriver.iUserCAKey.getUserCAKey()
		if err != nil {
			log.Printf("Error Get User CA key, %v", err)
			d.userDriver.iTrustedCerts.setCert("", err)
		}
		d.userDriver.iTrustedCerts.setCert(cert, nil)
		d.userDriver.iTrustedCerts.useTrustedCerts()

		// restart ssh

		time.Sleep(time.Second * 2)
	}
}

// enhancedRefresher stores the data required for both user and host certificate
type enhancedDriver struct {
	//driver     *driver
	userDriver *userDriver
	hostDriver *hostDriver
}

// refreshCert implements the interface `certRefresher`
func (ed *enhancedDriver) refreshCert() {}

// ensureSSHdCfg
func (ed *enhancedDriver) ensureSSHdCfg() error {
	return nil
}
