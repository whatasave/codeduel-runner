FROM golang:1.22

WORKDIR /app

COPY base/ .

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh && \
    go mod init main && \
    go mod tidy


CMD ["/bin/sh", "-c", "echo \"$CODE\" > main.go && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"go run main.go\""]
