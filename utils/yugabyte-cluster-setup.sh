#!/bin/bash
set -ex
YB_VER="2.25.2.0"
YB_DIR="/opt/yugabyte"
TAR_NAME="yugabyte-${YB_VER}-b359-linux-x86_64.tar.gz"
URL="https://downloads.yugabyte.com/releases/${YB_VER}/${TAR_NAME}"

sudo mkdir -p "$YB_DIR"
cd /tmp
wget -q "$URL"
tar -xzf "$TAR_NAME"
sudo mv "yugabyte-${YB_VER}" "$YB_DIR"
rm "$TAR_NAME"
cd "$YB_DIR"
./bin/post_install.sh

echo "export PATH=\$PATH:$YB_DIR/bin" >> ~/.zshrc
source ~/.zshrc

echo "YugabyteDB installed at $YB_DIR and added to PATH"
