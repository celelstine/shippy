## shippy-service-consignment

Service to management consigment

## tips
- if you wanna use docker, ensure to install the service in the servive directory.

## setup with docker
https://docs.docker.com/engine/userguide/eng-image/multistage-build/#name-your-build-stages
We used a multi phase build with docker to make the image smaller, it contains 2 step:
- build the service 
- copy only the build file and run it 
- for micro version , run `docker run -p 50051:50051  -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment`

## using micro to build the microservice
- we need to set the host address as MICRO_SERVER_ADDRESS=:<port>