#!/bin/bash
set -ex
YB_VER="2.25.2.0"
YB_DIR="/opt/yugabyte/yugabyte-${YB_VER}"
rm -rf "$HOME/.yugabyte_data"

echo "cluster tore down"
