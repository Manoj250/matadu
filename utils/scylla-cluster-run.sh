#!/bin/bash

echo "Launching 3-node ScyllaDB cluster"

for i in 1 2 3; do
  ip="127.0.0.$i"
  node_dir=/root/scylla/node$i

  echo "Preparing node $i at $ip"
  mkdir -p $node_dir/{data,commitlog,hints,logs}
  
  cat > $node_dir/scylla.yaml <<EOF
cluster_name: 'chatgod_cluster'
listen_address: $ip
rpc_address: $ip
seed_provider:
  - class_name: org.apache.cassandra.locator.SimpleSeedProvider
    parameters:
      - seeds: "127.0.0.1"
data_file_directories:
  - $node_dir/data
commitlog_directory: $node_dir/commitlog
hints_directory: $node_dir/hints
view_hints_directory: $node_dir/view_hints
EOF

  echo "Starting node $i..."
  SCYLLA_CONF=$node_dir scylla --developer-mode 1 \
    --log-to-stdout 1 > scylla_$i.log 2>&1 &
done

echo "All nodes launched"