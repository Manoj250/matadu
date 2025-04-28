defmodule SocketGateway.GRPC.SocketGRPCServer do
  use GRPC.Service, service: SocketGateway.Protobuf.SocketService.Service

  alias SocketGateway.Protobuf.{MessageDetails, CommonResponse}

	@spec poke_user(MessageDetails.t(), GRPC.Server.Stream.t()) :: CommonResponse.t()
 def poke_user(%MessageDetails{message_content: message, to_user_id: to_id, from_user_id: from_id, sent_at: sent_at, chat_id: chat_id}, _stream) do
    IO.inspect(request, label: "Received poke request 🔥")

    %CommonResponse{
      message: "Poked successfully 🫱",
      code: 200,
      success: true
    }
  end

end
