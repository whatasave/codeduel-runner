FROM python:3.12.1-slim

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y

COPY base .

CMD echo "$CODE" > main.py && ./start.sh "$TIMEOUT" "$INPUT" "python3 main.py"