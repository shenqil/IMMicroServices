protoc --go_out=../pb --go_opt=paths=source_relative  --go-grpc_out=../pb --go-grpc_opt=paths=source_relative common.proto
protoc --go_out=../pb --go_opt=paths=source_relative  --go-grpc_out=../pb --go-grpc_opt=paths=source_relative demo.proto
protoc --go_out=../pb --go_opt=paths=source_relative  --go-grpc_out=../pb --go-grpc_opt=paths=source_relative user.proto
protoc --go_out=../pb --go_opt=paths=source_relative  --go-grpc_out=../pb --go-grpc_opt=paths=source_relative captcha.proto
protoc --go_out=../pb --go_opt=paths=source_relative  --go-grpc_out=../pb --go-grpc_opt=paths=source_relative token.proto