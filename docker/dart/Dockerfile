FROM dart:2.14.4

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y

COPY base .

CMD echo "$CODE" > main.dart && ./start.sh "$TIMEOUT" "$INPUT" "dart main.dart"