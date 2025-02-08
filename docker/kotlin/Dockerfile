FROM zenika/kotlin

WORKDIR /app

RUN apk update && \
    apk add --no-cache jq

COPY base .

CMD [ "sh", "-c", "echo \"$CODE\" > main.kt && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"kotlinc main.kt -include-runtime -d main.jar && java -jar main.jar\"" ]
