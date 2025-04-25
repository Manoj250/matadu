.PHONY: protogen setupGrpc

protogen:
	@if [ -z "$(PROTO_DIR)" ]; then \
		echo "❌ Usage: make protogen PROTO_DIR=./proto"; \
		exit 1; \
	fi
	@echo "📁 Using PROTO_DIR=$(PROTO_DIR)"
	@echo "📁 Generating Go code to ./../cache/internal/pb"
	@echo "📁 Generating Elixir code to ./../socket_gateway/lib/protobuf"

	protoc \
	  -I$(shell realpath $(PROTO_DIR)) \
	  --plugin=protoc-gen-go=$(shell realpath $$(go env GOPATH)/bin/protoc-gen-go) \
	  --plugin=protoc-gen-elixir=$(shell realpath /home/manoj/.asdf/installs/elixir/1.18.3-otp-27/.mix/escripts/protoc-gen-elixir) \
	  --plugin=protoc-gen-go-grpc=$(shell realpath $$(go env GOPATH)/bin/protoc-gen-go-grpc) \
	  --go_out=./../cache/internal/pb \
	  --go-grpc_out=./../cache/internal/pb \
	  --elixir_out=./../socket_gateway/lib/protobuf \
	  --elixir_opt=plugins=grpc \
	  $(shell realpath $(PROTO_DIR))/*.proto

	@echo "✅ protobuf generated"


setupGrpc:
	@echo "📦 Installing protoc & Go plugins..."
	wget -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v30.2/protoc-30.2-linux-x86_64.zip && \
	sudo apt install -y unzip && \
	unzip -o protoc.zip -d /tmp/protoc && \
	sudo rm -rf /usr/local/include/google && \
	sudo mv -f /tmp/protoc/bin/protoc /usr/local/bin/ && \
	sudo mv -f /tmp/protoc/include/* /usr/local/include/ && \
	rm protoc.zip && \
	rm -rf /tmp/protoc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	mix escript.install hex protobuf 0.14.0
	@echo "🎉 protoc installed"
