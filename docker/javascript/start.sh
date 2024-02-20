#!/bin/bash

CODE=$1
TIMEOUT=$2
INPUT=$3

echo $CODE > main.js
IFS='|' read -a inputs <<< "$INPUT"
for input in "${inputs[@]}"; do
    OUTPUT=$(timeout $TIMEOUT node -r /runner/polyfill.js main.js <<< "$input" 2> /tmp/Error)
    ERROR=$(</tmp/Error)
    if [[ $? -eq 0 ]]; then TERMINATED=true; else TERMINATED=false; fi
    jq --null-input --compact-output --arg output "$OUTPUT" --arg error "$ERROR" --argjson terminated "$TERMINATED" '$ARGS.named'
done | jq --slurp --compact-output --join-output '.'