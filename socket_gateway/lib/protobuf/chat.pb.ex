defmodule Cache.ChatDetails do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :chat_id, 1, type: :string, json_name: "chatId"
  field :chat_name, 2, type: :string, json_name: "chatName"
  field :chat_type, 3, type: :string, json_name: "chatType"
end

defmodule Cache.UsersByChatIdResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :user_list, 1, repeated: true, type: Cache.UserDetails, json_name: "userList"
  field :status_response, 2, type: Cache.CommonResponse, json_name: "statusResponse"
end
