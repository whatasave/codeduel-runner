FROM python:3.12.1-slim

WORKDIR /app

CMD echo $CODE > main.py && echo -n $INPUT | timeout $TIMEOUT python3 main.py