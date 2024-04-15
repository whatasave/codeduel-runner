# FROM mcr.microsoft.com/dotnet/samples:aspnetapp
FROM mcr.microsoft.com/dotnet/sdk:latest

WORKDIR /app

COPY base .

RUN apt-get update &&\
    apt-get install jq -y &&\
    dotnet new console

CMD echo "$CODE" > Program.cs && dotnet build . > /dev/null && ./start.sh "$TIMEOUT" "$INPUT" "dotnet run ."