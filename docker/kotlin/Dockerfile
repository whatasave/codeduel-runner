FROM openjdk:11-jre-slim AS BUILDER

COPY base/ .

RUN apt-get update && \
    apt-get install -y --no-install-recommends curl gnupg2 unzip jq && \
    rm -rf /var/lib/apt/lists/* && \
    chmod +x start.sh

ENV KOTLIN_VERSION=1.8.10

RUN curl -sSL "https://github.com/JetBrains/kotlin/releases/download/v${KOTLIN_VERSION}/kotlin-compiler-${KOTLIN_VERSION}.zip" -o kotlin.zip && \
    unzip kotlin.zip -d /opt && \
    rm kotlin.zip


ENV PATH="/opt/kotlinc/bin:${PATH}"

CMD ["sh", "-c", "echo \"$CODE\" > main.kt && ./start.sh \"$TIMEOUT\" \"$INPUT\" \"kotlinc -script main.kt\""]
