defmodule Api.Notification.V1.Platform do
  @moduledoc false

  use Protobuf, enum: true, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :PLATFORM_UNSPECIFIED, 0
  field :PLATFORM_WEB, 1
  field :PLATFORM_IOS, 2
  field :PLATFORM_ANDROID, 3
end

defmodule Api.Notification.V1.DeviceRegistration do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :id, 1, type: :string
  field :user_id, 2, type: :string, json_name: "userId"
  field :device_data, 3, type: Api.Notification.V1.UserDeviceData, json_name: "deviceData"
  field :registered_at, 4, type: Google.Protobuf.Timestamp, json_name: "registeredAt"
end

defmodule Api.Notification.V1.UserDeviceData do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  oneof :version, 0

  field :device_id, 1, type: :string, json_name: "deviceId"
  field :platform, 2, type: Api.Notification.V1.Platform, enum: true
  field :app_version, 3, type: :string, json_name: "appVersion", oneof: 0
  field :browser_version, 4, type: :string, json_name: "browserVersion", oneof: 0
  field :ip_address, 5, type: :string, json_name: "ipAddress"
  field :user_agent, 6, type: :string, json_name: "userAgent"
end
