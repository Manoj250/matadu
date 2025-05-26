#!/bin/bash

# Start the Redis nodes
echo "Starting Redis nodes..."
redis-server --port 6379 --cluster-enabled yes --cluster-config-file nodes_6379.conf --cluster-node-timeout 5000 &
redis-server --port 6380 --cluster-enabled yes --cluster-config-file nodes_6380.conf --cluster-node-timeout 5000 &
redis-server --port 6381 --cluster-enabled yes --cluster-config-file nodes_6381.conf --cluster-node-timeout 5000 &
redis-server --port 6382 --cluster-enabled yes --cluster-config-file nodes_6382.conf --cluster-node-timeout 5000 &
redis-server --port 6383 --cluster-enabled yes --cluster-config-file nodes_6383.conf --cluster-node-timeout 5000 &
redis-server --port 6384 --cluster-enabled yes --cluster-config-file nodes_6384.conf --cluster-node-timeout 5000 &

echo "Waiting 10s for nodes to start..."

sleep 10

expect <<EOF
log_file cluster_create_status.txt
spawn redis-cli --cluster create \
  127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:6381 \
  127.0.0.1:6382 127.0.0.1:6383 127.0.0.1:6384 \
  --cluster-replicas 1
expect "Can I set the above configuration*" {
    send "yes\r"
}
expect eof
EOF

echo "nodetool -h 127.0.0.1 status -> for helath check"