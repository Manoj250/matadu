#!/bin/bash
set -ex
YB_VER="2.25.2.0"
YB_DIR="/opt/yugabyte/yugabyte-${YB_VER}"

$YB_DIR/bin/yugabyted start \
  --base_dir "$HOME/.yugabyte_data" \
  --advertise_address 127.0.0.1

echo "YugabyteDB running at http://127.0.0.1:15433"
