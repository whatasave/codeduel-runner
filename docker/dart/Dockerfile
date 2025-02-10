FROM dart:2.14.4

WORKDIR /app

COPY base/ .

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh

CMD ["/bin/sh", "-c", "echo \"$CODE\" > main.dart && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"dart main.dart\""]
