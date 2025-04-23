defmodule SocketGateway.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      SocketGatewayWeb.Telemetry,
      {DNSCluster, query: Application.get_env(:socket_gateway, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: SocketGateway.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: SocketGateway.Finch},
      # Start a worker by calling: SocketGateway.Worker.start_link(arg)
      # {SocketGateway.Worker, arg},
      # Start to serve requests, typically the last entry
      SocketGatewayWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: SocketGateway.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    SocketGatewayWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
