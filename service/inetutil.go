package service

import (
	"log"
	"net"
)

// Adapted from https://stackoverflow.com/a/37382208
func getOutBoundAddr() net.IP {
	// Get preferred outbound ip of this machine
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
