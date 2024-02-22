FROM node:20.10.0-slim

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y\
 && npm init -y\
 && npm install readline-sync\
 && npm prune --production

COPY base .
COPY runner /runner

CMD echo "$CODE" > main.js && ./start.sh "$TIMEOUT" "$INPUT" "node -r /runner/polyfill.js main.js"