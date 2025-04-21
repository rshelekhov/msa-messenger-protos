# MSA Messenger Protos

This repository contains Protocol Buffers definitions for the Messenger microservices architecture.

## Services Overview

### Auth Service

Handles user authentication and session management:

- User registration and login
- Token management (access/refresh)
- Device registration for multi-device support
- Email updates and session management

### User Service

Manages user profiles and personal information:

- User profile storage and updates
- Profile information synchronization with other services
- User data management

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

## Common Types

The `common` package contains shared data types used across multiple services:

- Device information structures
- Platform enums
- Shared data types

## Directory Structure

```
api/
├── auth/v1/         # Authentication service protos
├── chat/v1/         # Chat service protos
├── common/v1/       # Shared types and enums
├── notification/v1/ # Notification service protos
├── subscriber/v1/   # Friend management service protos
└── user/v1/        # User profile service protos
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
