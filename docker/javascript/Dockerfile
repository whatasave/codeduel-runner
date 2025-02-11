FROM node:23-alpine3.20

WORKDIR /app

COPY base .
COPY runner /runner

RUN apk add --no-cache jq &&\
    npm init -y && \
    npm prune --production && \
    chmod +x start.sh

ENTRYPOINT [ "sh", "-c", "echo \"$CODE\" > main.js && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"node -r /runner/polyfill.js main.js\"" ]