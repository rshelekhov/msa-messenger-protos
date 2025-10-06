defmodule Api.Notification.V1.NotificationService.Service do
  @moduledoc false

  use GRPC.Service,
    name: "api.notification.v1.NotificationService",
    protoc_gen_elixir_version: "0.15.0"

  rpc :RegisterDevice,
      Api.Notification.V1.RegisterDeviceRequest,
      Api.Notification.V1.RegisterDeviceResponse

  rpc :UnregisterDevice,
      Api.Notification.V1.UnregisterDeviceRequest,
      Api.Notification.V1.UnregisterDeviceResponse
end

defmodule Api.Notification.V1.NotificationService.Stub do
  @moduledoc false

  use GRPC.Stub, service: Api.Notification.V1.NotificationService.Service
end
