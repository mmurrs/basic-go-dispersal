package main

import (
	"github.com/Layr-Labs/eigenda/api/clients/v2"
)

func main() {
	config := &clients.DisperserClientConfig{
		Hostname:          "localhost",
		Port:              "50051",
		UseSecureGrpcFlag: false,
	}
}
