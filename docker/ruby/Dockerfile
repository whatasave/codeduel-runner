FROM ruby:3.0.3-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/* && \
    gem install bundler

COPY base/ .

CMD ["sh", "-c", "echo \"$CODE\" > main.rb && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"ruby main.rb\""]