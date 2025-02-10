FROM ruby:3.0.3-slim

WORKDIR /app

COPY base/ .

RUN apt-get update && \
    apt-get install -y --no-install-recommends jq && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh && \
    gem install bundler

CMD ["sh", "-c", "echo \"$CODE\" > main.rb && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"ruby main.rb\""]