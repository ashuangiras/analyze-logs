FROM golang:latest

COPY . /app

WORKDIR /app

EXPOSE 8000

RUN go get -v -d
RUN go build -o main

HEALTHCHECK CMD curl --fail http://localhost:8000/health || exit 1

CMD ./main