defmodule SocketGateway.Endpoint do
  use GRPC.Endpoint

  intercept GRPC.Server.Interceptors.Logger
  run SocketGateway.SocketService.Server
end
