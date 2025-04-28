defmodule Cache.MessageDetails do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :message_content, 1, type: :string, json_name: "messageContent"
  field :to_user_id, 2, type: :string, json_name: "toUserId"
  field :from_user_id, 3, type: :string, json_name: "fromUserId"
  field :sent_at, 4, type: :string, json_name: "sentAt"
  field :chat_id, 5, type: :string, json_name: "chatId"
end
