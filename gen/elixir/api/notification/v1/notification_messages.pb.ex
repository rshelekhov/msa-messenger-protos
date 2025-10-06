defmodule Api.Notification.V1.RegisterDeviceRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :user_id, 1, type: :string, json_name: "userId"
  field :device_data, 2, type: Api.Notification.V1.UserDeviceData, json_name: "deviceData"
end

defmodule Api.Notification.V1.RegisterDeviceResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :registration, 1, type: Api.Notification.V1.DeviceRegistration
end

defmodule Api.Notification.V1.UnregisterDeviceRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :registration_id, 1, type: :string, json_name: "registrationId"
  field :user_id, 2, type: :string, json_name: "userId"
end

defmodule Api.Notification.V1.UnregisterDeviceResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.15.0", syntax: :proto3

  field :success, 1, type: :bool
end
