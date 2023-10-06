#!/bin/bash

source ./scripts/env.sh

# Check if the output directory exists, create it if not
if [ ! -d "$OUTPUT_DIRECTORY" ]; then
    mkdir -p "$OUTPUT_DIRECTORY"
fi

go build -o $BINARY $SOURCE_FILE

echo "Build project successfully. Binary file: $Binary"
