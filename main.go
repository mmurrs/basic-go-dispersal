package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Layr-Labs/eigenda/api/clients/v2"
	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser/v2"
	"github.com/Layr-Labs/eigenda/core"
	corev2 "github.com/Layr-Labs/eigenda/core/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := &clients.DisperserClientConfig{
		Hostname:          "localhost",
		Port:              "50051",
		UseSecureGrpcFlag: false,
	}

	client, err := NewDisperserClient(config)
	if err != nil {
		log.Fatalf("Failed to create disperser client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	data := []byte("test data")
	blobVersion := corev2.BlobVersion(1)
	quorums := []core.QuorumID{1, 2, 3}
	salt := uint32(12345)

	blobStatus, blobKey, err := client.DisperseBlob(ctx, data, blobVersion, quorums, salt)
	if err != nil {
		log.Fatalf("Failed to disperse blob: %v", err)
	}

	fmt.Printf("Blob Status: %v, Blob Key: %v\n", blobStatus, blobKey)
}

func NewDisperserClient(config *clients.DisperserClientConfig) (clients.DisperserClient, error) {
	var opts []grpc.DialOption
	if config.UseSecureGrpcFlag {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Hostname, config.Port), opts...)
	if err != nil {
		return nil, err
	}

	client := disperser_rpc.NewDisperserClient(conn)
	return &clients.disperserClient{
		config: config,
		conn:   conn,
		client: client,
	}, nil
}
