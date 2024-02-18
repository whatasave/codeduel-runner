FROM node:20.10.0-slim

WORKDIR /app

RUN stty -echo
RUN npm init -y
RUN npm install readline-sync
RUN npm prune --production

COPY library.js library.js
CMD echo 'const { readline } = require("./library.js");\n'$CODE > main.js && echo -n $INPUT | timeout $TIMEOUT node main.js