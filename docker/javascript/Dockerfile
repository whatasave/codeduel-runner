FROM node:20.10.0-slim

WORKDIR /app

RUN npm init -y
RUN npm install readline-sync
RUN npm prune --production

COPY polyfill.js polyfill.js
COPY readline.js readline.js
CMD echo $CODE > main.js && echo -n $INPUT | timeout $TIMEOUT node -r ./polyfill.js main.js