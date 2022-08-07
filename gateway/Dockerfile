FROM golang:1.18

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o gateway ./cmd/gateway/main.go


CMD ["/src/gateway"]