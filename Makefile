OTEL_DOCKER_PROTOBUF ?= otel/build-protobuf:0.2.1

PROTOC := docker run --rm -u ${shell id -u} -v${PWD}:${PWD} -w${PWD}/$(OTLP_PROTO_INTERMEDIATE_DIR) ${OTEL_DOCKER_PROTOBUF} --proto_path=${PWD}/$(OTLP_PROTO_INTERMEDIATE_DIR)
PROTO_INCLUDES := -I/usr/include/github.com/gogo/protobuf

.PHONY: proto
proto:
	$(PROTOC) $(PROTO_INCLUDES) --gogo_out=plugins=grpc,:./ proto/signalfx_metrics.proto
	mv proto/signalfx_metrics.pb.go model/
