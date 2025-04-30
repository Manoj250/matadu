.PHONY: protogen setupGrpc

ELIXIR_PLUGIN := $(shell asdf which protoc-gen-elixir)
GOPATH_BIN := $(shell realpath $$(go env GOPATH)/bin)

protogen:
	@if [ -z "$(PROTO_DIR)" ]; then \
		echo "❌ Usage: make protogen PROTO_DIR=./proto"; \
		exit 1; \
	fi
	@if [ ! -x "$(ELIXIR_PLUGIN)" ]; then \
		echo "🤡 protoc-gen-elixir not found bro! Run: make setupGrpc"; \
		exit 1; \
	fi
	@echo "📁 Using PROTO_DIR=$(PROTO_DIR)"
	@echo "📁 Generating Go code to ./../cache/internal/pb"
	@echo "📁 Generating Elixir code to ./../socket_gateway/lib/protobuf"

	protoc \
	  -I$(shell realpath $(PROTO_DIR)) \
	  --plugin=protoc-gen-go=$(GOPATH_BIN)/protoc-gen-go \
	  --plugin=protoc-gen-go-grpc=$(GOPATH_BIN)/protoc-gen-go-grpc \
	  --plugin=protoc-gen-elixir=$(ELIXIR_PLUGIN) \
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
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
	mix do local.hex --force, local.rebar --force, escript.install hex protobuf 0.14.0 --force
	asdf reshim elixir
	@echo "🎉 protoc and plugins installed like a legend"
