#!/bin/bash

TIMEOUT=$1
INPUT=$2
RUN=$3
MAX_OUTPUT_LENGTH=1000
MAX_ERROR_LENGTH=1000

echo -n "$INPUT" | jq -c '.[]' | while read -r input; do
    OUTPUT=$(jq -rc <<< "$input" | timeout $TIMEOUT $RUN 2> /tmp/Error)
    STATUS=$?
    ERROR=$(</tmp/Error)
    jq --null-input --compact-output --arg output "${OUTPUT:0:MAX_OUTPUT_LENGTH}" --arg error "${ERROR:0:MAX_ERROR_LENGTH}" --argjson status "$STATUS" '$ARGS.named'
    if [[ $TERMINATED == false ]]; then
        break
    fi
done | jq --slurp --compact-output --join-output '.'