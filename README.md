# testgolangrpc2
other version of golang grpc without heroku to be executed locally


# locally


git clone https://github.com/czendee/testgolangrpc2.git

1. git clone https://github.com/czendee/testgolangrpc2.git


2. cd testgolangrpc2

3.1 (o) go get github.com/czendee/testgolangrpc2/genproto/go

3.2 (o) cd genproto/go 

3.3  (o)  go build 

3.4 (o) cd ..
        cd ..

4. go build main.go

5. go run  main.go

     running on port 8090
     
## Test using Postman


### example:


http://127.0.0.1:8080/v2/multiplica/greeter/say-multiplica

method POST
using Postman:   raw json

{
   "numero" :"18",
    "veces" :"20"
}

result

{
    "message": "360.000000"
}


### test case:


http://127.0.0.1:8080/v2/multiplica/greeter/say-multiplica

method POST
using Postman:   raw json

{
   "numero" :"1",
    "veces" :"2"
}

result

{
    "message": "2.000000"
}

### test case:


http://127.0.0.1:8080/v2/multiplica/greeter/say-multiplica

method POST
using Postman:   raw json

{
   "numero" :"0",
    "veces" :"2"
}

result

{
    "message": "2.000000"
}


### test case:


http://127.0.0.1:8080/v2/multiplica/greeter/say-multiplica

method POST
using Postman:   raw json

{
   "numero" :"0",
    "veces" :"0"
}

result

{
    "message": "2.000000"
}



### test case:


http://127.0.0.1:8080/v2/multiplica/greeter/say-multiplica

method POST
using Postman:   raw json

{
   "numero" :"ANA",
    "veces" :"2"
}

result

{
    "message": "2.000000"
}







# When the multiplica.proto is modified , execute these steps 

## the protoc to generate the pb.go, gb.go for the server and the client





1. protoc -I ./proto  -I .   --go_out ./genproto/go --go_opt paths=source_relative    --go-grpc_out ./genproto/go --go-grpc_opt paths=source_relative   --grpc-gateway_out ./genproto/go  --grpc-gateway_opt paths=source_relative proto/*.proto



2. go mod init github.com/czendee/testgolangrpc2/genproto/go

3.  cd genproto/go 

4.    go build 


5.   edit the multiplica_grpc.pb

     5.1  change this package _ for this package __
     
     5.2  comment 
    
//	codes "google.golang.org/grpc/codes"

//	status "google.golang.org/grpc/status"

//	mustEmbedUnimplementedGreeterServer()

    5.2 remove these lines

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayMultiplica(context.Context, *MultiplicaRequest) (*MultiplicaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayMultiplica not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

6. 
cd ..
cd ..


after creating it we need to do a 

git add .

git commit -m "adding the new go files generated by protoc"

git push


