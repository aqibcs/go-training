#!/bin/bash

source ./scripts/env.sh

# Check if the output directory exists, create it if not
if [ ! -d "$OutputDirectory" ]; then
    mkdir -p "$OutputDirectory"
fi

go build -o $Binary $SourceFile

echo "CSV to JSON conversion completed successfully. Output file: $Binary"
