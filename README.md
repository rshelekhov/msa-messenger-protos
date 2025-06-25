# MSA Messenger Protos

This repository contains Protocol Buffers definitions for the Messenger microservices architecture.

## Services Overview

### Chat Service

Handles one-to-one chat functionality:

- Chat creation and management
- Message sending and receiving
- Chat history and message storage
- Real-time message delivery

### Subscriber Service

Manages friend relationships between users:

- Friend request management (send/accept/decline)
- Friends list management
- Friend relationship status tracking

### Notification Service

Handles push notifications and device management:

- Device registration for notifications
- Push notification delivery
- Multi-device notification support
- Integration with external notification services

## Directory Structure

```
api/
├── chat/v1/         # Chat service protos
├── notification/v1/ # Notification service protos
├── subscriber/v1/   # Friend management service protos
```

## Code Generation

This repository uses [buf](https://buf.build) for Protocol Buffers code generation.

To generate code:

```bash
buf generate .
```

## Dependencies

- Go 1.24+
- Protocol Buffers
- buf (for code generation)
- gRPC
