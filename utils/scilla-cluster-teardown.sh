#!/bin/bash

echo "💀 Killing Scylla processes like a boss"

pkill -f scylla || true
sleep 2
pkill -9 -f scylla || true

echo "🧹 Removing loopback IPs..."
sudo ip addr del 127.0.0.2/32 dev lo 2>/dev/null || echo "127.0.0.2 IP not found, chill"
sudo ip addr del 127.0.0.3/32 dev lo 2>/dev/null || echo "127.0.0.3 IP not found, chill"

echo "🧽 Cleaning data dirs and logs..."
rm -rf ~/scylla/node{1,2,3}
rm -f scylla_{1,2,3}.log

echo "Cleanup done. No survivors."
