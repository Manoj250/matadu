defmodule SocketGateway.TestRedis do
  alias SocketGateway.RedisClusterCache

  def run do
    key = "croc:test:#{:rand.uniform(1000)}"
    value = %{lang: "elixir", skill: "🐊 giga brained"}

    RedisClusterCache.put(key, value)
    IO.inspect(RedisClusterCache.get(key), label: "GET VALUE BACK")

    RedisClusterCache.delete(key)
    IO.inspect(RedisClusterCache.get(key), label: "AFTER DELETE")
  end
end
