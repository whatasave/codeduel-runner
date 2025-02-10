FROM node:20.10.0-slim

WORKDIR /app

COPY base/ ./base
COPY runner/ ./runner

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    npm init -y && \
    npm prune --production && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh

ENTRYPOINT [ "sh", "-c", "echo \"$CODE\" > main.js && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"node -r /runner/polyfill.js main.js\"" ]