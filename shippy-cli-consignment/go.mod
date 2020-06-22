module github.com/celelstine/shippy/shippy-cli-consignment

go 1.14

require (
	github.com/celelstine/shippy/shippy-service-consignment v0.0.0-20200611190251-9bb42b902954
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	golang.org/x/tools v0.0.0-20200609164405-eb789aa7ce50 // indirect
	google.golang.org/grpc v1.29.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
