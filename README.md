# grpc-test
Very simple gRPC server and client written in Golang/Python

## Server side
This is a very simple gRPC server Golang implementation with two endpoints: getting an user and creating one.

In order to generate the golang protobuf and gRPC files you can run the following command while in the `server` dir:
```
protoc --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative internal.proto --proto_path=<path to the proto/user folder>
```

To run the golang server you can run:
```
go run main.go
```
The server will be accessible on the `5050` port.

To do 'curl' like commands to the grpc server the `grpcurl` tool (https://github.com/fullstorydev/grpcurl) can be used. An example command to get a user looks like this:
```
grpcurl -plaintext -proto <path to the proto definition> -d '{"user_id": "XVlBzgbaiCMRAjWwhTHc"}' <server_host:port> user.UserInternalService/GetUser
```

## Client side
This is a very simple HTTP to gRPC proxy done using FastAPI for the HTTP part in Python.

In order to generate the proto files you can run the following:
```
python -m grpc_tools.protoc -I=../proto --python_out=./proto --grpc_python_out=./proto <path to the proto definition>
```

To run the HTTP API you can do:
```
uvicorn main:app
```
While being in the `client` dir. The API will be accessible on the `8000` port.

The `requirements.txt` file has all the current dependencies taken using the `pip freeze` command.