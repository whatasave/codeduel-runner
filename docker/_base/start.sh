#!/bin/bash

TIMEOUT=$1
INPUT=$2
RUN=$3
MAX_OUTPUT_LENGTH=1000
MAX_ERROR_LENGTH=1000

echo -n "$INPUT" | jq -c '.[]' | while read -r item; do
    result=$(jq -rc <<< "$item" | timeout $TIMEOUT $RUN 2> /tmp/Error)
    exit_status=$?
    err=$(</tmp/Error)
    jq --null-input --compact-output --arg output "${result:0:MAX_OUTPUT_LENGTH}" --arg error "${err:0:MAX_ERROR_LENGTH}" --argjson status "$exit_status" '$ARGS.named'
    if [[ $TERMINATED == false ]]; then
        break
    fi
done | jq --slurp --compact-output --join-output '.'