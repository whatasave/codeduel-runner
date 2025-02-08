FROM ghcr.io/gleam-lang/gleam:nightly-elixir

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/*

COPY base .

CMD [ "sh", "-c", "echo \"$CODE\" > main.gleam && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"gleam run main.gleam\"" ]