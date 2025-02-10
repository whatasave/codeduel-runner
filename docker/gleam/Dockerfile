FROM ghcr.io/gleam-lang/gleam:nightly-elixir

WORKDIR /app

COPY base/ .

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh


CMD [ "sh", "-c", "echo \"$CODE\" > main.gleam && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"gleam run main.gleam\"" ]