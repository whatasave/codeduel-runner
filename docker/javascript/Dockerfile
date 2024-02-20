FROM node:20.10.0-slim

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y\
 && npm init -y\
 && npm install readline-sync\
 && npm prune --production

COPY runner /runner
COPY start.sh start.sh

CMD ./start.sh "$CODE" "$TIMEOUT" "$INPUT"