FROM golang:latest

COPY . /src
WORKDIR /src

RUN go mod download
RUN go build -o auth ./cmd/auth/main.go


CMD ["/src/auth"]