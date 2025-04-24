defmodule SocketGatewayWeb.UserSocket do
  use Phoenix.Socket

  channel "chat:*", SocketGatewayWeb.ChatChannel

  def connect(_params, socket, _connect_info) do
    {:ok, socket}
  end

  def id(_socket), do: nil
end
