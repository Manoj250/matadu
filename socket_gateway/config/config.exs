# This file is responsible for configuring your application
# and its dependencies with the aid of the Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

config :socket_gateway,
  generators: [timestamp_type: :utc_datetime]

# Configures the endpoint
config :socket_gateway, SocketGatewayWeb.Endpoint,
  url: [host: "localhost"],
  adapter: Bandit.PhoenixAdapter,
  render_errors: [
    formats: [json: SocketGatewayWeb.ErrorJSON],
    layout: false
  ],
  pubsub_server: SocketGateway.PubSub,
  live_view: [signing_salt: "Pd6MAd6H"]

# Configures the mailer
#
# By default it uses the "Local" adapter which stores the emails
# locally. You can see the emails in your browser, at "/dev/mailbox".
#
# For production it's recommended to configure a different adapter
# at the `config/runtime.exs`.
config :socket_gateway, SocketGateway.Mailer, adapter: Swoosh.Adapters.Local

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

config :socket_gateway, SocketGateway.RedisClusterCache,
  mode: :redis_cluster,
  redis_cluster: [
    configuration_endpoints: [
      node1: [host: "localhost", port: 6379],
      node2: [host: "localhost", port: 6380],
      node3: [host: "localhost", port: 6381],
      node4: [host: "localhost", port: 6382],
      node5: [host: "localhost", port: 6383],
      node6: [host: "localhost", port: 6384]
    ]
  ]


# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{config_env()}.exs"
