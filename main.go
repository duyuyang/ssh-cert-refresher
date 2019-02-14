package main

import (
	"errors"
	"log"
	"time"
)

// getUserCAKey retrive the DNS response
func getUserCAKey(DNS string) (string, error) {
	log.Println("get cert")
	return "cert", nil
}

// trustedCerts write the User CA Pub Key into /etc/ssh/trusted_certs
func trustedCerts(cert string, err error) error {
	if err != nil {
		return errors.New("not getting response")
	}
	log.Println("write trusted cert")
	return nil
}

// refreshUserCert is an infinite loop delay time.Duration
func refreshUserCert(DNS string, jobTimeout string, jobInterval string) {
	for {
		if r := trustedCerts(getUserCAKey(DNS)); r != nil {
			log.Printf("Failed to write trusted cert, %v", r)
		}
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var (
		certDNSRecord = ""
		jobTimeout    = ""
		jobInterval   = ""
	)

	refreshUserCert(certDNSRecord, jobTimeout, jobInterval)
}
