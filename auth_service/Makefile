
auth-api:
	protoc  --proto_path=./api \
            --proto_path=./third_party \
			--go_out=paths=source_relative:./api \
			--validate_out=paths=source_relative,lang=go:./api \
         	--go-grpc_out=paths=source_relative:./api \
           --grpc-gateway_out=paths=source_relative:./api \
           --openapi_out==paths=source_relative:. \
       		 api/auth/v1/auth.proto

