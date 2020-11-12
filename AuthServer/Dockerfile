FROM golang:latest

RUN mkdir /authserver
WORKDIR /authserver

COPY . ./

RUN go build -o authserver

CMD ./authserver