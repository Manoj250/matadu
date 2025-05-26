defmodule SocketGateway.SocketService.Server do
  use GRPC.Server, service: Socketpb.SocketService.Service

  alias Socketpb.{MessageDetails, CommonResponse}

  @spec poke_user(MessageDetails.t(), GRPC.Server.Stream.t()) :: CommonResponse.t()
  def poke_user(%MessageDetails{message_content: msg}, _stream) do
    IO.inspect(msg, label: "Poked with 🔥")

    %CommonResponse{
      message: "Poked successfully 🫱",
      code: 200,
      success: true
    }
  end
end
