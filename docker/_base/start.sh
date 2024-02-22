#!/bin/bash

TIMEOUT=$1
INPUT=$2
RUN=$3
MAX_OUTPUT_LENGTH=1000
MAX_ERROR_LENGTH=1000

SKIP_OTHERS=false
IFS=$'\n' read -d '' -a inputs <<< "$INPUT"
for input in "${inputs[@]}"; do
    if [[ $SKIP_OTHERS == false ]]; then
        OUTPUT=$(timeout $TIMEOUT $RUN <<< "$input" 2> /tmp/Error)
        ERROR=$(</tmp/Error)
        if [[ $? -eq 0 ]]; then
            TERMINATED=true
            SKIP_OTHERS=true
        else
            TERMINATED=false
        fi
        jq --null-input --compact-output --arg output "${OUTPUT:0:MAX_OUTPUT_LENGTH}" --arg error "${ERROR:0:MAX_ERROR_LENGTH}" --argjson terminated "$TERMINATED" '$ARGS.named'
    else
        jq --null-input --compact-output --argjson skipped "true" '$ARGS.named'
    fi
done | jq --slurp --compact-output --join-output '.'