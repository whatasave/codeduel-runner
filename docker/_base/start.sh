#!/bin/sh

TIMEOUT=$1
INPUT=$2
RUN=$3
MAX_OUTPUT_LENGTH=1000
MAX_ERROR_LENGTH=1000

printf '%s' "$INPUT" | jq -c '.[]' | while IFS= read -r item; do
    result=$(printf '%s' "$item" | jq -rc | timeout "$TIMEOUT" $RUN 2> /tmp/Error)
    exit_status=$?
    err=$(cat /tmp/Error)
    
    output=$(printf "%.*s" "$MAX_OUTPUT_LENGTH" "$result")
    error=$(printf "%.*s" "$MAX_ERROR_LENGTH" "$err")
    
    jq --null-input --compact-output \
       --arg output "$output" \
       --arg errors "$error" \
       --argjson status "$exit_status" \
       '$ARGS.named'
    
    if [ "$TERMINATED" = "false" ]; then
        break
    fi
done | jq --slurp --compact-output --join-output '.'