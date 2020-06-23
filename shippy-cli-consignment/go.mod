module github.com/celelstine/shippy/shippy-cli-consignment

go 1.14

require (
	github.com/celelstine/shippy/shippy-service-consignment v0.0.0-20200611190251-9bb42b902954
	github.com/celelstine/shippy/shippy-service-vessel v0.0.0-20200611220908-755556816183
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.3.4
	google.golang.org/grpc v1.29.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
