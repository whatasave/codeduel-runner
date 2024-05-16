FROM saschpe/kotlinc

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y

COPY base .

CMD echo "$CODE" > main.kt && ./start.sh "$TIMEOUT" "$INPUT" "kotlinc main.kt -include-runtime -d main.jar && java -jar main.jar"