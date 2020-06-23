package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/celelstine/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/celelstine/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	// get db uri or use default
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	// use get db client from datastore
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	// get the consignment collection
	consignmentCollection := client.Database("shippy").Collection("consignments")

	// get the model repository
	repository := &MongoRepository{consignmentCollection}

	// get the vessel client and create a handler via handler.go
	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}