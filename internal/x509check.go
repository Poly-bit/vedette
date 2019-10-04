package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"
)

func getCerts(port int, ip string) ([]*x509.Certificate, error) {

	address := fmt.Sprintf("%s:%v", ip, port)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	ipConn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return nil, fmt.Errorf("error on dial to '%s' : %v", address, err)
	}
	defer ipConn.Close()

	conn := tls.Client(ipConn, tlsConfig)
	defer conn.Close()

	if err := conn.Handshake(); err != nil {
		return nil, fmt.Errorf("error on ssl handshake with '%s' : %v", address, err)
	}

	certs := conn.ConnectionState().PeerCertificates
	fmt.Printf("found %v certs for: '%s'\n", len(certs), address)
	return certs, nil
}

func calcExpiry(cert x509.Certificate) int64 {
	now := time.Now()
	return int64(cert.NotAfter.Sub(now).Seconds())
}
