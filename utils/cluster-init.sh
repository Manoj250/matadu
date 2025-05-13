#!/bin/sh

# Wait for Redis to start and others to become reachable
for host in redis-1 redis-2 redis-3 redis-4 redis-5 redis-6; do
  until redis-cli -h $host -p 6379 ping | grep -q PONG; do
    echo "Waiting for $host to be ready..."
    sleep 1
  done
done

echo "All Redis nodes are up! Creating cluster..."

redis-cli --cluster create \
  redis-1:6379 redis-2:6379 redis-3:6379 \
  redis-4:6379 redis-5:6379 redis-6:6379 \
  --cluster-replicas 1 --cluster-yes
