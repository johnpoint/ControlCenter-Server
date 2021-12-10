protoc --go_out=../ proto/server_info/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/server_info/*.proto
