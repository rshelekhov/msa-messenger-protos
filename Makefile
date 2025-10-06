# Use bin in the current directory to install protoc plugins
LOCAL_BIN := $(CURDIR)/bin

# Add bin in the current directory to the PATH when running protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Path to protobuf files
PROTO_PATH := $(CURDIR)/api

# Path to generated .pb.go files
PKG_PROTO_PATH := $(CURDIR)/gen/go

# Path to vendored protobuf files
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

# Install necessary plugins
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	# Go plugins
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

	# Elixir plugin
	@if command -v mix >/dev/null 2>&1; then \
		echo "Installing Elixir protobuf plugin..."; \
		mix archive.install hex protobuf --force; \
	else \
		echo "Warning: Elixir/Mix not found. Skipping Elixir plugin installation."; \
		echo "Please install Elixir and run 'make install-elixir-deps' manually."; \
	fi

generate: export GOBIN := $(LOCAL_BIN)
generate:
	$(MAKE) generate-go
	$(MAKE) generate-elixir

# Generate Go code using protoc
generate-go:
	@echo "Generating Go code..."
	@mkdir -p gen/go
	@echo "Generating Chat and Notification services..."
	$(PROTOC) --go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/api/chat/v1/*.proto \
		proto/api/notification/v1/*.proto
	@echo "Generating Subscriber service with validation..."
	$(PROTOC) --go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		--validate_out=lang=go,paths=source_relative:gen/go \
		--proto_path=proto \
		--proto_path=vendor \
		proto/api/subscriber/v1/*.proto
	@echo "✓ Generated complete Go code for all services"

# Generate Elixir code using protoc (full gRPC services for chat and notification)
generate-elixir:
	@echo "Generating Elixir code..."
	@mkdir -p gen/elixir
	@echo "Generating Chat and Notification services..."
	PATH="$$PATH:$$HOME/.mix/escripts" protoc --elixir_out=plugins=grpc:gen/elixir \
		--proto_path=proto \
		proto/api/chat/v1/*.proto \
		proto/api/notification/v1/*.proto
	@echo "✓ Generated complete Elixir gRPC services for chat and notification"

# Linter
lint: .proto-lint

# Linter proto files
.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

# Format protobuf files
proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w
	
# Install Elixir dependencies separately
install-elixir-deps:
	$(info Installing Elixir protobuf dependencies...)
	mix archive.install hex protobuf --force
	mix escript.install hex protobuf --force

# Declare that the current commands are not files and
# instrument Makefile to not search for changes in the file system
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	.proto-lint \
	proto-format \
	vendor \
	generate \
	build \
	lint \
	install-elixir-deps