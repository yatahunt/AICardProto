


sudo apt-get install protobuf-compiler
go install google.golang.org/protobuf@latest
go install google.golang.org/grpc@latest
 
go install github.com/yatahunt/PokerDeck@latest

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative   AICardProto/AICardProto.proto

pip install grpcio-tools
pip install googleapis-common-protos
python -m grpc_tools.protoc --proto_path=. --python_out=./AICardProto/Python --grpc_python_out=./AICardProto/Python AICardProto/AICardProto.proto
cd .\AICardProto\ 
del .\Go\AICardProto\AICardProto.pb.go
del .\Go\AICardProto\AICardProto_grpc.pb.go 
move .\AICardProto.pb.go .\Go\AICardProto
move .\AICardProto_grpc.pb.go .\Go\AICardProto
cd ..