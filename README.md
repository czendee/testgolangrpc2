# testgolangrpc2
other version of golang grpc without heroku to be executed locally


# locally


git clone https://github.com/czendee/testgolangrpc2.git


# execute the protoc to generate the pb.go, gb.go for the server and the client

protoc -I ./proto  -I .   --go_out ./proto --go_opt paths=source_relative    --go-grpc_out ./genproto/goclient --go-grpc_opt paths=source_relative   --grpc-gateway_out ./genproto/goserver  --grpc-gateway_opt paths=source_relative proto/*.proto
