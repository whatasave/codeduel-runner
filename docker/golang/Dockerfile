FROM golang:1.22 as build-stage

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y\
 && go mod init main\
 && go mod tidy

COPY base .

CMD echo "$CODE" > main.go && ./start.sh "$TIMEOUT" "$INPUT" "go run main.go"