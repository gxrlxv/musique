FROM golang:1.18

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o artist ./cmd/artist/main.go


CMD ["/src/artist"]