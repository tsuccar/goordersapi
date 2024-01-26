#!/bin/bash -x

# Ensure the required services are started
# sudo service docker start # Required for install-tools script

(
  # cd /app || return 1
  echo "hello world"
  # Gather dependencies and install required tools
#   go mod download
#   scripts/install-tools.sh
)