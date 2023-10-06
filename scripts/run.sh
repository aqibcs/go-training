#!/bin/bash

source ./scripts/env.sh


build=false

# Parse command line arguments
while getopts "b" opt; do
  case $opt in
    b)
      Build=true
      ;;
  esac
done

if [ "$Build" = true ]; then
   ./scripts/build.sh
else
    $BINARY $CSV_FILE
fi
