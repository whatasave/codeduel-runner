FROM ruby:3.0.3-slim

WORKDIR /app

RUN apt-get update\
 && apt-get install jq -y\
 && gem install bundler

COPY base .

CMD echo "$CODE" > main.rb && ./start.sh "$TIMEOUT" "$INPUT" "ruby main.rb"