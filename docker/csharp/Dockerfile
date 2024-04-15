# FROM mcr.microsoft.com/dotnet/samples:aspnetapp
FROM mcr.microsoft.com/dotnet/sdk:8.0

WORKDIR /app

COPY base .

RUN apt-get update &&\
    apt-get install jq -y &&\
    dotnet new console

CMD echo "$CODE" > Program.cs && dotnet build . > /dev/null && ./start.sh "$TIMEOUT" "$INPUT" "./bin/Debug/net8.0/app"