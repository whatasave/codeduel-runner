#!/bin/bash

TIMEOUT=$1
INPUT=$2
RUN=$3
MAX_OUTPUT_LENGTH=1000
MAX_ERROR_LENGTH=1000

IFS=$'\n' read -d '' -a inputs <<< "$INPUT"
for input in "${inputs[@]}"; do
    OUTPUT=$(timeout $TIMEOUT $RUN <<< "$input" 2> /tmp/Error)
    if [[ $? -eq 0 ]]; then TERMINATED=true; else TERMINATED=false; fi
    ERROR=$(</tmp/Error)
    jq --null-input --compact-output --arg output "${OUTPUT:0:MAX_OUTPUT_LENGTH}" --arg error "${ERROR:0:MAX_ERROR_LENGTH}" --argjson terminated "$TERMINATED" '$ARGS.named'
    if [[ $TERMINATED == false ]]; then
        break
    fi
done | jq --slurp --compact-output --join-output '.'