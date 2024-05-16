FROM ghcr.io/gleam-lang/gleam:nightly-elixir

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y

COPY base .

CMD echo "$CODE" > main.gleam && ./start.sh "$TIMEOUT" "$INPUT" "gleam run main.gleam"