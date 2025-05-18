#!/bin/bash

echo "Slaying Redis demons... 🪓"

# Kill all Redis servers
pkill -f "redis-server --port 6379"
pkill -f "redis-server --port 6380"
pkill -f "redis-server --port 6381"
pkill -f "redis-server --port 6382"
pkill -f "redis-server --port 6383"
pkill -f "redis-server --port 6384"

# Delete all cluster config files and dump files
rm -f nodes_6379.conf appendonly.aof dump.rdb
rm -f nodes_6380.conf
rm -f nodes_6381.conf
rm -f nodes_6382.conf
rm -f nodes_6383.conf
rm -f nodes_6384.conf
rm -f cluster_create_status.txt

sudo lsof -ti :6379-6384 | xargs sudo kill -9

echo "🔥 Redis cluster dismantled. Back to ashes."
