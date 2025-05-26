defmodule SocketGatewayWeb.ChatChannel do
  use Phoenix.Channel

  def join("chat:lobby", _params, socket) do
    {:ok, socket}
  end

  def handle_in(_event, _payload, socket) do
    push(socket, "echo", %{"msg" => "world"})
    {:noreply, socket}
  end
end
