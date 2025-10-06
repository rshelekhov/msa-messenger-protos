# MSA Messenger Protos

This repository contains Protocol Buffers definitions for the Messenger microservices architecture.

## Services Overview

### Chat Service (Elixir Implementation)

Handles one-to-one chat functionality:

- Chat creation and management
- Message sending and receiving
- Chat history and message storage
- Real-time message delivery

### Notification Service (Elixir Implementation)

Handles push notifications and device management:

- Device registration for notifications
- Push notification delivery
- Multi-device notification support
- Integration with external notification services

### Subscriber Service (Go Implementation)

Manages friend relationships between users:

- Friend request management (send/accept/decline)
- Friends list management
- Friend relationship status tracking
- Input validation with protoc-gen-validate

## Architecture

This is a polyglot microservices setup:

- **Chat Service**: Elixir (uses generated Elixir server stubs)
- **Notification Service**: Elixir (uses generated Elixir server stubs)
- **Subscriber Service**: Go (uses generated Go server stubs with validation)

## Directory Structure

```
proto/
├── api/
│   ├── chat/v1/         # Chat service protos
│   ├── notification/v1/ # Notification service protos
│   └── subscriber/v1/   # Subscriber service protos (with validation)
vendor/
└── validate/            # Vendored protoc-gen-validate proto files
gen/
├── go/                  # Generated Go code (all services)
└── elixir/              # Generated Elixir code (chat & notification only)
```

## Code Generation

This repository uses `protoc` directly for Protocol Buffers code generation.

### Prerequisites

Install required tools:

```bash
make .bin-deps
```

### Generate All Code

```bash
make generate
```

### Generate Go Code Only

```bash
make generate-go
```

### Generate Elixir Code Only

```bash
make generate-elixir
```

## Validation

- **Subscriber Service**: Uses protoc-gen-validate for request validation
- **Chat/Notification Services**: No validation (Elixir compatible)

## Dependencies

- Go 1.21+
- Protocol Buffers (protoc)
- protoc-gen-go
- protoc-gen-go-grpc
- protoc-gen-validate
- Elixir + Mix (for Elixir code generation)
- protobuf Elixir package
