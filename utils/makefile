.PHONY: protogen setupGrpc

protogen:
	@if [ -z "$(PROTO_DIR)" ] || [ -z "$(OUT_DIR)" ]; then \
		echo "❌ Usage: make protogen PROTO_DIR=./proto OUT_DIR=./internal/pb"; \
		exit 1; \
	fi
	@echo "📁 Using PROTO_DIR=$(PROTO_DIR)"
	@echo "📁 Using OUT_DIR=$(OUT_DIR)"
	protoc \
	  -I$(shell realpath $(PROTO_DIR)) \
	  --plugin=protoc-gen-go=$(shell realpath $$(go env GOPATH)/bin/protoc-gen-go) \
	  --plugin=protoc-gen-go-grpc=$(shell realpath $$(go env GOPATH)/bin/protoc-gen-go-grpc) \
	  --go_out=$(OUT_DIR) \
	  --go-grpc_out=$(OUT_DIR) \
	  $(shell realpath $(PROTO_DIR))/*.proto
	@echo "✅ Code generated in $(OUT_DIR)"

setupGrpc:
	@echo "📦 Installing protoc & Go plugins..."
	wget -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v27.1/protoc-27.1-linux-x86_64.zip && \
	sudo apt install -y unzip && \
	unzip -o protoc.zip -d /tmp/protoc && \
	sudo rm -rf /usr/local/include/google && \
	sudo mv -f /tmp/protoc/bin/protoc /usr/local/bin/ && \
	sudo mv -f /tmp/protoc/include/* /usr/local/include/ && \
	rm protoc.zip && \
	rm -rf /tmp/protoc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "🎉 All set, legend"
