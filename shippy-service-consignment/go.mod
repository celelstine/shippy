module github.com/celelstine/shippy/shippy-service-consignment

go 1.14

require (
	github.com/celelstine/shippy/shippy-service-vessel v0.0.0-20200611220908-755556816183
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/micro/v2 v2.8.1-0.20200608175027-6e673c965a19 // indirect
	golang.org/x/sys v0.0.0-20200620081246-981b61492c35 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200620020550-bd6e04640131 // indirect
	google.golang.org/grpc v1.29.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
