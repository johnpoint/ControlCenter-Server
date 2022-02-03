rm proto/controlproto/*.go
protoc --go_out=../ proto/controlproto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/controlproto/*.proto

rm proto/mqproto/*.go
protoc --go_out=../ proto/mqproto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/mqproto/*.proto

rm proto/tcpproto/*.go
protoc --go_out=../ proto/tcpproto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/tcpproto/*.proto
