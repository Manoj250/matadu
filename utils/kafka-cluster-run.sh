#!/bin/bash

# ==== ENV VARS DEFAULTS ====
export BROKER_ID="${BROKER_ID:-1}"
export KAFKA_PORT="${KAFKA_PORT:-9092}"
export INSTALL_DIR="/mnt/d"
export KAFKA_VERSION="4.0.0"
export SCALA_VERSION="2.13"
export KAFKA_DIR="$INSTALL_DIR/kafka_${SCALA_VERSION}-${KAFKA_VERSION}"
export CLUSTER_ID="${CLUSTER_ID:-$("$KAFKA_DIR/bin/kafka-storage.sh" random-uuid)}"

cd "$KAFKA_DIR" || { echo "Kafka dir not found. Run setup first"; exit 1; }

# ==== CONFIGURE ====
mkdir -p kraft-logs

cat > kraft-server.properties <<EOF
process.roles=broker,controller
node.id=$BROKER_ID
controller.quorum.voters=$BROKER_ID@localhost:9093
listeners=PLAINTEXT://localhost:$KAFKA_PORT,CONTROLLER://localhost:9093
inter.broker.listener.name=PLAINTEXT
log.dirs=kraft-logs
num.network.threads=3
num.io.threads=8
log.retention.hours=168
log.segment.bytes=1073741824
log.retention.check.interval.ms=300000
offsets.topic.replication.factor=1
transaction.state.log.replication.factor=1
transaction.state.log.min.isr=1
listener.security.protocol.map=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT
controller.listener.names=CONTROLLER
EOF

# ==== FORMAT STORAGE IF FIRST TIME ====
if [ ! -f kraft-logs/meta.properties ]; then
  echo "Formatting logs with cluster ID: $CLUSTER_ID"
  bin/kafka-storage.sh format -t "$CLUSTER_ID" -c kraft-server.properties
fi

# ==== RUN ====
echo "Launching Kafka on localhost:$KAFKA_PORT"
bin/kafka-server-start.sh kraft-server.properties
