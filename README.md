Go gRPC Test Project
===

### Description
User Management Microservice with Go, gRPC, and implement Authorization and Authentication.

### Prerequisites
- Go (version 1.25 or newer is recommended)
- PostgreSQL
- Redis
- Docker
- Protocol Buffers compiler ( protoc )
```bash
# Install protocol buffers compiler
sudo apt-get install -y protobuf-compiler # Ubuntu/Debian
```
- Go plugins for protoc
```bash
# Install Go plugins for the protocol compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Getting Started
These instructions will get a copy of the project up and running on your local machine for development and testing purpose

#### Install Dependencies
First, clone the repository and navigate into the project directory:
```bash
git clone https://github.com/adhyttungga/go-grpc-test.git
cd go-grpc-test
```
#### Generate Go Code from Protobuf
gRPC uses Protocol Buffers to define service methods and messages. You will need to generate the Go code from the ".proto" files, you can run the following command:
```bash
# Generate Go code from the proto files
protoc --proto_path=internal/proto \
    --go_out=internal/pb \
    --go-grpc_out=internal/pb \
    auth.proto user.proto
```

### Running the Application
#### Run the Server
To start the gRPC server, execute the following command:
```bash
./scripts/start.sh
```
The server will start as docker container and listen on a specified port (e.g., ":8080")

### Project Structure
This project structure is as follows:
```
.
├── api/
│   └── v1/
├── cmd/
├── internal/
│   ├── models/
│   │   └── entity/
│   ├── pb/
│   ├── proto/
│   ├── repository/
│   └── usecase/
├── pkg/
│   ├── config/
│   ├── middleware/
│   ├── seeder/
│   └── utils/
├── scripts/
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```
