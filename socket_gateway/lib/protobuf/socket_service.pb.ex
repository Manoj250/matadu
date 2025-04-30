defmodule Socketpb.SocketService.Service do
  @moduledoc false

  use GRPC.Service, name: "socketpb.SocketService", protoc_gen_elixir_version: "0.14.0"

  rpc :PokeUser, Socketpb.MessageDetails, Socketpb.CommonResponse
end

defmodule Socketpb.SocketService.Stub do
  @moduledoc false

  use GRPC.Stub, service: Socketpb.SocketService.Service
end
