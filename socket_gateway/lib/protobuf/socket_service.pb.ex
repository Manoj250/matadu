defmodule Cache.SocketService.Service do
  @moduledoc false

  use GRPC.Service, name: "cache.SocketService", protoc_gen_elixir_version: "0.14.0"

  rpc :PokeUser, Cache.MessageDetails, Cache.CommonResponse
end

defmodule Cache.SocketService.Stub do
  @moduledoc false

  use GRPC.Stub, service: Cache.SocketService.Service
end
