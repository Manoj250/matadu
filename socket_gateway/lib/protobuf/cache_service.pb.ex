defmodule Cache.CacheService.Service do
  @moduledoc false

  use GRPC.Service, name: "cache.cache_service", protoc_gen_elixir_version: "0.14.0"

  rpc :RegisterServer, Cache.ServerDetails, Cache.CommonResponse

  rpc :UnRegisterServer, Cache.ServerDetails, Cache.CommonResponse

  rpc :GetServerLoad, Google.Protobuf.Empty, Cache.GetServerLoadResponse

  rpc :SetServerLoad, Cache.SetServerLoadRequest, Cache.CommonResponse

  rpc :RegisterUserToCache, Cache.UserDetails, Cache.CommonResponse

  rpc :UnRegisterUserFromCache, Cache.UserDetails, Cache.CommonResponse

  rpc :GetUserFromCache, Cache.UserDetails, Cache.UserDetails

  rpc :RegisterUserToChat, Cache.UserToChat, Cache.CommonResponse

  rpc :UnRegisterUserFromChat, Cache.UserToChat, Cache.CommonResponse

  rpc :GetUsersByChatId, Cache.ChatDetails, Cache.UsersByChatIdResponse
end

defmodule Cache.CacheService.Stub do
  @moduledoc false

  use GRPC.Stub, service: Cache.CacheService.Service
end
