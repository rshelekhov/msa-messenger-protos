defmodule Api.Chat.V1.ChatService.Service do
  @moduledoc false

  use GRPC.Service, name: "api.chat.v1.ChatService", protoc_gen_elixir_version: "0.15.0"

  rpc :CreateChat, Api.Chat.V1.CreateChatRequest, Api.Chat.V1.CreateChatResponse

  rpc :GetChat, Api.Chat.V1.GetChatRequest, Api.Chat.V1.GetChatResponse

  rpc :DeleteChat, Api.Chat.V1.DeleteChatRequest, Google.Protobuf.Empty

  rpc :ListChats, Api.Chat.V1.ListChatsRequest, Api.Chat.V1.ListChatsResponse

  rpc :SendMessage, Api.Chat.V1.SendMessageRequest, Api.Chat.V1.SendMessageResponse
end

defmodule Api.Chat.V1.ChatService.Stub do
  @moduledoc false

  use GRPC.Stub, service: Api.Chat.V1.ChatService.Service
end
