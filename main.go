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

var (
	certDNSRecord    string
	jobTimeout       string
	jobInterval      string
	refresherType    string
	sshdCfgPath      string
	sshdCfgPathMnt   string
	sshdCfgFile      string
	trustedCertsFile string
)

// init functions collect configuration data
func init() {
	refresherType = "default"
	certDNSRecord = "sshephalopod-ca-cert.cd2e-hub.realestate.com.au"
	sshdCfgPath = "/etc/ssh/"
	sshdCfgPathMnt = "/tmp/" // /host/etc/ssh/
	sshdCfgFile = "sshd_config"
	trustedCertsFile = "trusted_certs"
}

func main() {

	switch refresherType {
	case "default":
		refresh(&defaultDriver{
			driver: &driver{
				iSSHdConfiger: &userSSHdConfig{},
			},
			userDriver: &userDriver{
				iUserCAKey: &dnsCA{
					DNS: certDNSRecord,
				},
				iTrustedCerts: &userCert{},
			},
		})
	case "enhanced":
		refresh(&enhancedDriver{})
	}
}
