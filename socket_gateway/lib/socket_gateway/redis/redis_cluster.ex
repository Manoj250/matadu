defmodule SocketGateway.RedisClusterCache do
  use Nebulex.Cache,
    otp_app: :socket_gateway,
    adapter: NebulexRedisAdapter
end
