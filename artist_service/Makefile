artist-api:
	protoc  --proto_path=./api \
            --proto_path=./third_party \
			--go_out=paths=source_relative:./api \
         	--go-grpc_out=paths=source_relative:./api \
           --grpc-gateway_out=paths=source_relative:./api \
           --openapi_out==paths=source_relative:. \
       		 api/artist/v1/artist.proto
