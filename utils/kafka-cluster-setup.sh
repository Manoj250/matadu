#!/bin/bash

# ==== CONFIG ====
KAFKA_VERSION="4.0.0"
SCALA_VERSION="2.13"
TGZ_FILE="kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz"
TGZ_URL="https://downloads.apache.org/kafka/${KAFKA_VERSION}/${TGZ_FILE}"

# ==== DESTINATION ====
INSTALL_DIR="/mnt/d"
KAFKA_DIR="$INSTALL_DIR/kafka_${SCALA_VERSION}-${KAFKA_VERSION}"

mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR" || exit 1

# ==== CHECK EXISTING ====
if [ -d "$KAFKA_DIR" ]; then
  echo "Kafka already installed at $KAFKA_DIR"
  exit 0
fi

# ==== DOWNLOAD & UNZIP ====
curl -fLo "$TGZ_FILE" "$TGZ_URL"
tar -xzf "$TGZ_FILE"
rm -f "$TGZ_FILE"

# ==== JACKSON PATCH ====
cd "$KAFKA_DIR/libs" || exit 1
curl -O https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-core/2.15.2/jackson-core-2.15.2.jar
curl -O https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-annotations/2.15.2/jackson-annotations-2.15.2.jar
curl -O https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/2.15.2/jackson-databind-2.15.2.jar

echo "Kafka setup complete at $KAFKA_DIR"
