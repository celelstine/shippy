package main

import (
	"context"
	"log"

	// Import the generated protobuf code
	pb "github.com/celelstine/shippy/shippy-service-consignment/proto/consignment"

	micro "github.com/micro/go-micro/v2"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository create a struct for hosting the data
type Repository struct {
	// use mutex to lock resources while we use them
	// we need this to manually lock the resources, micro would handle this
	// mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create implement the create method of the repository interface to create consigment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	// repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	// repo.mu.Unlock()
	return consignment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// define the service which would use the repository
type consignmentService struct {
	repo repository
}

// implement the create consigment method
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	// return &pb.Response{Created: true, Consignment: consignment}, nil

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	// create an empty repository
	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	service := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register service
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
