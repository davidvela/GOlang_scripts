# gRPC
Google Remote Procedure Call
https://gRPC.io/

muchasfolders  y files 
Secure connection and protobuf libraries (faster)


## Message: 
message Repository {
	int64 id  = 1;
	string name = 2;
	int64 userId = 3;
	bool isPrivate = 4;
}


## gRPC Service



# Layout
build
   → proto-gen.bat
   → proto-gen.sh
cmd
   → gRPC
           → server
               → main.go
           → client
               → main.go
Internal
   → gRPC
   → proto-files
           → domain
                   → repository.proto
           → service
                   → repository-service.proto
pkg


