defmodule Api.Chat.V1.Chat do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :id, 1, type: :string
  field :user_id, 2, type: :string, json_name: "userId"
  field :friend_id, 3, type: :string, json_name: "friendId"
  field :created_at, 4, type: Google.Protobuf.Timestamp, json_name: "createdAt"
end

defmodule Api.Chat.V1.Message do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :id, 1, type: :string
  field :sender_id, 2, type: :string, json_name: "senderId"
  field :content, 3, type: :string
  field :created_at, 4, type: Google.Protobuf.Timestamp, json_name: "createdAt"
end
