
proto_kafka_messages:
	@echo Generating kafka messages proto
	cd proto/kafka_messages && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. kafka.proto

# proto auth
proto_auth:
	@echo Generating auth microservice proto
	cd authentication/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. auth_account.proto

proto_auth_message:
	@echo Generating auth messages microservice proto
	cd authentication/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. auth_message.proto

# proto account
proto_account:
	@echo Generating account microservice proto
	cd account/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. account.proto

proto_account_message:
	@echo Generating account messages microservice proto
	cd account/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. account_message.proto


swagger_gateway:
	@echo Starting swagger generating
	cd gateway && swag init -g ./cmd/main.go