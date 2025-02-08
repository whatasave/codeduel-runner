FROM python:3.12.1-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/*

COPY base .

CMD ["sh", "-c", "echo \"$CODE\" > main.py && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"python3 main.py\""]