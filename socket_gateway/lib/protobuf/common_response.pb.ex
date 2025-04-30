defmodule Socketpb.CommonResponse do
  @moduledoc false

  use Protobuf, protoc_gen_elixir_version: "0.14.0", syntax: :proto3

  field :message, 1, type: :string
  field :code, 2, type: :int32
  field :success, 3, type: :bool
end
