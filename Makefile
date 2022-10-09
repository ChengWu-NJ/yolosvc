PBDIR="./pkg/pb"

.PHONY: all
all:
	@$(MAKE) --no-print-directory deps
	@$(MAKE) --no-print-directory protobuf
#	@$(MAKE) --no-print-directory app

.PHONY: deps
deps:
	go mod tidy

.PHONY: protobuf
protobuf:
	### 1. grpc first
	protoc -I ${PBDIR} -I /usr/local/include \
	--go_out ${PBDIR} \
	--go_opt paths=source_relative \
	--go-grpc_out ${PBDIR} \
	--go-grpc_opt paths=source_relative \
	${PBDIR}/*.proto

	### 2. gateway of restful
	protoc -I ${PBDIR} -I /usr/local/include \
	--grpc-gateway_out ${PBDIR} \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ${PBDIR} \
	--openapiv2_opt logtostderr=true \
	${PBDIR}/*.proto

#.PHONY: app
#app:
#	go build

