#!/bin/bash

# Start the Redis nodes
echo "Starting Redis nodes..."
redis-server --port 6379 --cluster-enabled yes --cluster-config-file nodes_6379.conf --cluster-node-timeout 5000 &
redis-server --port 6380 --cluster-enabled yes --cluster-config-file nodes_6380.conf --cluster-node-timeout 5000 &
redis-server --port 6381 --cluster-enabled yes --cluster-config-file nodes_6381.conf --cluster-node-timeout 5000 &
redis-server --port 6382 --cluster-enabled yes --cluster-config-file nodes_6382.conf --cluster-node-timeout 5000 &
redis-server --port 6383 --cluster-enabled yes --cluster-config-file nodes_6383.conf --cluster-node-timeout 5000 &
redis-server --port 6384 --cluster-enabled yes --cluster-config-file nodes_6384.conf --cluster-node-timeout 5000 &

