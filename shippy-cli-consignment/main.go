package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	// import the auto generate protoc file in the consignment service
	pb "github.com/celelstine/shippy/shippy-service-consignment/proto/consignment"
	micro "github.com/micro/go-micro/v2"
)

const (
	// for docker change localhost to `docker-machine ip`
	// address = "192.168.99.100:50051"
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

// parseFile local method to parse a file to an consignment object
func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("obj: %vCreated: %t", r, r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for index, v := range getAll.Consignments {
		log.Println(index, v)
	}
}
