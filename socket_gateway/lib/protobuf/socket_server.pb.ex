defmodule Cache.ServerDetails do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :server_address, 1, type: :string, json_name: "serverAddress"
  field :server_id, 2, type: :string, json_name: "serverId"
end

defmodule Cache.ServerLoad do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :server_address, 1, type: :string, json_name: "serverAddress"
  field :load, 2, type: :int32
end

defmodule Cache.GetServerListResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :server_list, 1, repeated: true, type: Cache.ServerDetails, json_name: "serverList"
  field :status_response, 2, type: Cache.CommonResponse, json_name: "statusResponse"
end

defmodule Cache.GetServerLoadResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :server_load, 1, repeated: true, type: Cache.ServerLoad, json_name: "serverLoad"
  field :status_response, 2, type: Cache.CommonResponse, json_name: "statusResponse"
end

defmodule Cache.SetServerLoadRequest do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :server_load, 1, repeated: true, type: Cache.ServerLoad, json_name: "serverLoad"
end
