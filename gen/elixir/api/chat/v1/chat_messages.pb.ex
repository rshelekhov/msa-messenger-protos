defmodule Api.Chat.V1.CreateChatRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :user_id, 1, type: :string, json_name: "userId"
  field :friend_id, 2, type: :string, json_name: "friendId"
end

defmodule Api.Chat.V1.CreateChatResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chat, 1, type: Api.Chat.V1.Chat
end

defmodule Api.Chat.V1.GetChatRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chat_id, 1, type: :string, json_name: "chatId"
  field :user_id, 2, type: :string, json_name: "userId"
  field :page_token, 3, type: :string, json_name: "pageToken"
end

defmodule Api.Chat.V1.GetChatResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chat, 1, type: Api.Chat.V1.Chat
  field :messages, 2, repeated: true, type: Api.Chat.V1.Message
  field :next_page_token, 3, type: :string, json_name: "nextPageToken"
end

defmodule Api.Chat.V1.DeleteChatRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chat_id, 1, type: :string, json_name: "chatId"
  field :user_id, 2, type: :string, json_name: "userId"
end

defmodule Api.Chat.V1.DeleteChatResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :success, 1, type: :bool
end

defmodule Api.Chat.V1.ListChatsRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :user_id, 1, type: :string, json_name: "userId"
  field :page_token, 2, type: :string, json_name: "pageToken"
end

defmodule Api.Chat.V1.ListChatsResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chats, 1, repeated: true, type: Api.Chat.V1.Chat
  field :next_page_token, 2, type: :string, json_name: "nextPageToken"
end

defmodule Api.Chat.V1.SendMessageRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :chat_id, 1, type: :string, json_name: "chatId"
  field :user_id, 2, type: :string, json_name: "userId"
  field :content, 3, type: :string
end

defmodule Api.Chat.V1.SendMessageResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :message, 1, type: Api.Chat.V1.Message
end
