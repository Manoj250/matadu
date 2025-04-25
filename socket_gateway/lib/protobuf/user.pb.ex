defmodule Cache.UserDetails do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :user_id, 1, type: :string, json_name: "userId"
  field :user_name, 2, type: :string, json_name: "userName"
  field :user_email, 3, type: :string, json_name: "userEmail"
  field :user_phone, 4, type: :string, json_name: "userPhone"
  field :user_status, 5, type: :string, json_name: "userStatus"
  field :device_id, 6, type: :string, json_name: "deviceId"
end
